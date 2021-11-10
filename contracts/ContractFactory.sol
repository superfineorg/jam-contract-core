pragma solidity ^0.8.0;
import "./SimpleERC721.sol";
import "./SimpleERC20.sol";

contract ContractFactory {
    UniqueItem[] private uniqueItems;
    CoinToken[] private coinTokens;
    
    
    event CreateERC721(address indexed _owner, address contract_address);
    event CreateERC20(address indexed _owner, address contract_address);


    function createERC721(string memory _name, string memory _symbol) public {
        UniqueItem item = new UniqueItem(msg.sender, _name, _symbol);
        uniqueItems.push(item);
        emit CreateERC721(msg.sender, item.nftAddress());
    }
    
    function createERC20(uint256 capacity, uint256 initialSupply, uint8 d, string memory _name, string memory _symbol) public {
        CoinToken token = new CoinToken(msg.sender, capacity, initialSupply, d, _name, _symbol);
        coinTokens.push(token);
        emit CreateERC20(msg.sender, token.coinTokenAddress());
    }

    function getERC721(uint _index)
        public
        view
        returns (
            address nftAddress,
            address owner,
            string memory name,
            string memory symbol,
            uint256 totalSupply
        )
    {
        UniqueItem item = uniqueItems[_index];

        return (item.nftAddress(), item.owner(), item.name(), item.symbol(), item.totalSupply());
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