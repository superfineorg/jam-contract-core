/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";

contract JamMarketplaceHub is Ownable {
    mapping(bytes32 => address) private _marketplaceById;
    mapping(address => bytes32) private _marketplaceIds;

    function isMarketplace(address addr) external view returns (bool) {
        return _marketplaceIds[addr] != 0x0;
    }

    function getMarketplace(bytes32 id) external view returns (address) {
        return _marketplaceById[id];
    }

    function registerMarketplace(bytes32 id) external {
        require(
            Ownable(msg.sender).owner() == owner(),
            "JamMarketplaceHub: invalid caller contract"
        );
        require(id != 0x0, "JamMarketplaceHub: invalid marketplace id");
        _marketplaceById[id] = msg.sender;
        _marketplaceIds[msg.sender] = id;
    }

    function unregisterMarketplace(bytes32 id) external onlyOwner {
        delete _marketplaceIds[_marketplaceById[id]];
        delete _marketplaceById[id];
    }

    function unregisterMarketplace(address addr) external onlyOwner {
        delete _marketplaceById[_marketplaceIds[addr]];
        delete _marketplaceIds[addr];
    }
}
