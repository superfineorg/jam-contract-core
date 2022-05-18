/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "./JamMarketplaceHelpers.sol";

contract JamTraditionalAuction is JamMarketplaceHelpers {
    // The information of an auction
    struct Auction {
        address seller;
        address currency;
        address highestBidder;
        uint256 highestBidAmount;
        uint256 endAt;
    }

    // Mapping from (NFT address + token ID) to the NFT's auction.
    mapping(address => mapping(uint256 => Auction)) private _auctions;

    event AuctionCreated(
        address indexed nftAddress,
        uint256 indexed tokenId,
        address currency,
        uint256 startingPrice,
        uint256 endAt,
        address seller
    );

    event AuctionUpdated(
        address indexed nftAddress,
        uint256 indexed tokenId,
        address currency,
        uint256 startingPrice,
        uint256 endAt,
        address seller
    );

    event AuctionBidded(
        address indexed nftAddress,
        uint256 indexed tokenId,
        address currency,
        uint256 bidAmount,
        address bidder
    );

    event AuctionSuccessful(
        address indexed nftAddress,
        uint256 indexed tokenId,
        address currency,
        uint256 highestBidAmount,
        address winner
    );

    event AuctionCancelled(address indexed nftAddress, uint256 indexed tokenId);

    modifier owns(
        address nftAddress,
        address claimant,
        uint256 tokenId
    ) {
        require(
            IERC721(nftAddress).ownerOf(tokenId) == claimant,
            "JamTraditionalAuction: sender is not owner of NFT"
        );
        _;
    }

    modifier onlySeller(address nftAddress, uint256 tokenId) {
        require(
            msg.sender == _auctions[nftAddress][tokenId].seller ||
                msg.sender ==
                JamMarketplaceHub(_marketplaceHub).getMarketplace(
                    keccak256("JAM_P2P_TRADING")
                ),
            "JamTraditionalAuction: sender is not seller"
        );
        _;
    }

    modifier isOnAuction(address nftAddress, uint256 tokenId) {
        require(
            _auctions[nftAddress][tokenId].endAt > 0,
            "JamTraditionalAuction: auction not exists"
        );
        _;
    }

    constructor(address hubAddress, uint256 ownerCut_)
        JamMarketplaceHelpers(hubAddress, ownerCut_)
    {
        marketplaceId = keccak256("JAM_TRADITIONAL_AUCTION");
    }

    receive() external payable {}

    /**
     * @dev Returns auction info for an NFT currently on auction.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of NFT on auction.
     */
    function getAuction(address nftAddress, uint256 tokenId)
        external
        view
        isOnAuction(nftAddress, tokenId)
        returns (
            address seller,
            address currency,
            address highestBidder,
            uint256 highestBidAmount,
            uint256 endAt
        )
    {
        Auction memory auction = _auctions[nftAddress][tokenId];
        return (
            auction.seller,
            auction.currency,
            auction.highestBidder,
            auction.highestBidAmount,
            auction.endAt
        );
    }

    /**
     * @dev The auction is cancelable if no one has bidded so far
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction
     */
    function isAuctionCancelable(address nftAddress, uint256 tokenId)
        external
        view
        override
        returns (bool)
    {
        return _auctions[nftAddress][tokenId].highestBidder == address(0);
    }

    /**
     * @dev Creates and begins a new auction.
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction, sender must be owner.
     * @param currency - The address of an ERC20 currency that the owner wants to sell
     * @param startingPrice - Price of item (in wei) at beginning of auction.
     * @param endAt - The moment when this auction ends.
     */
    function createAuction(
        address nftAddress,
        uint256 tokenId,
        address currency,
        uint256 startingPrice,
        uint256 endAt
    ) external whenNotPaused owns(nftAddress, msg.sender, tokenId) {
        require(
            endAt - block.timestamp >= 10 minutes,
            "JamTraditionalAuction: too short auction"
        );
        IERC721(nftAddress).transferFrom(msg.sender, address(this), tokenId);
        _auctions[nftAddress][tokenId] = Auction(
            msg.sender,
            currency,
            address(0),
            startingPrice,
            endAt
        );
        emit AuctionCreated(
            nftAddress,
            tokenId,
            currency,
            startingPrice,
            endAt,
            msg.sender
        );
    }

    /**
     * @dev Updates an existent auction's information.
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction, sender must be owner.
     * @param currency - The address of an ERC20 currency that the owner wants to sell
     * @param startingPrice - Price of item (in wei) at beginning of auction.
     * @param endAt - The moment when this auction ends.
     */
    function updateAuction(
        address nftAddress,
        uint256 tokenId,
        address currency,
        uint256 startingPrice,
        uint256 endAt
    )
        external
        whenNotPaused
        isOnAuction(nftAddress, tokenId)
        onlySeller(nftAddress, tokenId)
    {
        Auction storage auction = _auctions[nftAddress][tokenId];
        require(
            endAt - block.timestamp >= 10 minutes,
            "JamTraditionalAuction: too short auction"
        );
        require(
            auction.highestBidder == address(0),
            "JamTraditionalAuction: cannot update auction after first bid"
        );
        auction.currency = currency;
        auction.highestBidAmount = startingPrice;
        auction.endAt = endAt;
        emit AuctionUpdated(
            nftAddress,
            tokenId,
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
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of token on auction
     */
    function cancelAuction(address nftAddress, uint256 tokenId)
        external
        override
        isOnAuction(nftAddress, tokenId)
        onlySeller(nftAddress, tokenId)
    {
        Auction memory auction = _auctions[nftAddress][tokenId];
        require(
            auction.highestBidder == address(0),
            "JamTraditionalAuction: cannot cancel auction after first bid"
        );
        _cancelAuction(nftAddress, tokenId, msg.sender);
    }

    /**
     * @dev Cancels an auction when the contract is paused.
     * Only the owner may do this, and NFTs are returned to
     * the seller. This should only be used in emergencies.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of the NFT on auction to cancel.
     */
    function cancelAuctionWhenPaused(address nftAddress, uint256 tokenId)
        external
        whenPaused
        onlyOwner
        isOnAuction(nftAddress, tokenId)
    {
        Auction memory auction = _auctions[nftAddress][tokenId];
        _cancelAuction(nftAddress, tokenId, auction.seller);
    }

    /**
     * @dev Bids on an open auction.
     * @param nftAddress - address of a deployed contract implementing the Nonfungible Interface.
     * @param tokenId - ID of token to bid on.
     * @param bidAmount - The amount a user wants to bid this NFT
     */
    function bid(
        address nftAddress,
        uint256 tokenId,
        uint256 bidAmount
    )
        external
        payable
        whenNotPaused
        nonReentrant
        isOnAuction(nftAddress, tokenId)
    {
        Auction storage auction = _auctions[nftAddress][tokenId];
        require(
            block.timestamp < auction.endAt,
            "JamTraditionalAuction: auction already ended"
        );
        if (auction.currency == address(0))
            require(
                bidAmount == msg.value,
                "JamTraditionalAuction: bid amount info mismatch"
            );
        else {
            (bool success, ) = payable(msg.sender).call{value: msg.value}("");
            require(success, "JamTraditionalAuction: return money failed");
        }
        require(
            bidAmount > auction.highestBidAmount,
            "JamTraditionalAuction: currently has higher bid"
        );

        // Return money to previous bidder
        if (auction.highestBidder != address(0))
            if (auction.currency == address(0)) {
                (bool success, ) = payable(auction.highestBidder).call{
                    value: auction.highestBidAmount
                }("");
                require(
                    success,
                    "JamTraditionalAuction: return money to previous bidder failed"
                );
            } else {
                IERC20 currencyContract = IERC20(auction.currency);
                bool success = currencyContract.transfer(
                    auction.highestBidder,
                    auction.highestBidAmount
                );
                require(
                    success,
                    "JamTraditionalAuction: return money to previous bidder failed"
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
            nftAddress,
            tokenId,
            auction.currency,
            bidAmount,
            msg.sender
        );
    }

    /**
     * @dev Finalize the auction, return the NFT to the winner and return money to the seller.
     * @param nftAddress - address of a deployed contract implementing the Nonfungible Interface.
     * @param tokenId - ID of token to bid on.
     */
    function finalizeAuction(address nftAddress, uint256 tokenId)
        external
        whenNotPaused
        isOnAuction(nftAddress, tokenId)
    {
        Auction memory auction = _auctions[nftAddress][tokenId];
        require(
            block.timestamp >= auction.endAt,
            "JamTraditionalAuction: auction not ends yet"
        );
        address winner = auction.highestBidder;
        address seller = auction.seller;
        address currency = auction.currency;
        uint256 price = auction.highestBidAmount;
        require(
            msg.sender == winner || msg.sender == seller,
            "JamTraditionalAuction: sender is not winner nor seller"
        );

        //  Auction is successful, remove its info
        delete _auctions[nftAddress][tokenId];

        // Compute auctioneer cut and royalty cut then return proceeds to seller
        if (price > 0) _handleMoney(nftAddress, seller, currency, price);

        // Give assets to winner
        IERC721(nftAddress).transferFrom(address(this), winner, tokenId);

        emit AuctionSuccessful(nftAddress, tokenId, currency, price, winner);
    }

    /** @dev Cancels an auction unconditionally. */
    function _cancelAuction(
        address nftAddress,
        uint256 tokenId,
        address recipient
    ) internal {
        delete _auctions[nftAddress][tokenId];
        IERC721(nftAddress).transferFrom(address(this), recipient, tokenId);
        emit AuctionCancelled(nftAddress, tokenId);
    }
}
