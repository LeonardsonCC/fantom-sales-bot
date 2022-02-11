import dotenv from "dotenv";
import nftkey from "./providers/nftkey";
import paintswap from "./providers/paintswap";
import Discord from "./providers/discord";
import * as Database from "./providers/database";

dotenv.config();

Database.connect();
Discord.init();

nftkey.subscribe();
paintswap.subscribe();
