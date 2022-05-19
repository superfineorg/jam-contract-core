// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "../JamMarketplaceHelpers.sol";

contract JamMarketplace721 is JamMarketplaceHelpers, ERC721Holder {
    using SafeMath for uint256;

    // Represents an auction on an NFT
    struct Auction {
        address seller;
        uint128 price;
        address erc20Address; // erc20Address, address(0) if place a price with (wei)
        address nftAddress;
        uint256 tokenId;
        uint64 startedAt;
    }

    // Map from an NFT (NFT address + token ID) to its corresponding auction.
    mapping(address => mapping(uint256 => Auction)) private _auctions;

    event AuctionCreated(
        address indexed nftAddress,
        uint256 indexed tokenId,
        uint256 price,
        address seller,
        address erc20Address
    );

    event AuctionSuccessful(
        address indexed nftAddress,
        uint256 indexed tokenId,
        uint256 price,
        address winner
    );

    event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId);

    // Modifiers to check that inputs can be safely stored with a certain
    // number of bits. We use constants and multiple modifiers to save gas.
    modifier canBeStoredWith64Bits(uint256 _value) {
        require(
            _value <= type(uint64).max,
            "JamMarketplace721: cannot be stored with 64 bits"
        );
        _;
    }

    modifier canBeStoredWith128Bits(uint256 _value) {
        require(
            _value < type(uint128).max,
            "JamMarketplace721: cannot be stored with 128 bits"
        );
        _;
    }

    constructor(address hubAddress, uint256 ownerCut_)
        JamMarketplaceHelpers(hubAddress, ownerCut_)
    {
        marketplaceId = keccak256("JAM_MARKETPLACE_721");
    }

    function _isOnAuction(Auction storage _auction)
        internal
        view
        returns (bool)
    {
        return (_auction.startedAt > 0);
    }

    function _getNftContract(address _nftAddress)
        internal
        pure
        returns (IERC721)
    {
        IERC721 candidateContract = IERC721(_nftAddress);
        return candidateContract;
    }

    function _getERC20Contract(address _erc20Address)
        internal
        pure
        returns (IERC20)
    {
        IERC20 candidateContract = IERC20(_erc20Address);
        return candidateContract;
    }

    function _owns(
        address _nftAddress,
        address _claimant,
        uint256 _tokenId
    ) internal view returns (bool) {
        IERC721 _nftContract = _getNftContract(_nftAddress);
        return (_nftContract.ownerOf(_tokenId) == _claimant);
    }

    function _transfer(
        address _nftAddress,
        address _receiver,
        uint256 _tokenId
    ) internal {
        IERC721 _nftContract = _getNftContract(_nftAddress);

        // It will throw if transfer fails
        _nftContract.safeTransferFrom(address(this), _receiver, _tokenId);
    }

    function _addAuction(
        address _nftAddress,
        uint256 _tokenId,
        Auction memory _auction,
        address _seller
    ) internal {
        _auctions[_nftAddress][_tokenId] = _auction;

        emit AuctionCreated(
            _nftAddress,
            _tokenId,
            uint256(_auction.price),
            _seller,
            address(_auction.erc20Address)
        );
    }

    function _escrow(
        address _nftAddress,
        address _owner,
        uint256 _tokenId
    ) internal {
        IERC721 _nftContract = _getNftContract(_nftAddress);

        // It will throw if transfer fails
        _nftContract.safeTransferFrom(_owner, address(this), _tokenId);
    }

    /// @dev Removes an auction from the list of open auctions.
    /// @param _tokenId - ID of NFT on auction.
    function _removeAuction(address _nftAddress, uint256 _tokenId) internal {
        delete _auctions[_nftAddress][_tokenId];
    }

    /// @dev Cancels an auction unconditionally.
    function _cancelAuction(
        address _nftAddress,
        uint256 _tokenId,
        address _recipient
    ) internal {
        _removeAuction(_nftAddress, _tokenId);
        _transfer(_nftAddress, _recipient, _tokenId);
        emit AuctionCancelled(_nftAddress, _tokenId);
    }

    function getAuction(address _nftAddress, uint256 _tokenId)
        external
        view
        returns (
            address seller,
            uint256 price,
            address erc20Address,
            uint256 startedAt
        )
    {
        Auction storage _auction = _auctions[_nftAddress][_tokenId];
        require(_isOnAuction(_auction), "JamMarketplace721: not on auction");
        return (
            _auction.seller,
            _auction.price,
            _auction.erc20Address,
            _auction.startedAt
        );
    }

    /// @dev Create an auction.
    /// @param _nftAddress - Address of the NFT.
    /// @param _tokenId - ID of the NFT.
    /// @param _price - Price of erc20 address.
    /// @param _erc20Address - Address of price need to be list on.
    function createAuction(
        address _nftAddress,
        uint256 _tokenId,
        uint256 _price,
        address _erc20Address
    ) external whenNotPaused canBeStoredWith128Bits(_price) {
        address _seller = msg.sender;
        require(
            _owns(_nftAddress, _seller, _tokenId),
            "JamMarketplace721: only owner can create auction"
        );
        _escrow(_nftAddress, _seller, _tokenId);

        Auction memory _auction = Auction(
            _seller,
            uint128(_price),
            _erc20Address,
            _nftAddress,
            _tokenId,
            uint64(block.timestamp)
        );
        _addAuction(_nftAddress, _tokenId, _auction, _seller);
    }

    /// @dev Cancels an auction.
    /// @param _nftAddress - Address of the NFT.
    /// @param _tokenId - ID of the NFT on auction to cancel.
    function cancelAuction721(address _nftAddress, uint256 _tokenId)
        external
        override
    {
        Auction storage _auction = _auctions[_nftAddress][_tokenId];
        require(_isOnAuction(_auction), "JamMarketplace721: not on auction");
        require(
            msg.sender == _auction.seller ||
                msg.sender ==
                JamMarketplaceHub(_marketplaceHub).getMarketplace(
                    keccak256("JAM_P2P_TRADING_721")
                ),
            "JamMarketplace721: only seller can cancel auction"
        );
        _cancelAuction(_nftAddress, _tokenId, msg.sender);
    }

    /// @dev Cancels an auction when the contract is paused.
    ///  Only the owner may do this, and NFTs are returned to
    ///  the seller. This should only be used in emergencies.
    /// @param _nftAddress - Address of the NFT.
    /// @param _tokenId - ID of the NFT on auction to cancel.
    function cancelAuctionWhenPaused(address _nftAddress, uint256 _tokenId)
        external
        whenPaused
        onlyOwner
    {
        Auction storage _auction = _auctions[_nftAddress][_tokenId];
        require(_isOnAuction(_auction), "JamMarketplace721: not on auction");
        _cancelAuction(_nftAddress, _tokenId, _auction.seller);
    }

    function buy(address _nftAddress, uint256 _tokenId)
        external
        payable
        whenNotPaused
    {
        Auction memory auction = _auctions[_nftAddress][_tokenId];
        require(
            auction.erc20Address == address(0),
            "JamMarketplace721: can only be paid with native token"
        );
        // _bid will throw if the bid or funds transfer fails
        _buy(_nftAddress, _tokenId, msg.value);
        _transfer(_nftAddress, msg.sender, _tokenId);
    }

    function buyByERC20(
        address _nftAddress,
        uint256 _tokenId,
        uint256 _maxPrice
    ) external whenNotPaused nonReentrant {
        Auction storage _auction = _auctions[_nftAddress][_tokenId];
        require(
            _auction.erc20Address != address(0),
            "JamMarketplace721: cannot buy this item with native token"
        );
        require(_isOnAuction(_auction), "JamMarketplace721: not on auction");
        uint256 _price = _auction.price;
        require(
            _price <= _maxPrice,
            "JamMarketplace721: item price is too high"
        );
        address _erc20Address = _auction.erc20Address;

        IERC20 erc20Contract = _getERC20Contract(_erc20Address);

        uint256 _allowance = erc20Contract.allowance(msg.sender, address(this));

        require(
            _allowance >= _price,
            "JamMarketplace721: not enough allowance"
        );

        address _seller = _auction.seller;
        bool success = erc20Contract.transferFrom(
            msg.sender,
            address(this),
            _price
        );
        require(success, "JamMarketplace721: not enough balance");
        if (_price > 0)
            _computeFeesAndPaySeller(
                _nftAddress,
                _seller,
                _erc20Address,
                _price
            );
        _transfer(_nftAddress, msg.sender, _tokenId);
        _removeAuction(_nftAddress, _tokenId);
    }

    /// @dev Computes the price and transfers winnings.
    /// Does NOT transfer ownership of token.
    function _buy(
        address _nftAddress,
        uint256 _tokenId,
        uint256 _bidAmount
    ) internal returns (uint256) {
        // Get a reference to the auction struct
        Auction storage _auction = _auctions[_nftAddress][_tokenId];

        require(_isOnAuction(_auction), "JamMarketplace721: not on auction");
        uint256 _price = _auction.price;

        require(
            _bidAmount >= _price,
            "JamMarketplace721: bid amount cannot be less than price"
        );

        address _seller = _auction.seller;

        _removeAuction(_nftAddress, _tokenId);

        if (_price > 0)
            _computeFeesAndPaySeller(_nftAddress, _seller, address(0), _price);

        if (_bidAmount > _price) {
            // Calculate any excess funds included with the bid. If the excess
            // is anything worth worrying about, transfer it back to bidder.
            // NOTE: We checked above that the bid amount is greater than or
            // equal to the price so this cannot underflow.
            uint256 _bidExcess = _bidAmount - _price;

            // Return the funds. Similar to the previous transfer, this is
            // not susceptible to a re-entry attack because the auction is
            // removed before any transfers occur.
            payable(msg.sender).transfer(_bidExcess);
        }

        // Tell the world!
        emit AuctionSuccessful(_nftAddress, _tokenId, _price, msg.sender);

        return _price;
    }
}
