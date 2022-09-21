import { getMintAsSale } from "../providers/blockchain";
import { tweet } from "../providers/twitter";
import { Sale } from "../types/sale";
import { Action, makeMessage } from "../utils/message";
import { getTokenUri } from "../providers/generic-contract";
import axios from "axios";
import sharp from "sharp";
import { getTokenHistory } from "../marketplaces";

const onSold = async (sale: Sale) => {
  console.log("Sale received: ", sale);

  let history = await getTokenHistory(sale.contract, sale.tokenId);
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

      let filetype: "PNG" | "JPG" | "GIF" = data.image
        .substr(data.image.length - 3)
        .toUpperCase();
      if (!["PNG", "GIF", "JPG"].includes(filetype)) filetype = "PNG";

      const imageResized: any = await sharp(image).resize(600, 600).toBuffer();

      tweet(message, imageResized, filetype);
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
