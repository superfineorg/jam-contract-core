/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "./JamNFT1155.sol";

contract GuardianOfGloryItems is JamNFT1155 {
    constructor(
        address owner,
        address minter,
        string memory uri
    ) JamNFT1155(uri) {
        _setupRole(DEFAULT_ADMIN_ROLE, owner);
        _setupRole(MINTER_ROLE, minter);
    }
}
