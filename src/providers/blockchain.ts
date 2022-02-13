import { ethers } from "ethers";
import { Marketplace } from "../types/marketplace";
import { Sale } from "../types/sale";
import { fetchContractTx } from "./ftmscan";

let provider: null | ethers.providers.JsonRpcProvider = null;

const getProvider = () => {
  if (!provider) {
    provider = new ethers.providers.JsonRpcProvider(process.env.FANTOM_RPC_URL);
  }

  return provider;
};

const isValidAddress = (addr: string): boolean => {
  try {
    ethers.utils.getAddress(addr);
    return true;
  } catch (err) {
    return false;
  }
};

const getMintAsSale = async (
  contractAddress: string,
  tokenId: ethers.BigNumber,
  seller: string
): Promise<Sale | null> => {
  const provider = getProvider();
  const txs = await fetchContractTx(seller);
  const sellerTxs = txs.filter((tx) => {
    return (
      tx.from.toLowerCase() === seller.toLowerCase() &&
      tx.to.toLowerCase() === contractAddress.toLowerCase() &&
      tx.contractAddress === "" &&
      ethers.BigNumber.from(tx.value).gt(0)
    );
  });

  let mintQty = 0;
  let mintTx;
  let foundIt = false;
  for (const tx of sellerTxs) {
    if (foundIt) break;
    const receipt = await provider.getTransactionReceipt(tx.hash);

    mintQty = 0;
    for (const log of receipt.logs) {
      if (log.topics.length === 4) {
        mintQty++;
        if (
          log.topics.filter((receipt) =>
            ethers.BigNumber.from(receipt).eq(tokenId)
          ).length
        ) {
          // Probably the mint event
          foundIt = true;
          mintTx = tx;
        }
      }
    }
  }
  if (mintTx) {
    return {
      contract: contractAddress,
      tokenId: tokenId,
      value: ethers.BigNumber.from(mintTx.value).div(mintQty),
      date: new Date(Number(mintTx.timeStamp) * 1000),
      txHash: mintTx.hash,
      seller: mintTx.from,
      marketplace: Marketplace.MINTED,
    };
  }

  if (process.env.MINT_PRICE && Number(process.env.MINT_PRICE) > 0) {
    const date = new Date();
    date.setDate(date.getDate() - 1);
    return {
      contract: contractAddress,
      tokenId: tokenId,
      value: ethers.BigNumber.from(process.env.MINT_PRICE).mul(
        ethers.BigNumber.from(10).pow(18)
      ),
      date: date,
      txHash: "",
      seller: seller,
      marketplace: Marketplace.MINTED,
    };
  }

  return null;
};

export { getMintAsSale, isValidAddress };
export default getProvider;
