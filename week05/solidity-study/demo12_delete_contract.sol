// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DeleteContract {
    uint public value = 10;

    constructor() payable {}
    
    receive() external payable { }

    function deleteContract() external {
        // 调用selfdestruct自我销毁，并把剩余的eth转给msg.sender
        selfdestruct(payable(msg.sender));
    }

    function getBalance() external view returns(uint) {
        return address(this).balance;
    }
}

contract DeployContract {
    struct DemoResult {
        address _address;
        uint balance;
        uint value;
    }

    constructor() payable {}

    function getBalance() external view returns (uint) {
        return address(this).balance;
    }

    function createAndDestroyContract() external payable returns(DemoResult memory) {
        DeleteContract deleteContract = new DeleteContract{value: msg.value}();
        DemoResult memory dr = DemoResult({
            _address: address(deleteContract),
            balance: deleteContract.getBalance(),
            value: deleteContract.value()
        });
        deleteContract.deleteContract();
        return dr;
    }
}