// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "./IERC20.sol";

contract AirDrop {
    /// @notice 失败转账记录
    ///
    mapping(address => uint) failTransferList;

    function getSum(uint256[] calldata _arr) internal pure returns(uint256 sum) {
        for (uint i=0; i<_arr.length; i++) {
            sum += _arr[i];
        }
    }

    /// @notice 向多个地址转账ERC20代币，使用前需要先授权
    ///
    /// @param _token 转账的ERC20代币地址
    /// @param _addresses 空投地址数组
    /// @param _amounts 代币数量数组（每个地址的空投数量）
    function multiTransferToken(
        address _token, address[] calldata _addresses, uint256[] calldata _amounts
    ) external {
        require(_addresses.length==_amounts.length);
        uint256 sum = getSum(_amounts);
        IERC20 token = IERC20(_token);
        // 校验发送者给这个空投合约授权的代币数量是否充足
        require(token.allowance(msg.sender, address(this))>=sum, "Need Approve ERC20 token");

        for (uint i=0; i<_addresses.length; i++) {
            token.transferFrom(msg.sender, _addresses[i], _amounts[i]);
        }
    }

    /// 向多个地址转账ETH
    function multiTransferETH(
        address payable[] calldata _addresses,
        uint256[] calldata _amounts
    ) external payable {
        require(_addresses.length==_amounts.length);
        uint256 sum = getSum(_amounts);
        require(msg.value==sum, "Transfer amount error");
        for (uint i=0; i<_addresses.length; i++) {
            // 这样写会有Dos攻击风险，具体参考 https://github.com/AmazingAng/WTF-Solidity/blob/main/S09_DoS/readme.md
            // _addresses[i].transfer(_amounts[i]);
            (bool success, ) = _addresses[i].call{value: _amounts[i]}("");
            if (!success) {
                failTransferList[_addresses[i]] += _amounts[i];
            }
        }
    }
}