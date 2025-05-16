#!/usr/bin/env bash

node_path="--home ./nodes/node0"

#rm -r ~/.exampled || true
rm -rf ./nodes
mkdir -p ./nodes/node0
exampled_BIN=$(which exampled)
# configure exampled
$exampled_BIN config set client chain-id demo $node_path
$exampled_BIN config set client keyring-backend test $node_path
$exampled_BIN keys add alice $node_path --keyring-backend=test
$exampled_BIN keys add bob $node_path --keyring-backend=test
$exampled_BIN init test --chain-id demo $node_path
# update genesis
$exampled_BIN genesis add-genesis-account alice 10000000000000uatom --keyring-backend test $node_path
$exampled_BIN genesis add-genesis-account bob 1000000000uatom --keyring-backend test $node_path
# create default validator
$exampled_BIN genesis gentx alice 1000000000000uatom --chain-id demo $node_path --keyring-backend test
$exampled_BIN genesis collect-gentxs $node_path

genesis_path="./nodes/node0/config/genesis.json"
jq '.app_state.staking.params.bond_denom = "uatom"
    | (.app_state.gov.params.min_deposit[] | select(.denom == "stake").denom) = "uatom"
    | (.app_state.gov.params.expedited_min_deposit[] | select(.denom == "stake").denom) = "uatom"' $genesis_path > temp.json && mv temp.json $genesis_path

nohup exampled start --home ./nodes/node0 > ./nodes/node0/node0.log 2>&1 &


