// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "./IERC20.sol";

// ERC20代币的水龙头合约
contract Faucet {
    // 记录ERC20合约地址
    address private tokenContract;

    // 每次领 100 单位代币
    uint256 private amountAllowed = 100;    
    
    // 记录已经获取过水龙头的地址
    mapping(address => bool) private requestedAddress;

    constructor(address _address) {
        tokenContract = _address;
    }

    // 转账记录
    event SendToken(address indexed Receiver, uint256 indexed Amount);

    // 用户领取代币函数
    function requestTokens() external {
        // 校验请求地址还没有申请过水龙头
        require(!requestedAddress[msg.sender], "Faucet: Address already requested!");
        
        // 余额还足够
        IERC20 token = IERC20(tokenContract);
        require(token.balanceOf(address(this))>=amountAllowed, "Faucet: Faucet empty!");

        // 转账并记录该地址
        token.transfer(msg.sender, amountAllowed);
        requestedAddress[msg.sender] = true;

        // 释放事件
        emit SendToken(msg.sender, amountAllowed);
    }
}