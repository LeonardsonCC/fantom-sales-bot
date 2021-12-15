import dotenv from "dotenv";
import { ethers } from "ethers";
import onSold from "./handlers/onSold";
import nftkey from "./providers/nftkey";
import paintswap from "./providers/paintswap";

dotenv.config();

nftkey.subscribe();
paintswap.subscribe();

// onSold({
//   contract: "0x7aCeE5D0acC520faB33b3Ea25D4FEEF1FfebDE73",
//   tokenId: ethers.BigNumber.from(4982),
//   value: ethers.BigNumber.from(240),
//   date: new Date(),
//   txHash: "0xd4ae7f2f1e7d1c8630dfde0d3f5922a10ea5d7d76a6b38ac60ce94f054dac89a",
//   seller: "0xA6e950aa70EBaAf99686A5d95aFe8aca8B5E353B",
//   marketplace: 0,
// });
