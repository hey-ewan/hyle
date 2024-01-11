#!/bin/bash

# rm -r ~/.minid || true
#MINID_BIN=$(which minid)
MINID_BIN=./minid
# configure minid
$MINID_BIN config set client chain-id demo
$MINID_BIN config set client keyring-backend test
$MINID_BIN config set app minimum-gas-prices 0.1mini
$MINID_BIN keys add alice
$MINID_BIN keys add bob
$MINID_BIN init test --chain-id demo --default-denom mini
# update genesis
$MINID_BIN genesis add-genesis-account alice 10000000mini --keyring-backend test
$MINID_BIN genesis add-genesis-account bob 1000mini --keyring-backend test
# create default validator
$MINID_BIN genesis gentx alice 1000000mini --chain-id demo
$MINID_BIN genesis collect-gentxs