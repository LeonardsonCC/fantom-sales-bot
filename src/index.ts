import dotenv from "dotenv";
import { ethers } from "ethers";
import { MarketplaceV2, Sold } from "@paintswap/marketplace-interactions";
import { fetchItemHistory, HistoryItem } from "./providers/ps-api";
import { fetchPrice } from "./providers/coingecko";

dotenv.config();

const provider = new ethers.providers.JsonRpcProvider(
  process.env.FANTOM_RPC_URL
);

const marketplace = new MarketplaceV2(provider);

console.log("Initializing listeners");
marketplace.onSold(async (sale) => {
  const timestamp = Date.now();
  const { nft } = await fetchItemHistory(
    sale.collection,
    sale.tokenID.toNumber()
  );

console.log("Initializing listeners");
marketplace.onSold((sale) => {
  onSold(sale, provider);
});
