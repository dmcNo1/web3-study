// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DemoETH {

    event receivedCalled(address sender, uint256 value);
    event fallbackCalled(address sender, uint256 value, bytes Data);

    receive() external payable {
        emit receivedCalled(msg.sender, msg.value);
    }

    fallback() external payable {
        emit fallbackCalled(msg.sender, msg.value, msg.data);
    }
}

contract ReceiveETH {
    // 接收eth事件，记录amount和gas
    event Log(uint amount, uint gas);

    receive() external payable {
        emit Log(msg.value, gasleft());
    }

    // 返回合约的ETH余额
    function gasBalance() view external returns(uint) {
        return address(this).balance;
    }
}

contract SendETH {
    // 构造函数，payable使得部署的时候可以转入ETH
    constructor() payable {}

    receive() external payable {}

    // 用transfer()发送ETH
    function transferETH(address payable _to, uint256 amount) payable external {
        _to.transfer(amount);
    }

    // 使用send发送失败的error
    error SendFailed();

    // send()发送ETH
    function sendETH(address payable _to, uint256 amount) payable external {
        bool success = _to.send(amount);
        if (!success) {
            revert SendFailed();
        }
    }

    // 使用call发送失败的error
    error CallFailed();

    // call()发送ETH
    function callETH(address payable _to, uint256 amount) payable external {
        // 处理下call的返回值，如果失败，revert交易并发送error
        (bool success, ) = _to.call{value: amount}("");
        if (!success) {
            revert CallFailed();
        }
    }
}