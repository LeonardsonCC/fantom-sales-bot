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

const CONTRACT_ADDRESS = "0x1a7d6ed890b6c284271ad27e7abe8fb5211d0739";

const subscribe = (provider: ethers.providers.JsonRpcProvider) => {
  console.log("Subscribing...");

  const contract = NFTKey__factory.connect(CONTRACT_ADDRESS, provider);

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
    marketplace: Marketplace.NFTKEY,
  };

  onSold(sale);
};

export default {
  subscribe,
};
