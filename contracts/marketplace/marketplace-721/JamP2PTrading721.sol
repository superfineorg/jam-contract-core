/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC721/utils/ERC721Holder.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "../JamMarketplaceHelpers.sol";

contract JamP2PTrading721 is JamMarketplaceHelpers, ERC721Holder {
    struct Offer {
        address offeror;
        address nftAddress;
        uint256 tokenId;
        address currency;
        uint256 amount;
    }

    // Mapping from an NFT (NFT address + token ID) to its pending offers
    mapping(address => mapping(uint256 => Offer[])) private _offersFor;

    // Mapping from a user address to the offers he has made so far
    mapping(address => Offer[]) private _offersOf;

    event OfferCreated(
        address offeror,
        address nftAddress,
        uint256 tokenId,
        address currency,
        uint256 amount
    );

    event OfferUpdated(
        address offeror,
        address nftAddress,
        uint256 tokenId,
        address currency,
        uint256 amount
    );

    event OfferAccepted(
        address offeror,
        address accepter,
        address nftAddress,
        uint256 tokenId,
        address currency,
        uint256 amount
    );

    event OfferCancelled(address offeror, address nftAddress, uint256 tokenId);

    constructor(address hubAddress, uint256 ownerCut_)
        JamMarketplaceHelpers(hubAddress, ownerCut_)
    {
        marketplaceId = keccak256("JAM_P2P_TRADING_721");
    }

    function getAllOffersFor(address nftAddress, uint256 tokenId)
        external
        view
        returns (Offer[] memory)
    {
        return _offersFor[nftAddress][tokenId];
    }

    function getAcceptableOffersFor(
        address nftAddress,
        uint256 tokenId,
        address user
    ) external view returns (Offer[] memory) {
        Offer[] memory allOffers = _offersFor[nftAddress][tokenId];
        Offer[] memory noOffers = new Offer[](0);
        address owner_ = IERC721(nftAddress).ownerOf(tokenId);
        if (JamMarketplaceHub(_marketplaceHub).isMarketplace(owner_)) {
            address cancelPermittee = JamMarketplaceHelpers(owner_)
                .auctionCancelPermittee(nftAddress, tokenId);
            if (cancelPermittee == user) return allOffers;
            else return noOffers;
        } else {
            if (owner_ == user) return allOffers;
            else return noOffers;
        }
    }

    function getOffersOf(address offeror)
        external
        view
        returns (Offer[] memory)
    {
        return _offersOf[offeror];
    }

    function getSpecificOffer(
        address offeror,
        address nftAddress,
        uint256 tokenId
    ) public view returns (Offer memory) {
        Offer memory offer;
        for (uint256 i = 0; i < _offersFor[nftAddress][tokenId].length; i++) {
            offer = _offersFor[nftAddress][tokenId][i];
            if (offer.offeror == offeror) return offer;
        }
        return offer;
    }

    function makeOffer(
        address nftAddress,
        uint256 tokenId,
        address currency,
        uint256 amount
    ) external payable {
        // Check offering conditions
        require(
            IERC721(nftAddress).ownerOf(tokenId) != address(0),
            "JamP2PTrading721: NFT not exists"
        );
        require(
            getSpecificOffer(msg.sender, nftAddress, tokenId).offeror ==
                address(0),
            "JamP2PTrading721: already offered before"
        );
        if (currency == address(0))
            require(
                amount == msg.value,
                "JamP2PTrading721: offer amount info mismatch"
            );
        else {
            (bool success, ) = payable(msg.sender).call{value: msg.value}("");
            require(success, "JamP2PTrading721: return money failed");
        }

        // Save the offer's information
        Offer memory offer = Offer(
            msg.sender,
            nftAddress,
            tokenId,
            currency,
            amount
        );
        _offersFor[nftAddress][tokenId].push(offer);
        _offersOf[msg.sender].push(offer);

        // Lock offeror's money
        if (currency != address(0))
            IERC20(currency).transferFrom(msg.sender, address(this), amount);

        emit OfferCreated(msg.sender, nftAddress, tokenId, currency, amount);
    }

    function updateOffer(
        address nftAddress,
        uint256 tokenId,
        address currency,
        uint256 amount
    ) external payable nonReentrant {
        Offer memory offer = getSpecificOffer(msg.sender, nftAddress, tokenId);

        // Check updating conditions
        require(
            offer.offeror == msg.sender,
            "JamP2PTrading721: no offer found"
        );
        if (currency == address(0))
            require(
                amount == msg.value,
                "JamP2PTrading721: offer amount info mismatch"
            );
        else {
            (bool success, ) = payable(msg.sender).call{value: msg.value}("");
            require(success, "JamP2PTrading721: return money failed");
        }

        // Save the new information
        for (uint256 i = 0; i < _offersFor[nftAddress][tokenId].length; i++)
            if (_offersFor[nftAddress][tokenId][i].offeror == msg.sender) {
                _offersFor[nftAddress][tokenId][i].currency = currency;
                _offersFor[nftAddress][tokenId][i].amount = amount;
                break;
            }
        for (uint256 i = 0; i < _offersOf[msg.sender].length; i++)
            if (
                _offersOf[msg.sender][i].nftAddress == nftAddress &&
                _offersOf[msg.sender][i].tokenId == tokenId
            ) {
                _offersOf[msg.sender][i].currency = currency;
                _offersOf[msg.sender][i].amount = amount;
                break;
            }

        // Return old locked offer amount
        if (offer.currency == address(0)) {
            (bool success, ) = payable(msg.sender).call{value: offer.amount}(
                ""
            );
            require(success, "JamP2PTrading721: return money failed");
        } else {
            bool success = IERC20(offer.currency).transfer(
                msg.sender,
                offer.amount
            );
            require(success, "JamP2PTrading721: return money failed");
        }

        // Lock new offer amount
        if (currency != address(0))
            IERC20(currency).transferFrom(msg.sender, address(this), amount);

        emit OfferUpdated(msg.sender, nftAddress, tokenId, currency, amount);
    }

    function cancelOffer(address nftAddress, uint256 tokenId)
        external
        nonReentrant
    {
        Offer memory offer = getSpecificOffer(msg.sender, nftAddress, tokenId);

        // Check cancelling conditions
        require(
            offer.offeror == msg.sender,
            "JamP2PTrading721: sender is not offeror"
        );

        // Delete offer's information
        for (uint256 i = 0; i < _offersFor[nftAddress][tokenId].length; i++)
            if (_offersFor[nftAddress][tokenId][i].offeror == msg.sender) {
                uint256 length = _offersFor[nftAddress][tokenId].length;
                _offersFor[nftAddress][tokenId][i] = _offersFor[nftAddress][
                    tokenId
                ][length - 1];
                _offersFor[nftAddress][tokenId].pop();
                break;
            }
        for (uint256 i = 0; i < _offersOf[msg.sender].length; i++)
            if (
                _offersOf[msg.sender][i].nftAddress == nftAddress &&
                _offersOf[msg.sender][i].tokenId == tokenId
            ) {
                uint256 length = _offersOf[msg.sender].length;
                _offersOf[msg.sender][i] = _offersOf[msg.sender][length - 1];
                _offersOf[msg.sender].pop();
                break;
            }

        // Return locked offer amount
        if (offer.currency == address(0)) {
            (bool success, ) = payable(msg.sender).call{value: offer.amount}(
                ""
            );
            require(success, "JamP2PTrading721: return money failed");
        } else {
            bool success = IERC20(offer.currency).transfer(
                msg.sender,
                offer.amount
            );
            require(success, "JamP2PTrading721: return money failed");
        }

        emit OfferCancelled(msg.sender, nftAddress, tokenId);
    }

    /**
     * @dev The owner of the NFT accepts the offer even when the NFT is on marketplace.
     * @param offeror The person who offered this NFT so far.
     * @param nftAddress - Address of the NFT.
     * @param tokenId - ID of the offered token.
     * @notice If the NFT is currently not on marketplace, the owner must approve it for this contract
     */
    function acceptOffer(
        address offeror,
        address nftAddress,
        uint256 tokenId
    ) external nonReentrant {
        // Check accepting conditions
        Offer memory offer = getSpecificOffer(offeror, nftAddress, tokenId);
        require(offer.offeror == offeror, "JamP2PTrading721: no offer found");

        address owner_ = IERC721(nftAddress).ownerOf(tokenId);

        // Cancel the auction if the NFT is on marketplace
        if (JamMarketplaceHub(_marketplaceHub).isMarketplace(owner_)) {
            JamMarketplaceHelpers marketplace = JamMarketplaceHelpers(owner_);
            address cancelPermittee = marketplace.auctionCancelPermittee(
                nftAddress,
                tokenId
            );
            if (msg.sender == cancelPermittee)
                marketplace.cancelAuction721(nftAddress, tokenId);
        }

        // Sell it to the offeror
        owner_ = IERC721(nftAddress).ownerOf(tokenId);
        IERC721(nftAddress).safeTransferFrom(owner_, offeror, tokenId);

        // Compute auctioneer cut and royalty cut then return proceeds to seller
        if (offer.amount > 0)
            _computeFeesAndPaySeller(
                nftAddress,
                msg.sender,
                offer.currency,
                offer.amount
            );

        emit OfferAccepted(
            offeror,
            msg.sender,
            nftAddress,
            tokenId,
            offer.currency,
            offer.amount
        );
    }
}
