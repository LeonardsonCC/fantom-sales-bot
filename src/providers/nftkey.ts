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

  if (process.env.COLLECTION) {
    contract.on(
      contract.filters.TokenBought(process.env.COLLECTION),
      onTokenBought
    );
    contract.on(
      contract.filters.TokenBidAccepted(process.env.COLLECTION),
      onTokenBidAccepted
    );
  } else {
    contract.on(contract.filters.TokenBought(), onTokenBought);
    contract.on(contract.filters.TokenBidAccepted(), onTokenBidAccepted);
  }
};

const onTokenBought: TypedListener<TokenBoughtEvent> = async (
  erc721Address,
  tokenId,
  buyer,
  listing,
  serviceFee,
  royaltyFee,
  event
) => {
  console.log(
    "TOKEN BOUGHT: ",
    erc721Address,
    tokenId,
    buyer,
    listing,
    serviceFee,
    royaltyFee,
    event
  );

  const sale: Sale = {
    contract: erc721Address,
    tokenId: tokenId,
    value: listing.value,
    date: new Date(),
    txHash: event.transactionHash,
    seller: listing.seller,
    marketplace: Marketplace.NFTKEY,
  };

  onSold(sale);
};

const onTokenBidAccepted: TypedListener<TokenBidAcceptedEvent> = async (
  erc721Address,
  tokenId,
  seller,
  bid,
  serviceFee,
  royaltyFee,
  event
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
    txHash: event.transactionHash,
    seller: seller,
    marketplace: Marketplace.NFTKEY,
  };

  onSold(sale);
};

const getTokenHistory = async (
  contractAddress: string,
  tokenId: ethers.BigNumber
): Promise<Sale[]> => {
  const contract = initContract();

  // Get old sales from filtering by nft
  const tokenBoughtFilter = contract.filters.TokenBought(
    contractAddress,
    tokenId
  );
  const tokenBoughtEvents = await contract.queryFilter(tokenBoughtFilter);

  const tokenBidAcceptedFilter = contract.filters.TokenBidAccepted(
    contractAddress,
    tokenId
  );
  const tokenBidAcceptedEvents = await contract.queryFilter(
    tokenBidAcceptedFilter
  );

  const sales: Sale[] = [];
  for (const boughtEvent of tokenBoughtEvents) {
    const tx = await boughtEvent.getBlock();

    const date = new Date(tx.timestamp * 1000);

    sales.push({
      contract: contractAddress,
      tokenId: tokenId,
      value: boughtEvent.args.listing.value,
      date: date,
      txHash: boughtEvent.transactionHash,
      seller: boughtEvent.args.listing.seller,
      marketplace: Marketplace.NFTKEY,
    });
  }

  for (const tokenBidAcceptedEvent of tokenBidAcceptedEvents) {
    const tx = await tokenBidAcceptedEvent.getTransaction();

    const date = new Date(tx.timestamp!);

    sales.push({
      contract: contractAddress,
      tokenId: tokenId,
      value: tokenBidAcceptedEvent.args.bid.value,
      date: date,
      txHash: tokenBidAcceptedEvent.transactionHash,
      seller: tokenBidAcceptedEvent.args.seller,
      marketplace: Marketplace.NFTKEY,
    });
  }

  return sales;
};

export default {
  subscribe,
  getTokenHistory,
};
