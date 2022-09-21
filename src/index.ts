import dotenv from "dotenv";
import { ethers } from "ethers";
import onSold from "./handlers/onSold";
import listenMarketplacesSales from "./marketplaces";

dotenv.config();

listenMarketplacesSales();
