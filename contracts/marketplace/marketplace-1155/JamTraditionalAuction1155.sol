/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC1155/utils/ERC1155Holder.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "../JamMarketplaceHelpers.sol";

contract JamTraditionalAuction1155 is JamMarketplaceHelpers, ERC1155Holder {
    using Counters for Counters.Counter;

    // The information of an auction
    struct Auction {
        uint256 auctionId;
        address nftAddress;
        uint256 tokenId;
        uint256 quantity;
        address seller;
        address currency;
        address highestBidder;
        uint256 highestBidAmount;
        uint256 endAt;
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
        uint256 startingPrice,
        uint256 endAt,
        address seller
    );

    event AuctionUpdated(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 startingPrice,
        uint256 endAt,
        address seller
    );

    event AuctionBidded(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 bidAmount,
        address bidder
    );

    event AuctionSuccessful(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 highestBidAmount,
        address winner
    );

    event AuctionCancelled(
        uint256 auctionId,
        address nftAddress,
        uint256 tokenId,
        uint256 quantity
    );

    modifier ownsEnough(
        address nftAddress,
        address claimant,
        uint256 tokenId,
        uint256 quantity
    ) {
        require(
            IERC1155(nftAddress).balanceOf(claimant, tokenId) >= quantity,
            "JamTraditionalAuction1155: not enough NFT"
        );
        _;
    }

    modifier onlySeller(uint256 auctionId) {
        require(
            msg.sender == _auctions[auctionId].seller ||
                msg.sender ==
                JamMarketplaceHub(_marketplaceHub).getMarketplace(
                    keccak256("JAM_P2P_TRADING_1155")
                ),
            "JamTraditionalAuction1155: sender is not seller nor trading marketplace"
        );
        _;
    }

    modifier isOnAuction(uint256 auctionId) {
        require(
            _auctions[auctionId].endAt > 0,
            "JamTraditionalAuction1155: auction not exists"
        );
        _;
    }

    constructor(address hubAddress, uint256 ownerCut_)
        JamMarketplaceHelpers(hubAddress, ownerCut_)
    {
        marketplaceId = keccak256("JAM_TRADITIONAL_AUCTION_1155");
    }

    receive() external payable {}

    /**
     * @dev Returns auction info for an NFT currently on auction.
     * @param auctionId - The ID of the auction to get information.
     */
    function getAuction(uint256 auctionId)
        external
        view
        isOnAuction(auctionId)
        returns (
            address nftAddress,
            uint256 tokenId,
            uint256 quantity,
            address seller,
            address currency,
            address highestBidder,
            uint256 highestBidAmount,
            uint256 endAt
        )
    {
        Auction memory auction = _auctions[auctionId];
        return (
            auction.nftAddress,
            auction.tokenId,
            auction.quantity,
            auction.seller,
            auction.currency,
            auction.highestBidder,
            auction.highestBidAmount,
            auction.endAt
        );
    }

    /**
     * @dev Creates and begins a new auction.
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction, sender must be owner.
     * @param quantity - The number of NFTs to sell.
     * @param currency - The address of an ERC20 currency that the owner wants to sell
     * @param startingPrice - Price of item (in wei) at beginning of auction.
     * @param endAt - The moment when this auction ends.
     */
    function createAuction(
        address nftAddress,
        uint256 tokenId,
        uint256 quantity,
        address currency,
        uint256 startingPrice,
        uint256 endAt
    )
        external
        whenNotPaused
        ownsEnough(nftAddress, msg.sender, tokenId, quantity)
    {
        require(
            endAt - block.timestamp >= 10 minutes,
            "JamTraditionalAuction1155: too short auction"
        );
        IERC1155(nftAddress).safeTransferFrom(
            msg.sender,
            address(this),
            tokenId,
            quantity,
            abi.encode("Create auction")
        );
        uint256 auctionId = _currentAuctionId.current();
        _auctions[auctionId] = Auction(
            auctionId,
            nftAddress,
            tokenId,
            quantity,
            msg.sender,
            currency,
            address(0),
            startingPrice,
            endAt
        );
        _currentAuctionId.increment();
        emit AuctionCreated(
            auctionId,
            nftAddress,
            tokenId,
            quantity,
            currency,
            startingPrice,
            endAt,
            msg.sender
        );
    }

    /**
     * @dev Updates an existent auction's information.
     * @param auctionId - The ID of the auction to update information.
     * @param currency - The address of an ERC20 currency that the owner wants to sell
     * @param startingPrice - Price of item (in wei) at beginning of auction.
     * @param endAt - The moment when this auction ends.
     */
    function updateAuction(
        uint256 auctionId,
        address currency,
        uint256 startingPrice,
        uint256 endAt
    ) external whenNotPaused isOnAuction(auctionId) onlySeller(auctionId) {
        Auction storage auction = _auctions[auctionId];
        require(
            endAt - block.timestamp >= 10 minutes,
            "JamTraditionalAuction1155: too short auction"
        );
        require(
            auction.highestBidder == address(0),
            "JamTraditionalAuction1155: cannot update auction after first bid"
        );
        auction.currency = currency;
        auction.highestBidAmount = startingPrice;
        auction.endAt = endAt;
        emit AuctionUpdated(
            auctionId,
            auction.nftAddress,
            auction.tokenId,
            auction.quantity,
            currency,
            startingPrice,
            endAt,
            msg.sender
        );
    }

    /**
     * @dev Cancels an auction that hasn't been won yet.
     * Returns the NFT to original owner.
     * @notice This is a state-modifying function that can be called while the contract is paused.
     * @param auctionId - The ID of the auction to cancel.
     */
    function cancelAuction1155(uint256 auctionId)
        external
        override
        isOnAuction(auctionId)
        onlySeller(auctionId)
    {
        Auction memory auction = _auctions[auctionId];
        require(
            auction.highestBidder == address(0),
            "JamTraditionalAuction1155: cannot cancel auction after first bid"
        );
        _cancelAuction(auctionId, msg.sender);
    }

    /**
     * @dev Cancels an auction when the contract is paused.
     * Only the owner may do this, and NFTs are returned to
     * the seller. This should only be used in emergencies.
     * @param auctionId - The ID of the auction to cancel.
     */
    function cancelAuctionWhenPaused(uint256 auctionId)
        external
        whenPaused
        onlyOwner
        isOnAuction(auctionId)
    {
        Auction memory auction = _auctions[auctionId];
        _cancelAuction(auctionId, auction.seller);
    }

    /**
     * @dev Bids on an open auction.
     * @param auctionId - The ID of the auction to bid on.
     * @param bidAmount - The amount a user wants to bid this NFT.
     */
    function bid(uint256 auctionId, uint256 bidAmount)
        external
        payable
        whenNotPaused
        nonReentrant
        isOnAuction(auctionId)
    {
        Auction storage auction = _auctions[auctionId];
        require(
            block.timestamp < auction.endAt,
            "JamTraditionalAuction1155: auction already ended"
        );
        if (auction.currency == address(0))
            require(
                bidAmount == msg.value,
                "JamTraditionalAuction1155: bid amount info mismatch"
            );
        else {
            (bool success, ) = payable(msg.sender).call{value: msg.value}("");
            require(success, "JamTraditionalAuction1155: return money failed");
        }
        require(
            bidAmount > auction.highestBidAmount,
            "JamTraditionalAuction1155: currently has higher bid"
        );

        // Return money to previous bidder
        if (auction.highestBidder != address(0))
            if (auction.currency == address(0)) {
                (bool success, ) = payable(auction.highestBidder).call{
                    value: auction.highestBidAmount
                }("");
                require(
                    success,
                    "JamTraditionalAuction1155: return money to previous bidder failed"
                );
            } else {
                IERC20 currencyContract = IERC20(auction.currency);
                bool success = currencyContract.transfer(
                    auction.highestBidder,
                    auction.highestBidAmount
                );
                require(
                    success,
                    "JamTraditionalAuction1155: return money to previous bidder failed"
                );
            }

        // Update auction info
        auction.highestBidder = msg.sender;
        auction.highestBidAmount = bidAmount;

        // Lock current bidder's money
        if (auction.currency != address(0))
            IERC20(auction.currency).transferFrom(
                msg.sender,
                address(this),
                bidAmount
            );

        emit AuctionBidded(
            auctionId,
            auction.nftAddress,
            auction.tokenId,
            auction.quantity,
            auction.currency,
            bidAmount,
            msg.sender
        );
    }

    /**
     * @dev Finalize the auction, return the NFT to the winner and return money to the seller.
     * @param auctionId - The ID of the auction to finalize.
     */
    function finalizeAuction(uint256 auctionId)
        external
        whenNotPaused
        isOnAuction(auctionId)
    {
        Auction memory auction = _auctions[auctionId];
        require(
            block.timestamp >= auction.endAt,
            "JamTraditionalAuction1155: auction not ends yet"
        );
        address nftAddress = auction.nftAddress;
        uint256 tokenId = auction.tokenId;
        uint256 quantity = auction.quantity;
        address winner = auction.highestBidder;
        address seller = auction.seller;
        address currency = auction.currency;
        uint256 price = auction.highestBidAmount;
        require(
            winner != address(0),
            "JamTraditionalAuction721: no-one bids on this auction"
        );
        require(
            msg.sender == winner || msg.sender == seller,
            "JamTraditionalAuction1155: sender is not winner nor seller"
        );

        //  Auction is successful, remove its info
        delete _auctions[auctionId];

        // Compute auctioneer cut and royalty cut then return proceeds to seller
        if (price > 0)
            _computeFeesAndPaySeller(nftAddress, seller, currency, price);

        // Give assets to winner
        IERC1155(nftAddress).safeTransferFrom(
            address(this),
            winner,
            tokenId,
            quantity,
            abi.encodePacked("Finalize auction")
        );

        emit AuctionSuccessful(
            auctionId,
            nftAddress,
            tokenId,
            quantity,
            currency,
            price,
            winner
        );
    }

    /** @dev Cancels an auction unconditionally. */
    function _cancelAuction(uint256 auctionId, address recipient) internal {
        Auction memory auction = _auctions[auctionId];
        address nftAddress = auction.nftAddress;
        uint256 tokenId = auction.tokenId;
        uint256 quantity = auction.quantity;
        delete _auctions[auctionId];
        IERC1155(nftAddress).safeTransferFrom(
            address(this),
            recipient,
            tokenId,
            quantity,
            abi.encode("Cancel auction")
        );
        emit AuctionCancelled(auctionId, nftAddress, tokenId, quantity);
    }
}
