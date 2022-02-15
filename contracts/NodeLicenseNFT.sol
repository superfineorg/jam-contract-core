// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract NodeLicenseNFT is ERC721Enumerable, Ownable {
    uint256 _tokenIds;
    string public tokenURIPrefix =
        "https://asset.gamejam.co/gamejam-nft/erc/721/node-license/";
    string public tokenURISuffix = ".json";

    /**
     * @dev Emitted when `tokenId` token is transferred from `from` to `to`.
     */
    event AwardItem(address indexed player, uint256 indexed tokenId);

    constructor() ERC721("Gamejam Node", "GJN") {}

    function mint(address player) public onlyOwner returns (uint256) {
        _tokenIds += 1;
        _mint(player, _tokenIds);
        return _tokenIds;
    }

    function changeTokenPrefix(string memory prefix) external onlyOwner {
        tokenURIPrefix = prefix;
    }

    function totalSupply() public view virtual override returns (uint256) {
        return _tokenIds;
    }

    function tokenURI(uint256 _tokenId)
        public
        view
        override
        returns (string memory)
    {
        bytes memory _tokenURIPrefixBytes = bytes(tokenURIPrefix);
        bytes memory _tokenURISuffixBytes = bytes(tokenURISuffix);
        uint256 _tmpTokenId = _tokenId;
        uint256 _length;

        do {
            _length++;
            _tmpTokenId /= 10;
        } while (_tmpTokenId > 0);

        bytes memory _tokenURIBytes = new bytes(
            _tokenURIPrefixBytes.length + _length + 5
        );
        uint256 _i = _tokenURIBytes.length - 6;

        _tmpTokenId = _tokenId;

        do {
            _tokenURIBytes[_i--] = bytes1(uint8(48 + (_tmpTokenId % 10)));
            _tmpTokenId /= 10;
        } while (_tmpTokenId > 0);

        for (_i = 0; _i < _tokenURIPrefixBytes.length; _i++) {
            _tokenURIBytes[_i] = _tokenURIPrefixBytes[_i];
        }

        for (_i = 0; _i < _tokenURISuffixBytes.length; _i++) {
            _tokenURIBytes[
                _tokenURIBytes.length + _i - 5
            ] = _tokenURISuffixBytes[_i];
        }

        return string(_tokenURIBytes);
    }

    function tokenOfOwner(address owner)
        public
        view
        returns (uint256[] memory)
    {
        uint256 totalToken = super.balanceOf(owner);
        uint256[] memory tokenIds = new uint256[](totalToken);
        for (uint256 i = 0; i < totalToken; i += 1) {
            uint256 tmp = super.tokenOfOwnerByIndex(owner, i);
            tokenIds[i] = tmp;
        }
        return tokenIds;
    }
}
