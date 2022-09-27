import { ethers } from "ethers";
import { fetchPrice } from "../providers/coingecko";
import { fetchContractName } from "../providers/generic-contract";
import { Marketplace } from "../types/marketplace";
import { Sale } from "../types/sale";
import { bigNumberToSimpleNumber } from "./price";

const roundValue = (value: number) => value.toFixed(3);

const makeMessage = async (
  sale: Sale,
  lastEvent: Sale | null
): Promise<string> => {
  if (lastEvent) {
    return await makeMessageWithLastEvent(sale, lastEvent);
  } else {
    return await makeMessageSimple(sale);
  }
};

const makeMessageSimple = async (sale: Sale): Promise<string> => {
  let salePrice;
  try {
    const saleResult = await fetchPrice("fantom", sale.date.getTime());
    salePrice = saleResult.market_data.current_price.usd;
  } catch (err) {
    console.log("error fetching fantom price", err);
  }

  if (salePrice) {
    const collectionName = await fetchContractName(sale.contract);

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

    let message = "";
    message += `${marketplaceName}\n`;
    message += `🧾 Collection: ${collectionName}\n`;
    message += `🖼️Token: #${sale.tokenId.toString()}\n\n`;

    message += `💰 Sold: ${Number(ethers.utils.formatUnits(sale.value)).toFixed(
      3
    )} FTM @ $${salePrice.toFixed(3)}\n`;

    message += `💸 Spent: $${(Number(ethers.utils.formatUnits(sale.value)) * salePrice).toFixed(3)}`;

    message += "\n";
    message += `${url}`;

    return message;
  }
  return "";
};

const makeMessageWithLastEvent = async (
  sale: Sale,
  lastEvent: Sale
): Promise<string> => {
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
    const diff =
      bigNumberToSimpleNumber(sale.value) -
      bigNumberToSimpleNumber(lastEvent.value);
    if (diff >= 0) {
      gains += `📈 Gain: ${roundValue(diff)} FTM`;
    } else {
      gains += `📉 Loss: ${roundValue(diff * -1)} FTM`;
    }

    const saleUsd = bigNumberToSimpleNumber(sale.value) * salePrice;
    const lastSaleUsd =
      bigNumberToSimpleNumber(lastEvent.value) * lastEventPrice;
    const diffUsd = saleUsd - lastSaleUsd;
    const percentDiffUsd = (diffUsd / lastSaleUsd) * 100;

    let usdGains = "";
    if (diffUsd >= 0) {
      usdGains += `💵 Profit: $${roundValue(
        diffUsd
      )} (📈 ${percentDiffUsd.toFixed(2)}%)`;
    } else {
      usdGains += `💵 Loss: $${roundValue(
        diffUsd * -1
      )} (📉 ${percentDiffUsd.toFixed(2)}%)`;
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

    let message = "";
    message += `${marketplaceName}\n`;
    message += `🧾 Collection: ${collectionName}\n`;
    message += `🖼️Token: #${sale.tokenId.toString()}\n\n`;

    message += `🛍 Bought: ${roundValue(
      Number(ethers.utils.formatUnits(lastEvent.value))
    )} FTM @ $${lastEventPrice.toFixed(3)}\n`;
    message += `💰 Sold: ${Number(ethers.utils.formatUnits(sale.value)).toFixed(
      3
    )} FTM @ $${salePrice.toFixed(3)}\n\n`;

    message += `🤝 HODL: ${Math.floor(
      (sale.date.getTime() - lastEvent.date.getTime()) / (1000 * 60 * 60 * 24)
    )} days\n`;

    // Show the gains
    message += `${gains}\n`;
    message += `${usdGains}\n`;
    message += `${url}`;

    return message;
  }
  return "";
};

export { makeMessage };
