import { Sold } from "@paintswap/marketplace-interactions";
import axios from "axios";
import { ethers, BigNumber } from "ethers";
import { fetchPrice } from "./providers/coingecko";
import { fetchContractTx } from "./providers/ftmscan";
import { fetchContractName } from "./providers/generic-contract";
import {
  fetchItemHistory,
  fetchNftMetadata,
  HistoryItem,
} from "./providers/ps-api";
import { bigNumberToSimpleNumber } from "./utils/price";

const onSold = async (
  sale: Sold,
  provider: ethers.providers.JsonRpcProvider
) => {
  const timestamp = Date.now();
  const { nft } = await fetchItemHistory(
    sale.collection,
    sale.tokenID.toNumber()
  );

  const history = nft.history
    .filter(
      (historyItem) =>
        Number(historyItem.id) === sale.marketplaceId.toNumber() ||
        historyItem.action === "Sold"
    )
    .sort((a, b) => {
      const aTimestamp = Number(a.timestamp);
      const bTimestamp = Number(b.timestamp);
      if (aTimestamp < bTimestamp) {
        return -1;
      }
      if (aTimestamp > bTimestamp) {
        return 1;
      }
      return 0;
    });

  let lastSale: null | HistoryItem = null;
  let mintedOrBought: "Minted" | "Bought";
  if (history.length > 1) {
    lastSale = history.slice(-1)[0];
    mintedOrBought = "Bought";
  } else {
    const txs = await fetchContractTx(sale.collection);
    const sellerTxs = txs.filter((tx) => {
      return (
        tx.from.toLowerCase() === sale.seller.toLowerCase() &&
        tx.contractAddress === "" &&
        BigNumber.from(tx.value).gt(0)
      );
    });

    let mintQty = 0;
    let mintTx;
    let foundIt = false;
    for (const tx of sellerTxs) {
      if (foundIt) break;
      const receipt = await provider.getTransactionReceipt(tx.hash);

      mintQty = 0;
      for (const log of receipt.logs) {
        if (log.topics.length === 4) {
          mintQty++;
          if (
            log.topics.filter((receipt) =>
              BigNumber.from(receipt).eq(sale.tokenID)
            ).length
          ) {
            // Probably the mint event
            foundIt = true;
            mintTx = tx;
          }
        }
      }
    }
    if (mintTx) {
      console.log("mintvalue", BigNumber.from(mintTx.value));
      lastSale = {
        action: "0",
        actionId: "0",
        data: BigNumber.from(mintTx.value).mul(10).div(mintQty).toString(),
        hash: mintTx.hash,
        id: "0",
        numericId: "0",
        timestamp: mintTx.timeStamp,
        version: 2,
      };
    }
    mintedOrBought = "Minted";
  }

  let salePrice, lastSalePrice;
  try {
    const result = await fetchPrice("fantom", timestamp);
    salePrice = result.market_data.current_price.usd;
  } catch (e) {
    // TODO make it retry after some seconds
    console.error(e);
  }

  if (lastSale) {
    try {
      if (lastSale.timestamp.length === 10)
        lastSale.timestamp = `${lastSale.timestamp}000`;
      const result = await fetchPrice("fantom", Number(lastSale.timestamp));
      lastSalePrice = result.market_data.current_price.usd;
    } catch (err) {
      // TODO make it retry after some seconds
      console.error(err);
    }
  }

  const collectionName = await fetchContractName(provider, sale.collection);

  if (lastSale !== null && lastSalePrice && salePrice) {
    const roundValue = (value: number) => value.toFixed(3);

    let tweetMessage = "";
    tweetMessage += `🧾 Collection: ${collectionName}\n\n`;
    tweetMessage += `🛍 ${mintedOrBought}: ${roundValue(
      Number(ethers.utils.formatUnits(lastSale.data, 19))
    )} FTM @ $${roundValue(lastSalePrice)}\n`;
    tweetMessage += `💰 Sold: ${roundValue(
      Number(ethers.utils.formatUnits(sale.priceTotal))
    )} FTM @ $${salePrice.toFixed(3)}\n`;

    const saleTime = new Date(timestamp);
    const lastSaleTime = new Date(Number(lastSale.timestamp));

    tweetMessage += `\n🤝 HODL: ${Math.floor(
      (saleTime.getTime() - lastSaleTime.getTime()) / (1000 * 60 * 60 * 24)
    )} days\n`;

    const diff =
      bigNumberToSimpleNumber(sale.priceTotal) -
      bigNumberToSimpleNumber(BigNumber.from(lastSale.data), 19);
    if (diff >= 0) {
      tweetMessage += `📈 Gain: ${roundValue(diff)} FTM\n`;
    } else {
      tweetMessage += `📉 Loss: ${roundValue(diff)} FTM\n`;
    }

    const saleUsd = bigNumberToSimpleNumber(sale.priceTotal) * salePrice;
    const lastSaleUsd =
      bigNumberToSimpleNumber(BigNumber.from(lastSale.data), 19) *
      lastSalePrice;
    const diffUsd = saleUsd - lastSaleUsd;
    const percentDiffUsd = (diffUsd / lastSaleUsd) * 100;
    if (diffUsd >= 0) {
      tweetMessage += `💵 Profit: $${roundValue(
        diffUsd
      )} (📈 ${percentDiffUsd.toFixed(2)}%)\n`;
    } else {
      tweetMessage += `💵 Loss: $${roundValue(
        diffUsd
      )} (📉 ${percentDiffUsd.toFixed(2)}%)\n`;
    }

    tweetMessage += `https://paintswap.finance/marketplace/${sale.marketplaceId}`;
    console.log(tweetMessage, await fetchNftMetadata(nft.uri));
  } else {
    console.log("No last sale");
  }
};

export default onSold;
