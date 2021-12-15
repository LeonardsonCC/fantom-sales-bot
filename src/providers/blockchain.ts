import { ethers } from "ethers";

let provider: null | ethers.providers.JsonRpcProvider = null;

const getProvider = () => {
  if (!provider) {
    provider = new ethers.providers.JsonRpcProvider(
      process.env.FANTOM_RPC_URL
    );
  }

  return provider;
};

export default getProvider;
