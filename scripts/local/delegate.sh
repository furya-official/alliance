#!/bin/bash

CHAIN_DIR=./data
CHAINID=${CHAINID:-kaiju}
COIN_DENOM=ukaijux
VAL_WALLET_ADDRESS=$(kaijud --home $CHAIN_DIR/$CHAINID keys show demowallet1 --keyring-backend test -a)
VAL_ADDR=$(kaijud query staking validators --output json | jq .validators[0].operator_address --raw-output)

printf "#1) Delegate 10000000000$COIN_DENOM thru x/kaiju $COIN_DENOM...\n\n"
kaijud tx kaiju delegate $VAL_ADDR 10000000000$COIN_DENOM --from=demowallet1 --home $CHAIN_DIR/$CHAINID --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

printf "\n#2) Query delegations from x/kaiju $COIN_DENOM...\n\n"
kaijud query kaiju kaiju $COIN_DENOM

printf "\n#3) Query delegation on x/kaiju by delegator, validator and $COIN_DENOM...\n\n"
kaijud query kaiju delegation $VAL_WALLET_ADDRESS $VAL_ADDR $COIN_DENOM
