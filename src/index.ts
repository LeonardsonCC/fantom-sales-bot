import dotenv from "dotenv";
import { ethers } from "ethers";
import { MarketplaceV2 } from "@paintswap/marketplace-interactions";
import onSold from "./onSold";

dotenv.config();

const provider = new ethers.providers.JsonRpcProvider(
  process.env.FANTOM_RPC_URL
);

const marketplace = new MarketplaceV2(provider);

console.log("Initializing listeners");
marketplace.onSold((sale) => {
  console.log("DEBUG SALE", sale);
  onSold(sale, provider);
});
