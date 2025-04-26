// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract DemoSelector {
    event Log(bytes data);
    event SelectorEvent(bytes4);

    struct User {
        uint256 id;
        string name;
    }

    // Enum School
    enum School { SCHOOL1, SCHOOL2, SCHOOL3 }

    // "mint(address)"： 0x6a627842
    function mint(address /*to*/) external {
        emit Log(msg.data);
    }

    // 0x046dab16
    function mintSelector() external pure returns(bytes4) {
        return bytes4(keccak256("mint(address)"));
    }

    // nonParamSelector() ： 0x03817936
    function nonParamSelector() external returns(bytes4) {
        // 可以通过method.selector来获取到selector
        emit SelectorEvent(this.nonParamSelector.selector);
        return bytes4(keccak256("nonParamSelector()"));
    }

    // 0x3ec37834
    function elementaryParamSelector(uint256 param1, bool param2) external returns(bytes4) {
        emit SelectorEvent(this.elementaryParamSelector.selector);
        return bytes4(keccak256("elementaryParamSelector(uint256,bool)"));
    }

    // 0xead6b8bd
    function fixedSizeParamSelector(uint256[3] memory param1) external returns(bytes4) {
        emit SelectorEvent(this.fixedSizeParamSelector.selector);
        return bytes4(keccak256("fixedSizeParamSelector(uint256[3])"));
    }

    // 不定长餐数据
    function nonFixedSizeParamSelector(uint256[] memory param1,string memory param2) external returns(bytes4 selectorWithNonFixedSizeParam){
        emit SelectorEvent(this.nonFixedSizeParamSelector.selector);
        return bytes4(keccak256("nonFixedSizeParamSelector(uint256[],string)"));
    }

    // 映射类型：
    function mappingParamSelector(DemoSelector demo, User memory user, uint256[] memory count, School mySchool) external returns(bytes4){
        emit SelectorEvent(this.mappingParamSelector.selector);
        return bytes4(keccak256("mappingParamSelector(address,(uint256,bytes),uint256[],uint8)"));
    }

    // 用selector来调用函数
    function callWithSignature() external{
        (bool success, bytes memory data) = address(this).call(abi.encodeWithSelector(0x3ec37834, 1, 0));
        if (success) {
            emit Log(data);
        }
        
    }
}