import { ethers } from "ethers";
import GenericContract from "../contracts/artifacts/GenericContract.json";

const fetchContractName = async (
  provider: ethers.providers.JsonRpcProvider,
  contractAddress: string
) => {
  const contract = new ethers.Contract(
    contractAddress,
    GenericContract,
    provider
  );
  return await contract.name();
};

export { fetchContractName };
