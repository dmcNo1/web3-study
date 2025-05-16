// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "@openzeppelin/contracts/access/Ownable.sol";
import "https://github.com/AmazingAng/WTF-Solidity/blob/main/34_ERC721/ERC721.sol";

contract DutchAuction is ERC721, Ownable {
    // NFT总量
    uint256 constant public COLLECTION_SIZE = 1000;
    // 荷兰拍卖起拍价，也是最高价
    uint256 constant public AUCTION_START_PRICE = 1 ether;
    // 荷兰拍卖结束价，也是最低价/地板价
    uint256 constant public AUCTION_END_PRICE = 0.1 ether;
    // 拍卖持续时长
    uint256 constant public AUCTION_TIME = 10 minutes;
    // 每过多久时间，价格衰减一次
    uint256 constant public AUCTION_DROP_INTERVAL = 1 minutes;
    // 每次价格衰减步长
    uint256 constant public AUCTION_DROP_PER_STEP = 
        (AUCTION_START_PRICE - AUCTION_END_PRICE) / (AUCTION_TIME / AUCTION_DROP_INTERVAL);
    // 拍卖起始时间（区块链时间戳，block.timestamp）
    uint256 public auctionStartTime;
    // metadata URI
    string private _baseTokenURI;   
    // 记录所有存在的tokenId
    uint256[] private _allTokens;

    constructor(string memory _name, string memory _symbol) Ownable(msg.sender) ERC721(_name, _symbol) {
        // 设置拍卖起始时间
        auctionStartTime = block.timestamp;
    }
    
    function setAuctionStartTime(uint32 time) external onlyOwner {
        auctionStartTime = time;
    }

    // 获取实时价格
    function getAuctionPrice() public view returns(uint256) {
        if (block.timestamp < auctionStartTime) {
            return AUCTION_START_PRICE;
        } else if (block.timestamp > auctionStartTime+AUCTION_TIME) {
            return AUCTION_END_PRICE;
        }
        uint256 times = (block.timestamp - auctionStartTime) / AUCTION_DROP_INTERVAL;
        return AUCTION_START_PRICE - times * AUCTION_DROP_PER_STEP;
    }

    // 参与拍卖
    /// @param quantity 拍卖的数量
    function auctionMint(uint256 quantity) external payable {
        uint256 _saleStartTime = uint256(auctionStartTime); // 建立local变量，减少gas花费
        require(_saleStartTime!=0 && _saleStartTime<=block.timestamp, "auction not start!");
        require(
            totalSupply()+quantity <= COLLECTION_SIZE, 
            "not enough remaining reserved for auction to support desired mint amount");
        
        // 计算花费
        uint256 totalCost = getAuctionPrice() * quantity;
        require(totalCost<=msg.value, "Need to send more ETH.");

        for (uint i=0; i<quantity; i++) {
            uint256 _mintIndex = totalSupply();
            _mint(msg.sender, _mintIndex);
            _addTokenToAllTokensEnumeration(_mintIndex);
        }

        // 多余ETH退款
        if (totalCost < msg.value) {
            // 注意一下这里是否有重入的风险
            payable(msg.sender).transfer(msg.value-totalCost);
        }
    }

    function totalSupply() public view virtual returns (uint256) {
        return _allTokens.length;
    }

    /**
     * Private函数，在_allTokens中添加一个新的token
     */
    function _addTokenToAllTokensEnumeration(uint256 tokenId) private {
        _allTokens.push(tokenId);
    }
}