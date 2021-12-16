import { ethers } from "ethers";
import { PaintSwap__factory } from "../contracts/types";
import { TypedListener } from "../contracts/types/common";
import { SoldEvent } from "../contracts/types/PaintSwap";
import getProvider from "./blockchain";
import { Sale } from "../types/sale";
import { Marketplace } from "../types/marketplace";
import axios, { AxiosResponse } from "axios";
import onSold from "../handlers/onSold";

const CONTRACT_ADDRESS = "0x6125fD14b6790d5F66509B7aa53274c93dAE70B9";

type ActionInfo = {
  sale: {
    addresses: string[];
    complete: true;
    endTime: string;
    id: string;
    isAuction: boolean;
    maxBid: string;
    maxBidder: string;
    maxOffer: string;
    maxOfferer: string;
    nextMinimum: string;
    price: string;
    seller: string;
    sold: boolean;
    startTime: string;
    tokenIds: string[];
    version: 1 | 2;
    minOffer: string;
    address: string;
    tokenId: string;
    isERC721: boolean;
  };
};

const initContract = () => {
  const contract = PaintSwap__factory.connect(CONTRACT_ADDRESS, getProvider());
  return contract;
};

const subscribe = () => {
  console.log("Subscribing...");

  const contract = initContract();

  contract.on(contract.filters.Sold(), onSoldHandler);
};

const onSoldHandler: TypedListener<SoldEvent> = async (
  marketplaceId,
  nfts,
  tokenIds,
  amountBatches,
  price,
  buyer,
  seller,
  _,
  event
) => {
  console.log(
    "PAINTSWAY: TOKEN SOLD: ",
    marketplaceId,
    nfts,
    tokenIds,
    amountBatches,
    price,
    buyer,
    seller
  );

  // Have to get the data from their API
  const { data: actionInfo } = await axios.get<
    undefined,
    AxiosResponse<ActionInfo>
  >(`https://api.paintswap.finance/v2/sales/${marketplaceId}`);
  console.log("RESULT API", actionInfo);

  const value = ethers.BigNumber.from(actionInfo.sale.price);

  const sale: Sale = {
    contract: actionInfo.sale.address,
    tokenId: ethers.BigNumber.from(actionInfo.sale.tokenId),
    value: ethers.BigNumber.from(value),
    date: new Date(),
    seller: actionInfo.sale.seller,
    txHash: event.transactionHash,
    marketplace: Marketplace.PAINTSWAP,
  };

  onSold(sale);
};

export default {
  subscribe,
};
