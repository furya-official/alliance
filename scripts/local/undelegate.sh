#!/bin/bash

DEMO_WALLET_ADDRESS=$(furyad --home ./data/furya keys show demowallet1 --keyring-backend test -a)
VAL_ADDR=$(furyad query staking validators --output json | jq .validators[0].operator_address --raw-output)
COIN_DENOM=ufuryx
COIN_AMOUNT=$(furyad query furya delegation $DEMO_WALLET_ADDRESS $VAL_ADDR $COIN_DENOM --home ./data/furya --output json | jq .delegation.balance.amount --raw-output | sed 's/\.[0-9]*//')
COINS=$COIN_AMOUNT$COIN_DENOM

# FIX: failed to execute message; message index: 0: invalid shares amount: invalid
printf "#1) Undelegate 5000000000$COIN_DENOM from x/furya $COIN_DENOM...\n\n"
furyad tx furya undelegate $VAL_ADDR $COINS --from=demowallet1 --home ./data/furya --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

printf "\n#2) Query delegations from x/furya $COIN_DENOM...\n\n"
furyad query furya furya $COIN_DENOM

printf "\n#3) Query delegation on x/furya by delegator, validator and $COIN_DENOM...\n\n"
furyad query furya delegation $DEMO_WALLET_ADDRESS $VAL_ADDR $COIN_DENOM --home ./data/furya
