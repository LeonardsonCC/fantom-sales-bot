import { ethers } from "ethers";
import { Marketplace } from "./marketplace";

export interface Sale {
  contract: string;
  tokenId: ethers.BigNumber;
  value: ethers.BigNumber;
  date: Date;
  txHash: string;
  seller: string;
  marketplace: Marketplace;
}
