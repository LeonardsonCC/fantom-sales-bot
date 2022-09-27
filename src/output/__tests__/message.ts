import * as axios from "axios";
import { ethers } from "ethers";
import { Sale } from "../../types/sale";
import { makeMessage } from "../message";

const mockSale: Sale = {
  contract: '0x3c301c85b191b3edada11198ef1f5277c1ee1f87',
  tokenId: ethers.BigNumber.from({ _hex: '0x06d3'}),
  value: ethers.BigNumber.from({ _hex: '0x14d1120d7b160000'}),
  date: new Date(), //2022-09-27T00:28:39.525Z,
  seller: '0xbc622c6bcc9200c68d525adf5c8b1b9c3d1c6de3',
  txHash: '0x8d85951a5506baf5a7c8e962beac98318019c3e77ece684293030c6fcc565486',
  marketplace: 2
}

describe("Make simple message", () => {
  beforeAll(() => {
    process.env.FANTOM_RPC_URL = "https://rpc.ftm.tools/"
  });

  it("should create a simple message", async () => {
    let expected = "🖌️ PaintSwap\n";
    expected += "🧾 Collection: CursedTransistor\n";
    expected += "🖼️Token: #1747\n";
    expected += "\n";
    expected += "💰 Sold: 1.500 FTM @ $0.228\n";
    expected += "💸 Spent: $0.342";
    expected += "\n";
    expected += "https://paintswap.finance/marketplace/assets/0x3c301c85b191b3edada11198ef1f5277c1ee1f87/1747";

    const actual = await makeMessage(mockSale, null);

    expect(actual).toBe(expected);
  });
});
