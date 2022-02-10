import dotenv from "dotenv";
import nftkey from "./providers/nftkey";
import paintswap from "./providers/paintswap";
import Discord from "./providers/discord";
import * as Database from "./providers/database";
import onSold from "./handlers/onSold";
import { Marketplace } from "./types/marketplace";
import { ethers } from "ethers";

dotenv.config();

Database.connect();
Discord.init();

onSold({
  contract: "0x7f9893ef9726d23e249518eaa1424677b7fed6a9",
  date: new Date(),
  marketplace: Marketplace.NFTKEY,
  seller: "0xd1acf184fd8768d9d3683f29d8113253b186b65a",
  tokenId: ethers.BigNumber.from("722"),
  txHash: "0x13dccbf9974f73fb05547d7bcb38d244c6674e06d063d9a8b84c53bfb132c3ab",
  value: ethers.BigNumber.from("60000000000000000000"),
});
// nftkey.subscribe();
// paintswap.subscribe();
