import { getMintAsSale } from "../providers/blockchain";
import nftkey from "../providers/nftkey";
import { Sale } from "../types/sale";
import { Action, makeMessage } from "../utils/message";

const onSold = async (sale: Sale) => {
  console.log("Sale received: ", sale);

  const nftkeyHistory = await nftkey.getTokenHistory(
    sale.contract,
    sale.tokenId
  );

  let history = [...nftkeyHistory];
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
  }
};

export default onSold;
