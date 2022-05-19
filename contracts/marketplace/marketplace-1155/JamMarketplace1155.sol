// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "../JamMarketplaceHelpers.sol";

contract JamMarketplace1155 is JamMarketplaceHelpers {
    using Counters for Counters.Counter;

    // Represents an auction on an NFT
    struct Auction {
        uint256 auctionId;
        address nftAddress;
        uint256 tokenId;
        uint256 quantity;
        address seller;
        uint128 price;
        address currency;
        uint64 startedAt;
    }

    // The current ID of auction
    Counters.Counter private _currentAuctionId;

    // Mapping from auction ID to the auction info
    mapping(uint256 => Auction) private _auctions;

    event AuctionCreated(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 price,
        address seller
    );

    event AuctionUpdated(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 price,
        address seller
    );

    event AuctionSuccessful(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        uint256 price,
        address currency,
        address winner
    );

    event AuctionCancelled(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    );

    // Modifiers to check that inputs can be safely stored with a certain number of bits
    modifier canBeStoredWith64Bits(uint256 value) {
        require(
            value <= type(uint64).max,
            "JamMarketplace1155: cannot be stored with 64 bits"
        );
        _;
    }

    modifier canBeStoredWith128Bits(uint256 value) {
        require(
            value < type(uint128).max,
            "JamMarketplace1155: cannot be stored with 128 bits"
        );
        _;
    }

    constructor(address hubAddress, uint256 ownerCut_)
        JamMarketplaceHelpers(hubAddress, ownerCut_)
    {
        marketplaceId = keccak256("JAM_MARKETPLACE_1155");
    }

    function _isOnAuction(Auction storage auction)
        internal
        view
        returns (bool)
    {
        return (auction.startedAt > 0);
    }

    function _getNFTContract(address nftAddress)
        internal
        pure
        returns (IERC1155)
    {
        IERC1155 nftContract = IERC1155(nftAddress);
        return nftContract;
    }

    function _getCurrencyContract(address currency)
        internal
        pure
        returns (IERC20)
    {
        IERC20 currencyContract = IERC20(currency);
        return currencyContract;
    }

    function _ownsEnough(
        address nftAddress,
        address claimant,
        uint256 tokenId,
        uint256 quantity
    ) internal view returns (bool) {
        IERC1155 nftContract = _getNFTContract(nftAddress);
        return (nftContract.balanceOf(claimant, tokenId) >= quantity);
    }

    function _transfer(
        address nftAddress,
        address recipient,
        uint256 tokenId,
        uint256 quantity
    ) internal {
        IERC1155 nftContract = _getNFTContract(nftAddress);
        nftContract.safeTransferFrom(
            address(this),
            recipient,
            tokenId,
            quantity,
            abi.encodePacked("Transfer")
        );
    }

    function _addAuction(Auction memory auction) internal {
        _auctions[auction.auctionId] = auction;
        emit AuctionCreated(
            auction.auctionId,
            auction.nftAddress,
            auction.tokenId,
            auction.quantity,
            auction.currency,
            uint256(auction.price),
            auction.seller
        );
    }

    function _escrow(
        address nftAddress,
        address owner,
        uint256 tokenId,
        uint256 quantity
    ) internal {
        IERC1155 nftContract = _getNFTContract(nftAddress);
        nftContract.safeTransferFrom(
            owner,
            address(this),
            tokenId,
            quantity,
            abi.encodePacked("Transfer")
        );
    }

    /**
     * @dev Removes an auction from the list of open auctions.
     * @param auctionId - The ID of the auction to remove
     */
    function _removeAuction(uint256 auctionId) internal {
        delete _auctions[auctionId];
    }

    /**
     * @dev Cancels an auction unconditionally.
     */
    function _cancelAuction(uint256 auctionId, address recipient) internal {
        Auction memory auction = _auctions[auctionId];
        address nftAddress = auction.nftAddress;
        uint256 tokenId = auction.tokenId;
        uint256 quantity = auction.quantity;
        _removeAuction(auctionId);
        _transfer(nftAddress, recipient, tokenId, quantity);
        emit AuctionCancelled(auctionId, nftAddress, tokenId, quantity);
    }

    function getAuction(uint256 auctionId)
        external
        view
        returns (
            address nftAddress,
            uint256 tokenId,
            uint256 quantity,
            address seller,
            address currency,
            uint256 price,
            uint256 startedAt
        )
    {
        Auction storage auction = _auctions[auctionId];
        require(_isOnAuction(auction), "JamMarketplace1155: not on auction");
        return (
            auction.nftAddress,
            auction.tokenId,
            auction.quantity,
            auction.seller,
            auction.currency,
            auction.price,
            auction.startedAt
        );
    }

    /**
     * @dev Create an auction.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of the NFT.
     * @param quantity - The quantity of NFTs listed for sale.
     * @param price - Price of the NFTs to sell.
     * @param currency - Address of price need to be listed on.
     */
    function createAuction(
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        uint256 price,
        address currency
    ) external whenNotPaused canBeStoredWith128Bits(price) {
        address seller = msg.sender;
        require(
            _ownsEnough(nftAddress, seller, tokenId, quantity),
            "JamMarketplace1155: not enough NFTs to sell"
        );
        _escrow(nftAddress, seller, tokenId, quantity);

        uint256 auctionId = _currentAuctionId.current();
        Auction memory auction = Auction(
            auctionId,
            nftAddress,
            tokenId,
            quantity,
            seller,
            uint128(price),
            currency,
            uint64(block.timestamp)
        );
        _currentAuctionId.increment();
        _addAuction(auction);
    }

    /**
     * @dev Cancels an auction.
     * @param auctionId - The ID of the auction to cancel
     */
    function cancelAuction1155(uint256 auctionId) external override {
        Auction storage auction = _auctions[auctionId];
        require(_isOnAuction(auction), "JamMarketplace1155: not on auction");
        require(
            msg.sender == auction.seller ||
                msg.sender ==
                JamMarketplaceHub(_marketplaceHub).getMarketplace(
                    keccak256("JAM_P2P_TRADING_1155")
                ),
            "JamMarketplace1155: only seller or trading contract can cancel auction"
        );
        _cancelAuction(auctionId, msg.sender);
    }

    /**
     * @dev Cancels an auction when the contract is paused.
     * @param auctionId - The ID of the auction to cancel.
     */
    function cancelAuctionWhenPaused(uint256 auctionId)
        external
        whenPaused
        onlyOwner
    {
        Auction storage auction = _auctions[auctionId];
        require(_isOnAuction(auction), "JamMarketplace1155: not on auction");
        _cancelAuction(auctionId, auction.seller);
    }

    function buy(uint256 auctionId) external payable whenNotPaused {
        Auction memory auction = _auctions[auctionId];
        require(
            auction.currency == address(0),
            "JamMarketplace1155: can only be paid with native token"
        );
        _buy(auctionId, msg.value);
        _transfer(
            auction.nftAddress,
            msg.sender,
            auction.tokenId,
            auction.quantity
        );
    }

    function buyByERC20(uint256 auctionId, uint256 maxPrice)
        external
        whenNotPaused
        nonReentrant
    {
        Auction storage auction = _auctions[auctionId];
        require(
            auction.currency != address(0),
            "JamMarketplace1155: cannot buy these items with native token"
        );
        require(_isOnAuction(auction), "JamMarketplace1155: not on auction");
        uint256 quantity = auction.quantity;
        uint256 price = auction.price;
        address nftAddress = auction.nftAddress;
        uint256 tokenId = auction.tokenId;
        address currency = auction.currency;
        address seller = auction.seller;
        require(
            price <= maxPrice,
            "JamMarketplace1155: items price is too high"
        );

        IERC20 currencyContract = _getCurrencyContract(currency);

        uint256 allowance = currencyContract.allowance(
            msg.sender,
            address(this)
        );

        require(allowance >= price, "JamMarketplace1155: not enough allowance");

        bool success = currencyContract.transferFrom(
            msg.sender,
            address(this),
            price
        );
        require(success, "JamMarketplace1155: not enough balance");
        if (price > 0)
            _computeFeesAndPaySeller(nftAddress, seller, currency, price);
        _transfer(nftAddress, msg.sender, tokenId, quantity);
        _removeAuction(auctionId);
    }

    /**
     * @dev Computes the price and transfers winnings.
     * @notice Does NOT transfer ownership of token.
     */
    function _buy(uint256 auctionId, uint256 bidAmount)
        internal
        returns (uint256)
    {
        // Get a reference to the auction struct
        Auction storage auction = _auctions[auctionId];

        require(_isOnAuction(auction), "JamMarketplace1155: not on auction");
        uint256 quantity = auction.quantity;
        uint256 price = auction.price;
        address seller = auction.seller;
        address nftAddress = auction.nftAddress;
        uint256 tokenId = auction.tokenId;

        require(
            bidAmount >= price,
            "JamMarketplace1155: insufficient bid amount"
        );

        _removeAuction(auctionId);

        if (price > 0)
            _computeFeesAndPaySeller(nftAddress, seller, address(0), price);

        if (bidAmount > price) {
            uint256 bidExcess = bidAmount - price;
            (bool success, ) = payable(msg.sender).call{value: bidExcess}("");
            require(success, "JamMarketplace1155: return bid excess failed");
        }

        // Tell the world!
        emit AuctionSuccessful(
            auctionId,
            nftAddress,
            tokenId,
            quantity,
            price,
            address(0),
            msg.sender
        );

        return price;
    }
}
