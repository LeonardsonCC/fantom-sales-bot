import { ethers } from "ethers";
import { Marketplace } from "./marketplace";

export interface Sale {
  contract: string;
  tokenId: ethers.BigNumber;
  value: ethers.BigNumber;
  date: Date;
  marketplace: Marketplace;
}
