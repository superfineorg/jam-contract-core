/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.6;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/interfaces/IERC2981.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165Checker.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "./JamMarketplaceHub.sol";

abstract contract JamMarketplaceHelpers is Ownable, Pausable, ReentrancyGuard {
    // The ID of this marketplace in Gamejam's marketplace system
    bytes32 public marketplaceId;

    // Cut owner takes on each auction. Values 0 - 10,000 map to 0% - 100%
    uint256 public ownerCut;

    // The minimum duration between 2 continuous royalty withdrawals
    uint256 private _withdrawDuration = 14 days;

    // The address of marketplace hub
    address internal _marketplaceHub;

    // Mapping from (royalty recipient, erc20 currency) to the amount of royalty cut he receives
    mapping(address => mapping(address => uint256)) internal _royaltyCuts;

    // The total amount of royalty cut which cannot be reclaimed by the owner of the contract
    mapping(address => uint256) internal _totalRoyaltyCut;

    // Mapping from (royalty recipient, erc20 currency) to the last moment he withdraws the royalty cut
    mapping(address => mapping(address => uint256)) private _lastWithdraws;

    /**
     * @dev Constructor that rejects incoming Ether
     * @dev The `payable` flag is added so we can access `msg.value` without compiler warning. If we
     * leave out payable, then Solidity will allow inheriting contracts to implement a payable
     * constructor. By doing it this way we prevent a payable constructor from working. Alternatively
     * we could use assembly to access msg.value.
     */
    constructor(address hubAddress, uint256 ownerCut_) payable {
        require(
            msg.value == 0,
            "JamMarketplaceHelpers: cannot send native token when deploying"
        );
        require(
            hubAddress != address(0),
            "JamMarketplaceHelpers: invalid hub address"
        );
        require(
            ownerCut_ <= 10000,
            "JamMarketplaceHelpers: owner cut cannot exceed 100%"
        );
        _marketplaceHub = hubAddress;
        ownerCut = ownerCut_;
    }

    /**
     * @dev Check if this auction is currently cancelable
     * @param nftAddress - address of a deployed contract implementing the non-fungible interface.
     * @param tokenId - ID of token to auction
     */
    function isAuctionCancelable(address nftAddress, uint256 tokenId)
        external
        virtual
        returns (bool)
    {
        return true;
    }

    function getReceivedRoyalty(address user, address[] memory currencies)
        external
        view
        returns (
            uint256[] memory receivedAmounts,
            uint256[] memory lastWithdraws
        )
    {
        receivedAmounts = new uint256[](currencies.length);
        lastWithdraws = new uint256[](currencies.length);
        for (uint256 i = 0; i < currencies.length; i++) {
            uint256 balance = _royaltyCuts[user][currencies[i]];
            receivedAmounts[i] = balance;
            uint256 lastWithdraw = _lastWithdraws[user][currencies[i]];
            lastWithdraws[i] = lastWithdraw;
        }
    }

    function registerWithHub() external onlyOwner {
        JamMarketplaceHub(_marketplaceHub).registerMarketplace(marketplaceId);
    }

    function cancelAuction(address nftAddress, uint256 tokenId)
        external
        virtual
    {}

    function withdrawRoyalty(address currency) public {
        uint256 lastWithdraw = _lastWithdraws[msg.sender][currency];
        require(
            lastWithdraw + _withdrawDuration <= block.timestamp,
            "JamMarketplaceHelpers: only withdraw after 14 days after previous withdraw"
        );
        uint256 royaltyCut = _royaltyCuts[msg.sender][currency];
        require(
            royaltyCut > 0,
            "JamMarketplaceHelpers: no royalty cut to withdraw"
        );
        _royaltyCuts[msg.sender][currency] = 0;
        _totalRoyaltyCut[currency] -= royaltyCut;
        _lastWithdraws[msg.sender][currency] = block.timestamp;
        if (currency == address(0)) {
            (bool success, ) = payable(msg.sender).call{value: royaltyCut}("");
            require(success, "JamMarketplaceHelpers: withdraw failed");
        } else {
            IERC20 erc20Contract = IERC20(currency);
            bool success = erc20Contract.transfer(msg.sender, royaltyCut);
            require(success, "JamMarketplaceHelpers: withdraw failed");
        }
    }

    function reclaim(address currency) external onlyOwner {
        if (currency == address(0)) {
            (bool success, ) = payable(owner()).call{
                value: address(this).balance - _totalRoyaltyCut[address(0)]
            }("");
            require(success, "JamMarketplaceHelpers: reclaim failed");
        } else {
            IERC20 currencyContract = IERC20(currency);
            bool success = IERC20(currency).transfer(
                owner(),
                currencyContract.balanceOf(address(this)) -
                    _totalRoyaltyCut[currency]
            );
            require(success, "JamMarketplaceHelpers: reclaim failed");
        }
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function _supportIERC2981(address nftAddress) internal view returns (bool) {
        bool success;
        success = ERC165Checker.supportsERC165(nftAddress);
        if (success) {
            success = IERC165(nftAddress).supportsInterface(
                type(IERC2981).interfaceId
            );
        }
        return success;
    }
}
