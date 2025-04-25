// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract OtherContract {
    uint256 private _x;

    event Log(uint amount, uint gas);

    // 返回合约的ETH余额
    function getBalance() view public returns(uint) {
        return address(this).balance;
    }

    // 可以调整状态变量_x，并且可以往合约转ETH
    function setX(uint256 x) external payable {
        _x = x;
        // 如果转入了ETH，释放Log
        if (msg.value > 0) {
            emit Log(msg.value, gasleft());
        }
    }

    // 读取x
    function getX() external view returns(uint256) {
        return _x;
    }
}

contract CallContract {
    function callSetX(address _address, uint256 x) external {
        OtherContract(_address).setX(x);
    }

    function callGetX(OtherContract oc) external view returns(uint) {
        return oc.getX();
    }

    function callGetX2(address _address) external view returns(uint) {
        OtherContract oc = OtherContract(_address);
        return oc.getX();
    }

    function setXTransferETH(address _address, uint256 x) payable external {
        OtherContract(_address).setX{value: msg.value}(x);
    }
}

contract OtherContract2 {
    uint256 private _x;

    event Log(uint amount, uint gas);

    fallback() external payable {
        emit Log(0, 0);
    }

    receive() external payable {}

    // 返回合约的ETH余额
    function getBalance() view public returns(uint) {
        return address(this).balance;
    }

    // 可以调整状态变量_x，并且可以往合约转ETH
    function setX(uint256 x) external payable {
        _x = x;
        // 如果转入了ETH，释放Log
        if (msg.value > 0) {
            emit Log(msg.value, gasleft());
        }
    }

    // 读取x
    function getX() external view returns(uint256) {
        return _x;
    }
}

contract CallContract2 {
    // 定义Response事件，输出call返回的结果success和data
    event Response(bool success, bytes data);

    function callSetX(address payable _address, uint256 x) external payable {
        // 通过encodeWithSignature来声明调用的方法和参数
        (bool success, bytes memory data) = _address.call{value: msg.value}(abi.encodeWithSignature("setX(uint256)", x));
        if (!success) {
            emit Response(success, data);
        }
    }

    function callGetX(address payable _address) external payable returns(uint256) {
        (bool success, bytes memory data) = _address.call{value: msg.value}(abi.encodeWithSignature("getX()"));
        emit Response(success, data);
        // 通过decode来解析发返回的data
        return abi.decode(data, (uint256));
    }

    function callNonExist(address _address) external {
        // call不存在的函数
        (bool success, bytes memory data) = _address.call(abi.encodeWithSignature("foo(uint256)"));
        // call了不存在的foo函数。call仍能执行成功，并返回success，但其实调用的目标合约fallback函数。
        emit Response(success, data);
    }
}

contract ContractC {
    uint public num;

    address public sender;

    function setVal(uint _num) public payable {
        num = _num;
        sender = msg.sender;
    }
}

contract ContractB {
    uint public num;
    address public sender;

    // 通过call调用
    function callSetVal(address _address, uint _num) external payable  {
        (bool success, bytes memory data) = _address.call(abi.encodeWithSignature("setVal(uint256)", _num));
    }

    // 调用之后，由于上下文是和合约B的上下文，合约B中的sender和num会被修改，合约C的不会
    // 首先，合约B必须和目标合约C的变量存储布局必须相同 —— 即存在两个 public 变量且变量类型顺序为 uint256 和 address
    function delegateCallSetVal(address _address, uint _num) external payable  {
        (bool success, bytes memory data) = _address.delegatecall(abi.encodeWithSignature("setVal(uint256)", _num));
    }
}