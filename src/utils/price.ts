import { ethers, BigNumber } from "ethers";

const bigNumberToSimpleNumber = (
  value: BigNumber,
  decimals: number = 18
): number => {
  return Number(ethers.utils.formatUnits(value, decimals));
};

export { bigNumberToSimpleNumber };
