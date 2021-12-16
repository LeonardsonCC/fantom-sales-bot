import dotenv from "dotenv";
import { ethers } from "ethers";
import onSold from "./handlers/onSold";
import nftkey from "./providers/nftkey";
import paintswap from "./providers/paintswap";

dotenv.config();

nftkey.subscribe();
paintswap.subscribe();
