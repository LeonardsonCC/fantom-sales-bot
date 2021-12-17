import { getMintAsSale } from "../providers/blockchain";
import nftkey from "../providers/nftkey";
import paintswap from "../providers/paintswap";
import { tweet } from "../providers/twitter";
import { Sale } from "../types/sale";
import { Action, makeMessage } from "../utils/message";
import { getTokenUri } from "../providers/generic-contract";
import axios from "axios";

const onSold = async (sale: Sale) => {
  console.log("Sale received: ", sale);

  const nftkeyHistory = await nftkey.getTokenHistory(
    sale.contract,
    sale.tokenId
  );

  const paintswapHistory = await paintswap.getTokenHistory(
    sale.contract,
    sale.tokenId
  );

  let history = [...nftkeyHistory, ...paintswapHistory];
  history = history.filter((item) => item.txHash !== sale.txHash);

  let message = "";
  if (history.length > 0) {
    // have sales after mint
    history.sort((a, b) => {
      if (a.date < b.date) {
        return -1;
      }
      if (a.date > b.date) {
        return 1;
      }
      return 0;
    });

    const bought = history[history.length - 1];
    message = await makeMessage(sale, bought, Action.BOUGHT);
  } else {
    const minted = await getMintAsSale(
      sale.contract,
      sale.tokenId,
      sale.seller
    );
    console.log("MINTED", minted);
    if (minted) {
      message = await makeMessage(sale, minted, Action.MINTED);
    }
  }

  if (message) {
    console.log("MESSAGE: ", message);

    try {
      const tokenUri = await getTokenUri(sale.contract, sale.tokenId);
      const { data } = await axios.get(
        tokenUri.replace("ipfs://", "https://cloudflare-ipfs.com/ipfs/")
      );
      console.log("DATA", data);
      const { data: image } = await axios.get(
        data.image.replace("ipfs://", "https://cloudflare-ipfs.com/ipfs/"),
        {
          responseType: "arraybuffer",
        }
      );

      console.log(typeof image);
      tweet(message, image);
    } catch (err) {
      console.log("can't get the image...", err);
      tweet(message);
    }
    return;
  } else {
    console.log("Some error ocurred during message creation");
  }
};

export default onSold;
