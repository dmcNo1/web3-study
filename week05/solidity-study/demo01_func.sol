// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DemoFunction {
    uint256 public number = 5;

    // pure修饰的函数既不能读取也不能操作内部变量
    function addPure(uint256 _number) external pure returns(uint256) {
        return _number + 1;
    }

    // view只能读取，不能更改
    function addView() external view returns(uint256) {
        return number + 1;
    }

    // 默认的话，既可以读取又可以更改
    function add() external returns (uint256) {
        number += 1;
        return number;
    }

    // internal修饰的函数只能内部访问，部署之后看不到这个方法
    function subInternal() internal returns(uint256) {
        number -= 1;
        return number;
    }

    // external修饰的函数可以外部访问，也可以调用内部方法
    function subExternal() external returns(uint256) {
        return subInternal();
    }

    // payable，能够给合约支付eth的函数
    function minusPayable() external payable returns (uint256) {
        subInternal();
        // 获取合约的地址，并且返回合约的余额
        return address(this).balance;
    }

    // 可以返回多个值，并且可以声明返回变量名，向Go一样
    function getReturnValues() internal pure returns (uint256 val1, uint256 val2) {
        val1 = 1;
        val2 = 2;
    }

    function getReturnValuesExternal() external pure returns (uint256) {
        uint256 val2;
        // 不要的参数直接空着就好
        (, val2) = getReturnValues();
        return val2;
    }
}