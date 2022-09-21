import dotenv from "dotenv";
import { ethers } from "ethers";
import onSold from "./handlers/onSold";
import { subscribe } from "./marketplaces";

dotenv.config();

subscribe();
