import { MessageEmbed } from "discord.js";
import { ethers } from "ethers";
import getProvider from "../../providers/blockchain";
import { fetchPrice } from "../../providers/coingecko";
import { fetchContractName } from "../../providers/generic-contract";
import { Marketplace } from "../../types/marketplace";
import { Sale } from "../../types/sale";
import { Action } from "../twitter/message";
import { bigNumberToSimpleNumber } from "./../../utils/price";

const roundValue = (value: number) => value.toFixed(3);

const makeMessage = async (
  sale: Sale,
  lastEvent: Sale,
  action: Action
): Promise<MessageEmbed | null> => {
  let salePrice, lastEventPrice;
  try {
    const saleResult = await fetchPrice("fantom", sale.date.getTime());
    salePrice = saleResult.market_data.current_price.usd;
  } catch (err) {
    console.log("error fetching fantom price", err);
  }

  try {
    const result = await fetchPrice("fantom", lastEvent.date.getTime());
    lastEventPrice = result.market_data.current_price.usd;
  } catch (err) {
    console.log("error fetching fantom price", err);
  }

  if (salePrice && lastEventPrice) {
    const collectionName = await fetchContractName(sale.contract);

    let gains = "";
    let gainOrLoss = "";
    const diff =
      bigNumberToSimpleNumber(sale.value) -
      bigNumberToSimpleNumber(lastEvent.value);
    if (diff >= 0) {
      gains = `${roundValue(diff)} FTM`;
      gainOrLoss = "Gain";
    } else {
      gains = `${roundValue(diff * -1)} FTM`;
      gainOrLoss = "Loss";
    }

    const saleUsd = bigNumberToSimpleNumber(sale.value) * salePrice;
    const lastSaleUsd =
      bigNumberToSimpleNumber(lastEvent.value) * lastEventPrice;
    const diffUsd = saleUsd - lastSaleUsd;
    const percentDiffUsd = (diffUsd / lastSaleUsd) * 100;

    let usdGains = "";
    let gainOrLossUsd = "";
    if (diffUsd >= 0) {
      usdGains += `$${roundValue(diffUsd)} (📈 ${percentDiffUsd.toFixed(2)}%)`;
      gainOrLossUsd = "Gain";
    } else {
      usdGains += `$${roundValue(diffUsd * -1)} (📉 ${percentDiffUsd.toFixed(
        2
      )}%)`;
      gainOrLossUsd = "Loss";
    }

    let url = "";
    let marketplaceName = "";
    switch (sale.marketplace) {
      case Marketplace.NFTKEY:
        url = `https://nftkey.app/token-details/?tokenAddress=${
          sale.contract
        }&tokenId=${sale.tokenId.toString()}`;
        marketplaceName = "🔑 NFTKEY";
        break;
      case Marketplace.PAINTSWAP:
        url = `https://paintswap.finance/marketplace/assets/${sale.contract}/${sale.tokenId}`;
        marketplaceName = "🖌️ PaintSwap";
        break;
    }

    const message = new MessageEmbed()
      .setColor(gainOrLoss === "Gain" ? "GREEN" : "RED")
      .setTitle(`${collectionName} - #${sale.tokenId.toString()}`)
      .setURL(url)
      .setTimestamp()
      .addFields([
        {
          name: "Marketplace",
          value: marketplaceName,
        },
        {
          name: action,
          value: `${roundValue(
            Number(ethers.utils.formatUnits(lastEvent.value))
          )} FTM @ $${lastEventPrice.toFixed(3)}`,
        },
        {
          name: "Sold",
          value: `${Number(ethers.utils.formatUnits(sale.value)).toFixed(
            3
          )} FTM @ $${salePrice.toFixed(3)}`,
        },
        {
          name: "HODL",
          value: `${Math.floor(
            (sale.date.getTime() - lastEvent.date.getTime()) /
              (1000 * 60 * 60 * 24)
          )} days`,
        },
        {
          name: gainOrLoss,
          value: gains,
        },
        {
          name: gainOrLossUsd,
          value: usdGains,
        },
      ]);

    return message;
  }
  return null;
};

export { makeMessage, Action };
