import { Sold } from "@paintswap/marketplace-interactions";
import { ethers } from "ethers";
import { fetchPrice } from "./providers/coingecko";
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
