import { ethers } from "ethers";
import nftkey from "./nftkey";
import paintswapV3 from "./paintswap-v3";

const marketplaces = [nftkey, paintswapV3];

export function subscribe() {
  marketplaces.forEach((mp) => mp.subscribe());
}

export async function getTokenHistory(
  contractAddress: string,
  tokenId: ethers.BigNumber
) {
  try {
    const histories = await Promise.all(
      marketplaces.map((mp) => mp.getTokenHistory(contractAddress, tokenId))
    );

    return histories.flat(1);
  } catch (e) {
    return [];
  }
}
