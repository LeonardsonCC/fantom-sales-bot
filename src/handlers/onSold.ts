import nftkey from "../providers/nftkey";
import { Sale } from "../types/sale";

const onSold = async (sale: Sale) => {
  console.log("Sale received: ", sale);

  const nftkeyHistory = await nftkey.getTokenHistory(
    sale.contract,
    sale.tokenId
  );

  const history = [...nftkeyHistory];

  console.log("HISTORY", history)
};

export default onSold;
