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
  const time = new Date(timestamp); // idk why but it works
  const timeString: string = `${String(time.getDay()).padStart(
    2,
    "0"
  )}-${String(time.getMonth()).padStart(2, "0")}-${time.getFullYear()}`;

  let data;
  try {
    const response = await axios.get(
      `https://api.coingecko.com/api/v3/coins/${coin}/history?date=${timeString}`
    );
    data = response.data;
  } catch (err) {
    setTimeout(() => {
      console.log("Error fetching price... Waiting 5s and trying again");
      fetchPrice(coin, timestamp);
    }, 5000);
  }

  return data;
};

export { fetchPrice };
