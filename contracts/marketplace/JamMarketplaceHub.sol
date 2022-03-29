/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/access/Ownable.sol";

contract JamMarketplaceHub is Ownable {
    mapping(bytes32 => address) private _marketplaceById;
    mapping(address => bytes32) private _marketplaceIds;

    function getMarketplace(bytes32 id) external view returns (address) {
        return _marketplaceById[id];
    }

    function registerMarketplace(bytes32 id, address addr) external onlyOwner {
        _marketplaceById[id] = addr;
        _marketplaceIds[addr] = id;
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
