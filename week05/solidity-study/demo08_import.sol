// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import {God} from "./demo06_extends.sol";
// 通过网址引用
import 'https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/Address.sol';
// 引用OpenZeppelin合约
import '@openzeppelin/contracts/access/Ownable.sol';

contract DemoImport {
    God god = new God();

    using Address for address;
    
    function testGod() external {
        god.foo();
    }
}