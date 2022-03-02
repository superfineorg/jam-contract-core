pragma solidity ^0.8.0;

/**
 * @dev Interface of gamejam .
 */
interface IOracle {
    function GetRate(address tokenAddress) external returns (uint256);
}
