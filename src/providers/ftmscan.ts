import axios, { AxiosResponse } from "axios";

interface IFTMScanTx {
  blockNumber: string;
  timeStamp: string;
  hash: string;
  nonce: string;
  blockHash: string;
  transactionIndex: string;
  from: string;
  to: string;
  value: string;
  gas: string;
  gasPrice: string;
  isError: string;
  txreceipt_status: string;
  input: string;
  contractAddress: string;
  cumulativeGasUsed: string;
  gasUsed: string;
  confirmations: string;
}

const fetchContractTx = async (
  contractAddress: string
): Promise<IFTMScanTx[]> => {
  console.log(
    `https://api.ftmscan.com/api?module=account&action=txlist&address=${contractAddress}&startblock=0&endblock=99999999&sort=asc&apikey=${process.env.FTMSCAN_API_KEY}`
  );
  const { data } = await axios.get(
    `https://api.ftmscan.com/api?module=account&action=txlist&address=${contractAddress}&startblock=0&endblock=99999999&sort=asc&apikey=${process.env.FTMSCAN_API_KEY}`
  );

  return data.result;
};

export { fetchContractTx };
