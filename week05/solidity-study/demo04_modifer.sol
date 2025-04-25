// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DemoModifier {

    address public owner;

    constructor(address _initOwner) {
        owner = _initOwner;
    }

    modifier onlyByOwner {
        require(msg.sender == owner);
        _; // 如果断言通过的话，继续运行函数主体；否则报错并revert交易
    }

    function changeOwner(address _newOwner) external onlyByOwner {
        owner = _newOwner;
    }
}