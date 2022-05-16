const FileSystem = require("fs");
const Straightener = require("sol-straightener");
const schema = require("./token-schema.json");

const RESULT = "./scripts/token-generator/result.sol";
const OPENZEPPELIN = {
  Ownable: "@openzeppelin/contracts/access/Ownable.sol",
  AccessControl: "@openzeppelin/contracts/access/AccessControl.sol",
  ERC20: "@openzeppelin/contracts/token/ERC20/ERC20.sol",
  ERC20Capped: "@openzeppelin/contracts/token/ERC20/extensions/ERC20Capped.sol",
  ERC20Burnable: "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol"
};
const VALUES = {
  FIXED: "fixed",
  UNLIMITED: "unlimited",
  CAPPED: "capped",
  OWNABLE: "ownable",
  ROLE_BASED: "roleBased",
  ERC20: "erc20",
  ERC721: "erc721",
  ERC1155: "erc1155"
};

module.exports = async function generateErc20() {
  // Prepare the parameters
  let tokenName = schema.tokenName.replaceAll(" ", "");
  let tokenSymbol = schema.tokenSymbol.replaceAll(" ", "");
  let capacity = schema.capacity?.toString();

  // List necessary Openzeppelin contracts
  let inheritedContracts = [OPENZEPPELIN.ERC20];
  if (schema.accessType === VALUES.OWNABLE)
    inheritedContracts.push(OPENZEPPELIN.Ownable);
  else
    inheritedContracts.push(OPENZEPPELIN.AccessControl);
  if (schema.supplyType === VALUES.CAPPED) {
    inheritedContracts = inheritedContracts.filter(contract => contract != OPENZEPPELIN.ERC20);
    inheritedContracts.push(OPENZEPPELIN.ERC20Capped);
  }
  if (schema.burnable) {
    inheritedContracts = inheritedContracts.filter(contract => contract != OPENZEPPELIN.ERC20);
    inheritedContracts.push(OPENZEPPELIN.ERC20Burnable);
  }
  if (schema.flatten) {
    inheritedContracts = inheritedContracts.map(contract => `./${contract.substring(1)}`);
  }

  // Generate contract header
  let header = `
    /* SPDX-License-Identifier: MIT */

    pragma solidity ^0.8.0;

    ${inheritedContracts.map(contract => `import "${contract}";`).join("\n")}

    contract ${tokenName} is ${inheritedContracts.map(contract => getContractName(contract))} {
  `;
  write(header);

  // Generate variables
  if (schema.accessType === VALUES.ROLE_BASED) {
    let vars = `bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");\n`;
    append(vars);
  }

  // Generate constructor
  let constructor = `
    constructor()
    ERC20("${tokenName}", "${tokenSymbol}")
    ${schema.supplyType === VALUES.CAPPED ? `ERC20Capped(${capacity}*10**18)` : ""}
    ${schema.accessType === VALUES.OWNABLE ? `Ownable()` : ""} {
      ${schema.accessType === VALUES.ROLE_BASED ? "_setupRole(DEFAULT_ADMIN_ROLE, msg.sender);" : ""}
      ${schema.initialSupply ? schema.supplyType === VALUES.CAPPED ?
      `require(${schema.initialSupply} <= ${schema.capacity}, "${tokenName}: initial supply exceeds capacity");\nERC20._mint(msg.sender, ${schema.initialSupply}*10**18);` :
      `_mint(msg.sender, ${schema.initialSupply}*10**18);` : ""}
    }
  `;
  append(constructor);

  // Override token decimals if necessary
  if (schema.tokenDecimals !== 18) {
    let decimals = `
      function decimals() public pure override returns (uint8) {
        return ${schema.tokenDecimals};
      }
    `;
    append(decimals);
  }

  // Generate mint function
  if (schema.mintable) {
    let mintFunction = `
      function mint(address recipient, uint256 amount) external ${schema.accessType === VALUES.OWNABLE ? "onlyOwner" : ""} {
        ${schema.accessType === VALUES.ROLE_BASED ? `require(hasRole(MINTER_ROLE, msg.sender), "${tokenName}: caller is not minter");` : ""}
        _mint(recipient, amount);
      }
    `;
    append(mintFunction);

    if (schema.supplyType === VALUES.CAPPED) {
      let internalMintFunction = `
        function _mint(address recipient, uint256 amount) internal override${schema.burnable ? "(ERC20, ERC20Capped)" : ""} {
            ERC20Capped._mint(recipient, amount);
        }
      `;
      append(internalMintFunction);
    }
  }

  // Finish the contract
  append("}");

  // Flatten the contract
  if (schema.flatten) {
    let flattenedCode = await Straightener.straighten(RESULT);
    let removeSPDX = flattenedCode.split("\n").filter(line => !line.includes("SPDX")).join("\n");
    let finalCode = "/* SPDX-License-Identifier: MIT */\n\n" + removeSPDX;
    write(finalCode);
  }
};

let getContractName = path => {
  let steps = path.split("/");
  let contract = steps[steps.length - 1];
  let parts = contract.split(".");
  return parts[0];
};

let write = content => {
  FileSystem.writeFileSync(RESULT, content, err => {
    if (err) {
      console.log("Error when generating!");
      return;
    } else
      console.log("Generate successfully!");
  });
};

let append = content => {
  FileSystem.appendFileSync(RESULT, content, err => {
    if (err) {
      console.log("Error when generating!");
      return;
    } else
      console.log("Generate successfully!");
  });
};