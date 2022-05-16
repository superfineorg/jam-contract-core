const fs = require("fs");
const rimraf = require("rimraf");

const BUILD_PATH = "./artifacts";
const ABI_PATH = BUILD_PATH + "/abi";
const BIN_PATH = BUILD_PATH + "/bin";
const CONTRACT_PATH = BUILD_PATH + "/contracts";
const RUNTIME_PATH = BUILD_PATH + "/runtime";

let browseBuildFolder = path => {
  if (fs.lstatSync(path).isDirectory()) {
    fs.readdir(path, (err, items) => {
      if (err) {
        console.error(`Error while browsing ${path}`);
        process.exit(1);
      }
      items.forEach((item, _) => browseBuildFolder(path + "/" + item));
    });
  } else if (path.slice(-9) !== ".dbg.json") {
    let splitPath = path.split("/");
    let baseName = splitPath[splitPath.length - 1].split(".")[0];
    let rawData = fs.readFileSync(path);
    let info = JSON.parse(rawData);
    let { abi, bytecode } = info;
    bytecode = bytecode.substring(2);

    if (abi.length !== 0) {
      fs.writeFileSync(ABI_PATH + "/" + baseName + ".abi", JSON.stringify(abi, null, "\t"));
      fs.writeFileSync(BIN_PATH + "/" + baseName + ".bin", bytecode);
    }
  }
};

let main = () => {
  // Remove old build
  rimraf.sync(ABI_PATH);
  rimraf.sync(BIN_PATH);
  rimraf.sync(RUNTIME_PATH);

  // Create empty dirs
  if (!fs.existsSync(BUILD_PATH))
    fs.mkdirSync(BUILD_PATH);
  if (!fs.existsSync(ABI_PATH))
    fs.mkdirSync(ABI_PATH);
  if (!fs.existsSync(BIN_PATH))
    fs.mkdirSync(BIN_PATH);
  if (!fs.existsSync(RUNTIME_PATH))
    fs.mkdirSync(RUNTIME_PATH);

  // Loop through all the files in the temp directory
  browseBuildFolder(CONTRACT_PATH);
};

main();