import axios from "axios";

interface CoinGeckoResponse {
  market_data: {
    current_price: {
      usd: number;
    };
  };
}

const fetchPrice = async (
  coin: string,
  timestamp: number
): Promise<CoinGeckoResponse> => {
  const time = new Date(timestamp); // idk why but it works
  const timeString: string = `${String(time.getDay()).padStart(
    2,
    "0"
  )}-${String(time.getMonth()).padStart(2, "0")}-${time.getFullYear()}`;

  const { data } = await axios.get(
    `https://api.coingecko.com/api/v3/coins/${coin}/history?date=${timeString}`
  );

  return data;
};

export { fetchPrice };
