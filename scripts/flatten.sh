#!/bin/sh

set -ex

rm -rf temp
mkdir -p temp

alias flatten="yarn run truffle-flattener"

flatten contracts/SimpleERC20.sol --output temp/SimpleERC20.sol
