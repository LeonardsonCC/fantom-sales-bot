import dotenv from "dotenv";
import nftkey from "./providers/nftkey";
// import paintswap from "./providers/paintswap";

dotenv.config();

nftkey.subscribe();
// paintswap.subscribe(provider);
