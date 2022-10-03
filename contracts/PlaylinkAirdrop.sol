/* SPDX-License-Identifier: MIT */

pragma solidity ^0.8.15;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC1155/IERC1155.sol";

contract PlaylinkAirdrop is Ownable, ReentrancyGuard {
    enum AssetType {
        ERC20,
        ERC721,
        ERC1155
    }

    struct Asset {
        AssetType assetType;
        address assetAddress;
        uint256 assetId; // 0 for ERC20
        uint256 availableAmount; // 1 for ERC721
    }

    struct AirdropCampaign {
        string campaignId;
        address creator;
        Asset[] assets;
        uint256 maxBatchSize;
        uint256 startingTime;
        uint256 totalAvailableAssets;
        uint256 airdropFee;
    }

    uint256 private _maxBatchSize;
    uint256 private _feePerBatch;
    mapping(string => AirdropCampaign) private _campaignById;
    mapping(address => bool) private _operators;

    event AirdropCampaignCreated(
        string campaignId,
        address creator,
        Asset[] assets,
        uint256 maxBatchSize,
        uint256 startingTime
    );

    event AirdropCampaignUpdated(
        string campaignId,
        address creator,
        Asset[] assets,
        uint256 maxBatchSize,
        uint256 startingTime
    );

    event AssetsAirdropped(
        string campaignId,
        address creator,
        Asset[] assets,
        address[] recipients
    );

    constructor(uint256 maxBatchSize, uint256 feePerBatch) Ownable() {
        require(
            maxBatchSize > 0,
            "PlaylinkAirdrop: batch size must be greater than zero"
        );
        _maxBatchSize = maxBatchSize;
        _feePerBatch = feePerBatch;
        _operators[msg.sender] = true;
    }

    modifier onlyOperators() {
        require(
            _operators[msg.sender],
            "PlaylinkAirdrop: caller is not operator"
        );
        _;
    }

    function getCampaignById(string memory campaignId)
        external
        view
        returns (AirdropCampaign memory)
    {
        return _campaignById[campaignId];
    }

    function estimateAirdropFee(uint256 numAssets)
        public
        view
        returns (uint256)
    {
        uint256 numRequiredBatches = (numAssets + _maxBatchSize - 1) /
            _maxBatchSize; // ceil(numAssets / _maxBatchSize)
        return numRequiredBatches * _feePerBatch;
    }

    function setOperators(address[] memory operators, bool[] memory isOperators)
        external
        onlyOwner
    {
        require(
            operators.length == isOperators.length,
            "PlaylinkAirdrop: lengths mismatch"
        );
        for (uint256 i = 0; i < operators.length; i++)
            _operators[operators[i]] = isOperators[i];
    }

    function setMaxBatchSize(uint256 newSize) external onlyOperators {
        require(
            newSize > 0,
            "PlaylinkAirdrop: batch size must be greater than zero"
        );
        _maxBatchSize = newSize;
    }

    function setFeePerBatch(uint256 newFee) external onlyOperators {
        _feePerBatch = newFee;
    }

    function createAirdropCampaign(
        string memory campaignId,
        Asset[] memory assets,
        uint256 startingTime
    ) external payable nonReentrant {
        AirdropCampaign storage campaign = _campaignById[campaignId];

        // Check if campaign exists
        require(
            campaign.creator == address(0),
            "PlaylinkAirdrop: campaign already created"
        );

        // Check payment
        uint256 airdropFee = estimateAirdropFee(assets.length);
        require(
            msg.value >= airdropFee,
            "PlaylinkAirdrop: insufficient airdrop fee"
        );
        if (msg.value > airdropFee) {
            (bool success, ) = payable(msg.sender).call{
                value: msg.value - airdropFee
            }("");
            require(success, "PlaylinkAirdrop: failed to return excess");
        }

        // Validate data
        require(
            block.timestamp < startingTime,
            "PlaylinkAirdrop: starting time too low"
        );
        for (uint256 i = 0; i < assets.length; i++) {
            Asset memory asset = assets[i];
            require(
                uint256(asset.assetType) <= 3,
                "PlaylinkAirdrop: invalid asset type"
            );
            if (asset.assetType == AssetType.ERC20)
                require(
                    asset.assetId == 0,
                    "PlaylinkAirdrop: invalid ERC20 asset ID"
                );
            else if (asset.assetType == AssetType.ERC721)
                require(
                    asset.availableAmount == 1,
                    "PlaylinkAirdrop: invalid ERC721 amount"
                );
        }

        // Create new airdrop campaign
        uint256 totalAvailableAssets = 0;
        for (uint256 j = 0; j < assets.length; j++)
            totalAvailableAssets += assets[j].availableAmount;
        campaign.campaignId = campaignId;
        campaign.creator = msg.sender;
        for (uint256 k = 0; k < assets.length; k++)
            campaign.assets.push(assets[k]);
        campaign.maxBatchSize = _maxBatchSize;
        campaign.startingTime = startingTime;
        campaign.totalAvailableAssets = totalAvailableAssets;
        campaign.airdropFee = airdropFee;

        emit AirdropCampaignCreated(
            campaignId,
            msg.sender,
            assets,
            _maxBatchSize,
            startingTime
        );
    }

    function updateCampaign(
        string memory campaignId,
        Asset[] memory assets,
        uint256 startingTime
    ) external payable nonReentrant {
        AirdropCampaign storage campaign = _campaignById[campaignId];

        // Check campaign ownership
        require(
            campaign.creator == msg.sender,
            "PlaylinkAirdrop: caller is not campaign owner"
        );

        // Make sure that this campaign has not started yet
        require(
            block.timestamp < campaign.startingTime,
            "PlaylinkAirdrop: campaign started, cannot update assets"
        );

        // Check payment
        uint256 newAirdropFee = estimateAirdropFee(assets.length);
        if (newAirdropFee > campaign.airdropFee) {
            require(
                msg.value >= newAirdropFee - campaign.airdropFee,
                "PlaylinkAirdrop:insufficient airdrop fee"
            );
            if (msg.value > newAirdropFee - campaign.airdropFee) {
                (bool success, ) = payable(msg.sender).call{
                    value: msg.value + campaign.airdropFee - newAirdropFee
                }("");
                require(success, "PlaylinkAirdrop: failed to return excess");
            }
        }

        // Validate data
        require(
            block.timestamp < startingTime,
            "PlaylinkAirdrop: starting time too low"
        );
        for (uint256 i = 0; i < assets.length; i++) {
            Asset memory asset = assets[i];
            require(
                uint256(asset.assetType) <= 3,
                "PlaylinkAirdrop: invalid asset type"
            );
            if (asset.assetType == AssetType.ERC20)
                require(
                    asset.assetId == 0,
                    "PlaylinkAirdrop: invalid ERC20 asset ID"
                );
            else if (asset.assetType == AssetType.ERC721)
                require(
                    asset.availableAmount == 1,
                    "PlaylinkAirdrop: invalid ERC721 amount"
                );
        }

        // Update campaign assets and airdrop fee
        uint256 totalAvailableAssets = 0;
        for (uint256 j = 0; j < assets.length; j++)
            totalAvailableAssets += assets[j].availableAmount;
        delete campaign.assets;
        for (uint256 k = 0; k < assets.length; k++)
            campaign.assets.push(assets[k]);
        campaign.maxBatchSize = _maxBatchSize;
        campaign.startingTime = startingTime;
        campaign.totalAvailableAssets = totalAvailableAssets;
        campaign.airdropFee = newAirdropFee;

        emit AirdropCampaignUpdated(
            campaignId,
            msg.sender,
            assets,
            campaign.maxBatchSize,
            startingTime
        );
    }

    function airdrop(
        string memory campaignId,
        uint256[] memory assetIndexes,
        address[] memory recipients
    ) external onlyOperators nonReentrant {
        require(
            _campaignById[campaignId].creator != address(0),
            "PlaylinkAirdrop: campaign does not exist"
        );
        AirdropCampaign storage campaign = _campaignById[campaignId];
        require(
            block.timestamp > campaign.startingTime,
            "PlaylinkAirdrop: campaign not start yet"
        );
        require(
            assetIndexes.length == recipients.length,
            "PlaylinkAirdrop: lengths mismatch"
        );
        require(
            assetIndexes.length <= campaign.maxBatchSize,
            "PlaylinkAirdrop: too many assets airdropped"
        );
        Asset[] memory airdroppedAssets = new Asset[](assetIndexes.length);
        for (uint256 i = 0; i < assetIndexes.length; i++) {
            airdroppedAssets[i] = campaign.assets[assetIndexes[i]];
            Asset storage asset = campaign.assets[assetIndexes[i]];
            if (asset.assetType == AssetType.ERC20) {
                bool success = IERC20(asset.assetAddress).transferFrom(
                    campaign.creator,
                    recipients[i],
                    asset.availableAmount
                );
                require(
                    success,
                    "PlaylinkAirdrop: failed to send ERC20 assets"
                );
                campaign.totalAvailableAssets -= asset.availableAmount;
                asset.availableAmount = 0;
            } else if (asset.assetType == AssetType.ERC721) {
                IERC721(asset.assetAddress).transferFrom(
                    campaign.creator,
                    recipients[i],
                    asset.assetId
                );
                campaign.totalAvailableAssets--;
                asset.availableAmount = 0;
            } else if (asset.assetType == AssetType.ERC1155) {
                IERC1155(asset.assetAddress).safeTransferFrom(
                    campaign.creator,
                    recipients[i],
                    asset.assetId,
                    asset.availableAmount,
                    abi.encodePacked("Airdrop ERC1155 assets")
                );
                campaign.totalAvailableAssets -= asset.availableAmount;
                asset.availableAmount = 0;
            }
        }
        emit AssetsAirdropped(
            campaignId,
            campaign.creator,
            airdroppedAssets,
            recipients
        );
    }

    function withdrawAirdropFee(address recipient) external onlyOwner {
        (bool success, ) = payable(recipient).call{
            value: address(this).balance
        }("");
        require(success, "PlaylinkAirdrop: failed to withdraw airdrop fee");
    }
}
