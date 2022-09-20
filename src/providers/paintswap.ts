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

type NftInfo = {
  nft: {
    address: string;
    collections: string[];
    creator: string;
    history: {
      action: "Sold" | "UpdatePrice" | "ForSale" | "Auction";
      actionId: string;
      data: string;
      hash: string;
      id: string;
      numericId: string;
      timestamp: string;
      version: 1 | 2;
    }[];
    isERC721: boolean;
    isNSFW: boolean;
    lastSellPrice: boolean;
    num: string;
    owner: string;
    tokenId: string;
    uri: string;
  };
};

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

const getTokenHistory = async (
  contractAddress: string,
  tokenId: ethers.BigNumber
): Promise<Sale[]> => {
  const contract = initContract();

  // I know, this is the centralized way... but paintswap does not like to expose his contracts :/
  const { nft } = await fetchTokenInfo(contractAddress, tokenId);

  const sales: Sale[] = nft.history
    .filter((item) => item.action === "Sold")
    .map((item): Sale => {
      return {
        contract: contractAddress,
        seller: "", // it's a sale, doesn't need the seller to get the mint
        tokenId: ethers.BigNumber.from(tokenId),
        value: ethers.BigNumber.from(item.data),
        txHash: item.hash,
        date: new Date(Number(item.timestamp + "000")), // 3 zeros for the seconds :p
        marketplace: Marketplace.PAINTSWAP,
      };
    });

  return sales;
};

const fetchTokenInfo = async (
  contractAddress: string,
  tokenId: ethers.BigNumber
) => {
  const { data } = await axios.get<any, AxiosResponse<NftInfo>>(
    `https://api.paintswap.finance/nft/${contractAddress}/${tokenId.toString()}`,
    {
      params: {
        numToFetch: 10,
        allowNSFW: true,
      },
    }
  );

  return data;
};

export default {
  subscribe,
  getTokenHistory,
};
