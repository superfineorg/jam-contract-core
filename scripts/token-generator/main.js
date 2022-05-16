const schema = require("./token-schema.json");
const generateErc20 = require("./generate-erc20");
const generateErc721 = require("./generate-erc721");
const generateErc1155 = require("./generate-erc1155");

let generateToken = async () => {
  switch (schema.tokenType) {
    case "erc20":
      await generateErc20();
      break;
    case "erc721":
      await generateErc721();
      break;
    case "erc1155":
      await generateErc1155();
      break;
    default:
      console.log("Invalid token type, process exits with error!");
      return;
  }
};

let main = async () => {
  await generateToken();
};

main();