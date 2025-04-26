// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract OnlyEven {
    constructor(uint a) {
        require(a!=0, "individual number");
        assert(a != 1);
    }

    function onlyEven(uint256 b) external pure returns(bool) {
        // 输入奇数时revert
        require(b % 2 == 0, "Ups! Reverting");
        return true;
    }
}

contract TryCatch {
    // 成功event
    event SuccessEvent();
    // 失败event
    event CatchEvent(string message);
    event CatchByte(bytes data);

    OnlyEven onlyEven;

    constructor() {
        onlyEven = new OnlyEven(2);
    }

    function execute(uint amount) external returns(bool success) {
        try onlyEven.onlyEven(amount) returns(bool _success) {
            // call成功
            emit SuccessEvent();
            return _success;
        } catch Error(string memory errMsg) {
            emit CatchEvent(errMsg);
        }
    }

    function executeNew(uint a) external returns (bool success) {
        try new OnlyEven(a) returns(OnlyEven _onlyEven) {
            // call成功
            emit SuccessEvent();
            success = _onlyEven.onlyEven(a);
        } catch Error(string memory reason) {
            emit CatchEvent(reason);
        } catch (bytes memory reason) {
            emit CatchByte(reason);
        }
    }
}