import { getMintAsSale } from "../providers/blockchain";
import nftkey from "../providers/nftkey";
import paintswap from "../providers/paintswap";
import { Sale } from "../types/sale";
import { Action } from "../providers/twitter/message";
import { getTokenUri } from "../providers/generic-contract";
import axios from "axios";
import sharp from "sharp";
import { providers } from "../providers/posters";
import { ethers } from "ethers";

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

  let action: Action | undefined = undefined;
  let actionBefore: Sale | undefined = undefined;
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

    actionBefore = history[history.length - 1];
    action = Action.BOUGHT;
  } else {
    const mint = await getMintAsSale(sale.contract, sale.tokenId, sale.seller);
    console.log("MINTED", mint);
    if (mint) {
      action = Action.MINTED;
      actionBefore = mint;
    }
  }

  console.log("HAS ACTION BEFORE: ", action, actionBefore);

  try {
    const image = await getImage(sale.contract, sale.tokenId);

    if (!image) throw new Error("Can't find image");

    providers.forEach((provider) => {
      if (
        typeof actionBefore !== "undefined" &&
        typeof action !== "undefined"
      ) {
        provider.post(
          sale,
          actionBefore,
          action,
          image.imageResized,
          image.filetype
        );
      }
    });
  } catch (err) {
    console.log(err);
    providers.forEach((provider) => {
      if (
        typeof actionBefore !== "undefined" &&
        typeof action !== "undefined"
      ) {
        provider.post(sale, actionBefore, action);
      }
    });
  }
  return;
};

const getImage = async (contract: string, tokenId: ethers.BigNumber) => {
  try {
    const tokenUri = await getTokenUri(contract, tokenId);
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

    return {
      imageResized,
      filetype,
    };
  } catch (err) {
    return null;
  }
};

export default onSold;
