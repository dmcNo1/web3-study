// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DemoVariable {

    // calldata中的数据不可变
    function fCalldata(uint[] calldata _x) external pure returns (uint[] calldata) {
        // _x[0] = 0; 这行会报错
        return _x;
    }

    string public s1;
    string public s2;

    // storage -> storage
    function copyS2S() external returns (string memory, string memory) {
        s1 = "hello";
        s2 = s1;
        s2 = "world";
        return (s1, s2);
    }

    // storage -> memory，这里带上view修饰符，仍然可以编译通过，说明这样不会修改到storage中的数据
    function copyS2M() external view returns (string memory, string memory) {
        // s1 = "hello";
        string memory s3 = s1;
        s3 = "copyS2M";
        return (s1, s3);
    }
}