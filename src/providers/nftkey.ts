import { ethers } from "ethers";
import { NFTKey__factory } from "../contracts/types";
import { TypedListener } from "../contracts/types/common";
import {
  TokenBidAcceptedEvent,
  TokenBoughtEvent,
} from "../contracts/types/NFTKey";
import onSold from "../handlers/onSold";
import { Marketplace } from "../types/marketplace";
import { Sale } from "../types/sale";
import getProvider from "./blockchain";

const CONTRACT_ADDRESS = "0x1a7d6ed890b6c284271ad27e7abe8fb5211d0739";

const initContract = () => {
  const contract = NFTKey__factory.connect(CONTRACT_ADDRESS, getProvider());

  return contract;
};

const subscribe = () => {
  console.log("Subscribing...");

  const contract = initContract();

  contract.on(contract.filters.TokenBought(), onTokenBought);
  contract.on(contract.filters.TokenBidAccepted(), onTokenBidAccepted);
};

const onTokenBought: TypedListener<TokenBoughtEvent> = (
  erc721Address,
  tokenId,
  buyer,
  listing,
  serviceFee,
  royaltyFee
) => {
  console.log(
    "TOKEN BOUGHT: ",
    erc721Address,
    tokenId,
    buyer,
    listing,
    serviceFee,
    royaltyFee
  );

  const sale: Sale = {
    contract: erc721Address,
    tokenId: tokenId,
    value: listing.value,
    date: new Date(),
    marketplace: Marketplace.NFTKEY,
  };

  onSold(sale);
};

const onTokenBidAccepted: TypedListener<TokenBidAcceptedEvent> = (
  erc721Address,
  tokenId,
  seller,
  bid,
  serviceFee,
  royaltyFee
) => {
  console.log(
    "TOKEN BID ACCEPTED: ",
    erc721Address,
    tokenId,
    seller,
    bid,
    serviceFee,
    royaltyFee
  );

  const sale: Sale = {
    contract: erc721Address,
    tokenId: tokenId,
    value: bid.value,
    date: new Date(),
    marketplace: Marketplace.NFTKEY,
  };

  onSold(sale);
};

const getTokenHistory = async (
  contractAddress: string,
  tokenId: ethers.BigNumber
) => {
  const contract = initContract();

  const tokenBoughtFilter = contract.filters.TokenBought();
  const tokenBoughtEvents = await contract.queryFilter(tokenBoughtFilter);

  const tokenBidAcceptedFilter = contract.filters.TokenBidAccepted();
  const tokenBidAcceptedEvents = await contract.queryFilter(
    tokenBidAcceptedFilter
  );

  console.log("ALL SALES", tokenBoughtEvents, onTokenBidAccepted);
};

export default {
  subscribe,
  getTokenHistory,
};
