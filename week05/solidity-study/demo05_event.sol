// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DemoEvent {

    mapping(address => uint256) _balance;

    event Transfer(address indexed from, address indexed to, uint256 value);

    function _transfer(address from, address to, uint256 value)  external {
        // 给一个初始值
        _balance[from] = 10000000;
        _balance[from] -= value;
        _balance[to] += value;

        // 触发时间
        emit Transfer(from, to, value);
    }
}