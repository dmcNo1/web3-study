// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

contract Pair {
    address factory;    // 工厂合约地址
    address token0; // 代币1
    address token1; // 代币2

    constructor() payable {
        factory = msg.sender;
    }

    function initialize(address _token0, address _token1) external {
        require(msg.sender == factory, "UniswapV2: FORBIDDEN");  // 只有工厂才能调用initialize
        token0 = _token0;
        token1 = _token1;
    }
}

contract PairFactory {
    // 通过两个代币地址查Pair地址，A=>B，指向支持的合约地址
    mapping(address => mapping(address => address)) public getPair;
    // 保存所有Pair地址
    address[] public allPairs;

    function createPair(address tokenA, address tokenB) external returns (address) {
        // 创建新的合约
        Pair pair = new Pair();
        // 调用合约的initialize方法
        pair.initialize(tokenA, tokenB);
        // 更新地址map
        address pairAddr = address(pair);
        allPairs.push(pairAddr);
        getPair[tokenA][tokenB] = pairAddr;
        getPair[tokenB][tokenA] = pairAddr;
        return pairAddr;
    }
}

contract Pair2 {
    address public factory;
    address public token0;
    address public token1;

    constructor() payable {
        factory = msg.sender;
    }

    function initialize(address _token0, address _token1) external {
        require(msg.sender == factory, "UniswapV2: FORBIDDEN");  // 只有工厂才能调用initialize
        token0 = _token0;
        token1 = _token1;
    }
}

contract PairFactory2 {
    mapping(address => mapping(address => address)) public getPair; // 通过两个代币地址查Pair地址
    address[] public allPairs; // 保存所有Pair地址

    function createPair(address tokenA, address tokenB) external returns(address) {
        require(tokenA!=tokenB, "IDENTICAL_ADDRESSES");
        // 用tokenA和tokenB地址计算salt
        (address token0, address token1) = tokenA < tokenB ? (tokenA, tokenB): (tokenB, tokenA);    // 将tokenA和tokenB按大小排序
        bytes32 salt = keccak256(abi.encodePacked(token0, token1));
        // 用create2部署新合约
        Pair2 pair = new Pair2{salt: salt}();
        // 更新地址map
        address pairAddr = address(pair);
        allPairs.push(pairAddr);
        getPair[tokenA][tokenB] = pairAddr;
        getPair[tokenB][tokenA] = pairAddr;
        return pairAddr;
    }

    function calculateAddress(address tokenA, address tokenB) public view returns(address) {
        require(tokenA != tokenB, "IDENTICAL_ADDRESSES"); //避免tokenA和tokenB相同产生的冲突
        // 计算用tokenA和tokenB地址计算salt
        (address token0, address token1) = tokenA < tokenB ? (tokenA, tokenB) : (tokenB, tokenA); //将tokenA和tokenB按大小排序
        bytes32 salt = keccak256(abi.encodePacked(token0, token1));
        // 计算合约地址方法 hash()
        address predictedAddress = address(uint160(uint(keccak256(abi.encodePacked(
            bytes1(0xff),
            address(this),  // 这里使用了this，所以函数必须声明为view，不能是pure
            salt,
            // 如果构造函数需要传参，这里计算的时候也要带上
            // keccak256(abi.encodePacked(type(Pair2).creationCode, abi.encode(address(this))))
            // 在Solidity中，type关键字主要用于获取与特定类型相关的元数据或接口信息。
            keccak256(type(Pair2).creationCode)
            )))));
        return predictedAddress;
    }
}
