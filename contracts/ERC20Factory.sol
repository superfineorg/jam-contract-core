pragma solidity ^0.8.0;
import "./SimpleERC20.sol";

contract ERC20Factory {
    CoinToken[] private coinTokens;


    event CreateERC20(address indexed _owner, string template_type, address contract_address, string name, string symbol);



    function createERC20(uint256 capacity, uint256 initialSupply, uint8 d, string memory _name, string memory _symbol) public {
        CoinToken token = new CoinToken(msg.sender, capacity, initialSupply, d, _name, _symbol);
        coinTokens.push(token);
        emit CreateERC20(msg.sender, "SimpleERC20-V1", token.coinTokenAddress(), _name, _symbol);
    }
    
    function getERC20(uint _index)
        public
        view
        returns (
            address coinTokenAddress,
            address owner,
            string memory name,
            string memory symbol,
            uint8 decimals,
            uint256 capacity
        )
    {
        CoinToken token = coinTokens[_index];

        return (token.coinTokenAddress(), token.owner(), token.name(), token.symbol(), token.decimals(), token.totalSupply());
    }
    
}