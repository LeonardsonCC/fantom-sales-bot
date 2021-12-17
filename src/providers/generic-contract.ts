import { ethers } from "ethers";
import { GenericContract__factory } from "../contracts/types";
import getProvider from "./blockchain";

const fetchContractName = async (contractAddress: string) => {
  const contract = GenericContract__factory.connect(
    contractAddress,
    getProvider()
  );
  return await contract.name();
};

const getTokenUri = async (
  contractAddress: string,
  tokenId: ethers.BigNumber
) => {
  const contract = GenericContract__factory.connect(
    contractAddress,
    getProvider()
  );
  return await contract.tokenURI(tokenId);
};

export { fetchContractName, getTokenUri };
