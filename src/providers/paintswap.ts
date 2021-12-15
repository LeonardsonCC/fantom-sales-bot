import { ethers } from "ethers";
import { PaintSwap__factory } from "../contracts/types";
import { TypedListener } from "../contracts/types/common";
import { SoldEvent } from "../contracts/types/PaintSwap";
import getProvider from "./blockchain";

const CONTRACT_ADDRESS = "0x6125fD14b6790d5F66509B7aa53274c93dAE70B9";

const initContract = () => {
  const contract = PaintSwap__factory.connect(CONTRACT_ADDRESS, getProvider());
  return contract;
};

const subscribe = () => {
  console.log("Subscribing...");

  const contract = initContract();

  contract.on(contract.filters.Sold(), onSold);
};

const onSold: TypedListener<SoldEvent> = (
  marketplaceId,
  nfts,
  tokenIds,
  amountBatches,
  price,
  buyer,
  seller
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

  // TODO send sale to handler
};

export default {
  subscribe,
};
