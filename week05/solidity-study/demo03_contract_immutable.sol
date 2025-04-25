// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DemoConstantImmutable {
    uint256 constant public CONSTANT_NUM = 10;
    uint256 immutable public IMMUTABLE_NUM = 20;
    address immutable public IMMUTABLE_ADDRESS;

    constructor() {
        IMMUTABLE_ADDRESS = address(this);
    }
}