// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./SimpleERC721.sol";

contract ERC721Factory {
    UniqueItem[] private uniqueItems;

    event CreateERC721(address indexed _owner, string template_type,  address contract_address, string name, string symbol);

    function createERC721(string memory _name, string memory _symbol, string memory _tokenUri) public {
        UniqueItem item = new UniqueItem(msg.sender, _name, _symbol, _tokenUri);
        uniqueItems.push(item);
        emit CreateERC721(msg.sender, "SimpleERC721-V1",item.nftAddress(), _name, _symbol);
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
}