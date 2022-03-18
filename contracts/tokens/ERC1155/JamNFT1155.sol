/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/token/ERC1155/presets/ERC1155PresetMinterPauser.sol";

contract JamNFT1155 is ERC1155PresetMinterPauser {
    using Strings for uint256;

    struct TokenInfo {
        uint256 tokenId;
        uint256 quantity;
        string uri;
    }

    uint256[] private _mintedTokenIds;
    mapping(uint256 => bool) private _isTokenIdMinted;

    constructor(string memory uri_) ERC1155PresetMinterPauser(uri_) {}

    function uri(uint256 tokenId) public view override returns (string memory) {
        require(
            _isTokenIdMinted[tokenId],
            "JamNFT1155: URI query for non-existent token"
        );
        return
            string(
                abi.encodePacked(
                    super.uri(tokenId),
                    Strings.toString(tokenId),
                    ".json"
                )
            );
    }

    function getOwnedTokens(address user)
        external
        view
        returns (TokenInfo[] memory)
    {
        // Get the number of owned ERC1155 NFTs
        uint256 numOwnedNFTs = 0;
        for (uint256 i = 0; i < _mintedTokenIds.length; i++)
            if (balanceOf(user, _mintedTokenIds[i]) > 0) numOwnedNFTs++;

        // Query all owned ERC1155 NFTs
        TokenInfo[] memory ownedNFTs = new TokenInfo[](numOwnedNFTs);
        uint256 nftCount = 0;
        for (uint256 j = 0; j < _mintedTokenIds.length; j++) {
            uint256 tokenId = _mintedTokenIds[j];
            if (balanceOf(user, tokenId) > 0) {
                ownedNFTs[nftCount] = TokenInfo(
                    tokenId,
                    balanceOf(user, tokenId),
                    uri(tokenId)
                );
                nftCount++;
            }
        }
        return ownedNFTs;
    }

    function setBaseTokenURI(string memory baseTokenURI_) external {
        require(
            hasRole(DEFAULT_ADMIN_ROLE, _msgSender()),
            "JamNFT1155: caller is not an admin"
        );
        _setURI(baseTokenURI_);
    }

    function _beforeTokenTransfer(
        address operator,
        address from,
        address to,
        uint256[] memory tokenIds,
        uint256[] memory quantities,
        bytes memory data
    ) internal virtual override {
        require(
            tokenIds.length == quantities.length,
            "JamNFT1155: lengths mismatch"
        );
        if (from == address(0))
            for (uint256 i = 0; i < tokenIds.length; i++)
                if (!_isTokenIdMinted[tokenIds[i]]) {
                    _isTokenIdMinted[tokenIds[i]] = true;
                    _mintedTokenIds.push(tokenIds[i]);
                }
        super._beforeTokenTransfer(
            operator,
            from,
            to,
            tokenIds,
            quantities,
            data
        );
    }
}
