// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "./ERC721.sol";

contract WTFApe is ERC721 {
    // 总量
    uint public MAX_APES = 10000;

    constructor(string memory _name, string memory _symbol) ERC721(_name, _symbol) {}

    function mint(address to, uint tokenId) external {
        require(to!=address(0) && tokenId<MAX_APES, "Invalid address");
        _mint(to, tokenId);
    }

    function _baseURI() override internal pure returns (string memory) {
        return "ipfs://QmeSjSinHpPnmXmspMjwiXyN6zS4E9zccariGR3jxcaWtq";
    }
}