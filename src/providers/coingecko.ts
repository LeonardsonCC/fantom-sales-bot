import axios from "axios";

interface CoinGeckoResponse {
  market_data: {
    current_price: {
      usd: number;
    };
  };
}

const fetchPrice = async (
  coin: "fantom" | "paint-swap",
  timestamp: number
): Promise<CoinGeckoResponse> => {
  let time;
  if (String(timestamp).length < 13) {
    time = new Date(timestamp * 1000); // idk why but it works
  } else {
    time = new Date(timestamp);
  }
  const timeString: string = `${String(time.getDate()).padStart(
    2,
    "0"
  )}-${String(time.getMonth() + 1).padStart(2, "0")}-${time.getFullYear()}`;

  let data;
  const response = await axios.get(
    `https://api.coingecko.com/api/v3/coins/${coin}/history?date=${timeString}`
  );
  data = response.data;

  return data;
};

export { fetchPrice };
