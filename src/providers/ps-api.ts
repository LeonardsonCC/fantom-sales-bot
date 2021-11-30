import axios from "axios";

interface ItemDetails {
  nft: {
    address: string;
    creator: string;
    history: HistoryItem[];
    tokenId: string;
    uri: string;
  };
}

interface HistoryItem {
  action: string;
  actionId: string;
  data: string;
  hash: string;
  id: string;
  numericId: string;
  timestamp: string;
  version: number;
}

const fetchItemHistory = async (
  contract: string,
  tokenId: number
): Promise<ItemDetails> => {
  const { data } = await axios.get(
    `https://api.paintswap.finance/nft/${contract}/${tokenId}`,
    {
      params: {
        allowNFSW: true,
        numToFetch: 10,
      },
    }
  );

  return data;
};

export { HistoryItem, fetchItemHistory };
