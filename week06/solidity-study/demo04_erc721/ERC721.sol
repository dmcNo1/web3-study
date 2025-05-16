// SPDX-License-Identifier: MIT
// by 0xAA
pragma solidity ^0.8.21;

import "./IERC165.sol";
import "./IERC721.sol";
import "./IERC721Receiver.sol";
import "./IERC721Metadata.sol";
import "./String.sol";

contract ERC721 is IERC721, IERC721Metadata{
    using Strings for uint256; // 使用Strings库，

    string public override name;

    string public override symbol;

    mapping(address => uint256) private _balanceOf;

    mapping(uint256 => address) private _ownerOf;

    mapping(address => mapping(address => bool)) private _operatorApprovals;

    mapping(uint256 => address) private _tokenApprovals;

    constructor(string memory _name, string memory _symbol) {
        name = _name;
        symbol = _symbol;
    }

    function tokenURI(uint256 tokenId) override external view returns (string memory) {
        
    }

    function supportsInterface(bytes4 interfaceId) override external pure returns (bool) {
        return interfaceId == type(IERC165).interfaceId ||
                interfaceId == type(IERC721).interfaceId ||
                interfaceId == type(IERC721Metadata).interfaceId;
    }

    function balanceOf(address owner) override external view returns (uint256 balance) {
        require(owner!=address(0), "owner = zero address");
        return _balanceOf[owner];
    }

    function ownerOf(uint256 tokenId) override public view returns (address owner) {
        owner = _ownerOf[tokenId];
        require(owner!=address(0), "token doesn't exist");
    }

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes calldata data
    ) override external {}

    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) override external {}

    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) override external {
        address owner = ownerOf(tokenId);
        require(_isApprovedOrOwner(owner, msg.sender, tokenId));
        _transfer(owner, from, to, tokenId);
    }

    function approve(address to, uint256 tokenId) override external {
        address owner = _ownerOf[tokenId];
        require(msg.sender==owner || _operatorApprovals[owner][msg.sender]);
        _approve(msg.sender, to, tokenId);
    }

    function setApprovalForAll(address operator, bool _approved) override external {
        _operatorApprovals[msg.sender][operator] = _approved;
    }

    function getApproved(uint256 tokenId) external view returns (address operator) {
        require(_ownerOf[tokenId] != address(0), "token doesn't exist");
        operator = _tokenApprovals[tokenId];
    }

    function isApprovedForAll(address owner, address operator) override external view returns (bool) {
        return _operatorApprovals[owner][operator];
    }

    // 授权函数。通过调整_tokenApprovals来，授权 to 地址操作 tokenId，同时释放Approval事件。
    function _approve(
        address owner,
        address to,
        uint tokenId
    ) private {
        _tokenApprovals[tokenId] = to;
        emit Approval(owner, to, tokenId);
    }


    // 查询 spender地址是否可以使用tokenId（需要是owner或被授权地址）
    function _isApprovedOrOwner(
        address owner,
        address spender,
        uint tokenId
    ) private view returns (bool) {
        return spender==owner || _operatorApprovals[owner][spender] || _tokenApprovals[tokenId]==spender;
    }

    /*
     * 转账函数。通过调整_balances和_owner变量将 tokenId 从 from 转账给 to，同时释放Transfer事件。
     * 条件:
     * 1. tokenId 被 from 拥有
     * 2. to 不是0地址
     */
    function _transfer(
        address owner,
        address from,
        address to,
        uint tokenId
    ) private {
        require(from == owner, "not owner");
        require(to != address(0), "transfer to the zero address");

        _approve(from, to, tokenId);
        _balanceOf[from] -= 1;
        _balanceOf[to] += 1;
        _tokenApprovals[tokenId] = to;

        emit Transfer(from, to, tokenId);
    }
}