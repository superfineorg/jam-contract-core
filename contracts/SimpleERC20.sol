pragma solidity ^0.8.0;


import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "./MinterRole.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract CoinToken is ERC20, ERC20Burnable,  Ownable {
    using SafeMath for uint256;

    uint256 private _totalMinted;
    uint256 private _cap;
    uint8 private _decimals;
    address public coinTokenAddress;
    
    constructor(address _owner, uint256 capacity, uint256 initialSupply, uint8 d, string memory _name, string memory _symbol) ERC20(_name, _symbol) {
        _cap = capacity;
        owner = _owner;
        _decimals = d;
        _mint(_owner, initialSupply);
        coinTokenAddress = address(this);
    }
    
    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }
    
    function _mint(address account, uint256 amount) internal override {
        super._mint(account, amount);
        _totalMinted = _totalMinted.add(amount);
    }

    function cap() public view returns (uint256) {
        return _cap;
    }

    function totalMinted() public view returns (uint256) {
        return _totalMinted;
    }

    function mintBulk(address[] memory accounts, uint256[] memory amounts)
        public
        onlyOwner
        returns (bool)
    {
        require(accounts.length == amounts.length, "arrays must have same length");
        for (uint256 i = 0; i < accounts.length; i++) {
            require(amounts[i] > 0, "amount must be greater than 0");
            _mint(accounts[i], amounts[i]);
        }
        return true;
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override {
        super._beforeTokenTransfer(from, to, amount);

        if (from == address(0)) {
            // When minting tokens
            require(totalMinted().add(amount) <= cap(), "ERC20Capped: cap exceeded");
        }
    }

    function addMinter(address account) public override onlyOwner {
        _addMinter(account);
    }

    function removeMinter(address account) public onlyOwner {
        _removeMinter(account);
    }
}