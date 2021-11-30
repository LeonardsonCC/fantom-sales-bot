// SPDX-License-Identifier: MIT
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

abstract contract GenericContract is ERC721 {
    constructor() {}

    function _baseURI() internal override pure returns (string memory) {}

    function claim() public payable {}

}
