import nftkey from "../providers/nftkey";
import { Sale } from "../types/sale";

const onSold = (sale: Sale) => {
  console.log("Sale received: ", sale);

  nftkey.getTokenHistory(sale.contract, sale.tokenId);
}

export default onSold;
