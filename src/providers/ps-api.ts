import axios from "axios";
import { ByteLengthQueuingStrategy } from "stream/web";

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

interface ERC711Metadata {
  name: string;
  description: string;
  image: string;
  attributes: {
    trait_type: string;
    value: string;
  }[];
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

const fetchNftMetadata = async (
  uri: string
): Promise<ERC711Metadata | null> => {
  try {
    const { data } = await axios.get(uri);

    return data;
  } catch (err) {
    console.log("error fetching metadata");
    return null;
  }
};

const fetchNftImage = async (uri: string) => {
  try {
    const { data } = await axios.get(uri, {
      responseType: "arraybuffer",
    });

    return data;
  } catch (err) {
    console.log("error fetching image");
    return {};
  }
};

export { HistoryItem, fetchItemHistory, fetchNftMetadata, fetchNftImage };
