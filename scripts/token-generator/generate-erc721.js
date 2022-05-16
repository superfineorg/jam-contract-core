const FileSystem = require("fs");
const Straightener = require("sol-straightener");
const schema = require("./token-schema.json");

const RESULT = "./scripts/token-generator/result.sol";
const OPENZEPPELIN = {
  Ownable: "@openzeppelin/contracts/access/Ownable.sol",
  AccessControl: "@openzeppelin/contracts/access/AccessControl.sol",
  Counters: "@openzeppelin/contracts/utils/Counters.sol",
  ERC721: "@openzeppelin/contracts/token/ERC721/ERC721.sol",
  ERC721Enumerable: "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol",
  ERC721Burnable: "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol"
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

module.exports = async function generateErc721() {
  // Prepare the parameters
  let tokenName = schema.tokenName.replaceAll(" ", "");
  let tokenSymbol = schema.tokenSymbol.replaceAll(" ", "");
  let baseUri = schema.baseUri.replaceAll(" ", "");

  // List necessary Openzeppelin contracts
  let inheritedContracts = [OPENZEPPELIN.Counters, OPENZEPPELIN.ERC721];
  if (schema.accessType === VALUES.OWNABLE)
    inheritedContracts.push(OPENZEPPELIN.Ownable);
  else
    inheritedContracts.push(OPENZEPPELIN.AccessControl);
  if (schema.enumerable) {
    inheritedContracts = inheritedContracts.filter(contract => contract != OPENZEPPELIN.ERC721);
    inheritedContracts.push(OPENZEPPELIN.ERC721Enumerable);
  }
  if (schema.burnable) {
    inheritedContracts = inheritedContracts.filter(contract => contract != OPENZEPPELIN.ERC721);
    inheritedContracts.push(OPENZEPPELIN.ERC721Burnable);
  }
  if (schema.flatten) {
    inheritedContracts = inheritedContracts.map(contract => `./${contract.substring(1)}`);
  }

  // Generate contract header
  let header = `
    /* SPDX-License-Identifier: MIT */

    pragma solidity ^0.8.0;

    ${inheritedContracts.map(contract => `import "${contract}";`).join("\n")}

    contract ${tokenName} is ${inheritedContracts.filter(contract => !contract.includes("Counters")).map(contract => getContractName(contract))} {
  `;
  write(header);

  // Generate variables
  let vars = `
    using Counters for Counters.Counter;

    string public baseTokenURI;
    Counters.Counter private _currentId;
  `;
  append(vars);
  if (schema.accessType === VALUES.ROLE_BASED) {
    let roleVars = `bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");`;
    append(roleVars);
  }

  // Generate constructor
  let constructor = `
    constructor()
    ERC721("${tokenName}", "${tokenSymbol}")
    ${schema.accessType === VALUES.OWNABLE ? `Ownable()` : ""} {
      ${schema.accessType === VALUES.ROLE_BASED ? "_setupRole(DEFAULT_ADMIN_ROLE, msg.sender);" : ""}
      baseTokenURI = "${baseUri}";
    }
  `;
  append(constructor);

  // Generate supportsInterface() function if necessary
  let interfaces = [];
  if (schema.enumerable)
    interfaces.push("ERC721Enumerable");
  if (schema.burnable)
    interfaces.push("ERC721");
  if (schema.accessType === VALUES.ROLE_BASED)
    interfaces.push("AccessControl");
  if (interfaces.length >= 2) {
    let supportsInterface = `
      function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(${interfaces})
        returns (bool)
      {
        return ${interfaces.map(interface => `${interface}.supportsInterface(interfaceId)`).join("||")};
      }
    `;
    append(supportsInterface);
  }

  // Generate tokenURI() function
  let tokenURI = `
    function tokenURI(uint256 tokenId)
      public
      view
      override
      returns (string memory)
    {
      require(_exists(tokenId), "${tokenName}: token does not exist");
      return
        string(
          abi.encodePacked(
            baseTokenURI,
            Strings.toString(tokenId),
            ".json"
          )
        );
    }
  `;
  append(tokenURI);

  // Generate a function to set base token URI
  let setBaseURI = `
    function setBaseTokenURI(string memory baseTokenURI_) external ${schema.accessType === VALUES.OWNABLE ? "onlyOwner" : ""} {
      ${schema.accessType === VALUES.ROLE_BASED ? `require(hasRole(DEFAULT_ADMIN_ROLE, msg.sender), "${tokenName}: must have admin role to set");` : ""}
      baseTokenURI = baseTokenURI_;
    }
  `;
  append(setBaseURI);

  // Generate mint function
  let mintFunction = `
    function mint(address to) external ${schema.accessType === VALUES.OWNABLE ? "onlyOwner" : ""} {
      ${schema.accessType === VALUES.ROLE_BASED ? `require(hasRole(MINTER_ROLE, msg.sender), "${tokenName}: must have minter role to mint");` : ""}
      _safeMint(to, _currentId.current());
      _currentId.increment();
    }
  `;
  append(mintFunction);

  // Generate _beforeTokenTransfer() function if necessary
  if (schema.burnable && schema.enumerable) {
    let beforeTokenTransfer = `
      function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId
      ) internal virtual override(ERC721, ERC721Enumerable) {
        ERC721Enumerable._beforeTokenTransfer(from, to, tokenId);
      }
    `;
    append(beforeTokenTransfer);
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