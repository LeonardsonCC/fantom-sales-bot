import { ethers } from "ethers";
import getProvider from "../providers/blockchain";
import { fetchPrice } from "../providers/coingecko";
import { fetchContractName } from "../providers/generic-contract";
import { Marketplace } from "../types/marketplace";
import { Sale } from "../types/sale";
import { bigNumberToSimpleNumber } from "./price";

enum Action {
  MINTED = "Minted",
  BOUGHT = "Bought",
}

const roundValue = (value: number) => value.toFixed(3);

const makeMessage = async (
  sale: Sale,
  lastEvent: Sale,
  action: Action
): Promise<string> => {
  let salePrice, lastEventPrice;
  try {
    console.log("TIME", sale.date.getTime());
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
    const collectionName = await fetchContractName(
      getProvider(),
      sale.contract
    );

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
    switch (sale.marketplace) {
      case Marketplace.NFTKEY:
        url = `https://nftkey.app/token-details/?tokenAddress=${
          sale.contract
        }&tokenId=${sale.tokenId.toString()}`;
    }

    return `
      🧾 Collection: ${collectionName}
      🖼️Token: #${sale.tokenId.toString()}

      🛍 ${action}: ${roundValue(
      Number(ethers.utils.formatUnits(lastEvent.value))
    )} FTM @ $${lastEventPrice.toFixed(3)}
      💰 Sold: ${roundValue(
        Number(ethers.utils.formatUnits(sale.value))
      )} FTM @ $${salePrice.toFixed(3)}   

      🤝 HODL: ${Math.floor(
        (sale.date.getTime() - lastEvent.date.getTime()) / (1000 * 60 * 60 * 24)
      )} days
      ${gains}
      ${usdGains}
      ${url}
    `;
  }
  return "";
};

export { makeMessage, Action };
