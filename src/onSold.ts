import { Sold } from "@paintswap/marketplace-interactions";
import { ethers, BigNumber } from "ethers";
import { isConstValueNode } from "graphql";
import { fetchPrice } from "./providers/coingecko";
import { fetchContractTx } from "./providers/ftmscan";
import { fetchContractName } from "./providers/generic-contract";
import { fetchItemHistory, HistoryItem } from "./providers/ps-api";

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
  if (history.length > 1) {
    lastSale = history.slice(-1)[0];
  } else {
    const txs = await fetchContractTx(sale.collection);
    const sellerTxs = txs.filter((tx) => {
      return (
        tx.from.toLowerCase() === sale.seller.toLowerCase() &&
        tx.contractAddress === "" &&
        BigNumber.from(tx.value).gt(0)
      );
    });

    let foundIt = false;
    for (const tx of sellerTxs) {
      if (foundIt) break;
      const receipt = await provider.getTransactionReceipt(tx.hash);

      for (const log of receipt.logs) {
        if (foundIt) break;

        if (log.topics.length === 4) {
          if (
            log.topics.filter((receipt) =>
              BigNumber.from(receipt).eq(sale.tokenID)
            ).length
          ) {
            // Probably the mint event
            foundIt = true;
            lastSale = {
              action: "0",
              actionId: "0",
              data: BigNumber.from(tx.value).toString(),
              hash: tx.hash,
              id: "0",
              numericId: "0",
              timestamp: tx.timeStamp,
              version: 2,
            };
            break;
          }
        }
      }
    }
    console.log("finished", sale, lastSale);
  }

  try {
    const {
      market_data: {
        current_price: { usd: salePrice },
      },
    } = await fetchPrice("fantom", timestamp);
  } catch (e) {
    // TODO make it retry after some seconds
    console.error(e);
  }

  if (lastSale) {
    try {
      if (lastSale.timestamp.length === 10)
        lastSale.timestamp = `${lastSale.timestamp}000`;
      const {
        market_data: {
          current_price: { usd: lastSalePrice },
        },
      } = await fetchPrice("fantom", Number(lastSale.timestamp));
    } catch (err) {
      // TODO make it retry after some seconds
      console.error(err);
    }
  }

  const collectionName = await fetchContractName(provider, sale.collection);
};

export default onSold;
