#!/bin/bash

CHAIN_DIR=./data
CHAINID=${CHAINID:-furya}
COIN_DENOM=ufuryx
VAL_WALLET_ADDRESS=$(furyad --home $CHAIN_DIR/$CHAINID keys show demowallet1 --keyring-backend test -a)
VAL_ADDR=$(furyad query staking validators --output json | jq .validators[0].operator_address --raw-output)

printf "#1) Delegate 10000000000$COIN_DENOM thru x/furya $COIN_DENOM...\n\n"
furyad tx furya delegate $VAL_ADDR 10000000000$COIN_DENOM --from=demowallet1 --home $CHAIN_DIR/$CHAINID --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

printf "\n#2) Query delegations from x/furya $COIN_DENOM...\n\n"
furyad query furya furya $COIN_DENOM

printf "\n#3) Query delegation on x/furya by delegator, validator and $COIN_DENOM...\n\n"
furyad query furya delegation $VAL_WALLET_ADDRESS $VAL_ADDR $COIN_DENOM
