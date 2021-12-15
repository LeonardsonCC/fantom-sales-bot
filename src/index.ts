import dotenv from "dotenv";
import { ethers } from "ethers";
import nftkey from "./providers/nftkey";
// import paintswap from "./providers/paintswap";

dotenv.config();

const provider = new ethers.providers.JsonRpcProvider(
  process.env.FANTOM_RPC_URL
);

nftkey.subscribe(provider);
// paintswap.subscribe(provider);
