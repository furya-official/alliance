#!/bin/bash

DEMO_WALLET_ADDRESS=$(kaijud --home ./data/kaiju keys show demowallet1 --keyring-backend test -a)
VAL_ADDR=$(kaijud query staking validators --output json | jq .validators[0].operator_address --raw-output)
COIN_DENOM=ukaijux
COIN_AMOUNT=$(kaijud query kaiju delegation $DEMO_WALLET_ADDRESS $VAL_ADDR $COIN_DENOM --home ./data/kaiju --output json | jq .delegation.balance.amount --raw-output | sed 's/\.[0-9]*//')
COINS=$COIN_AMOUNT$COIN_DENOM

# FIX: failed to execute message; message index: 0: invalid shares amount: invalid
printf "#1) Undelegate 5000000000$COIN_DENOM from x/kaiju $COIN_DENOM...\n\n"
kaijud tx kaiju undelegate $VAL_ADDR $COINS --from=demowallet1 --home ./data/kaiju --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

printf "\n#2) Query delegations from x/kaiju $COIN_DENOM...\n\n"
kaijud query kaiju kaiju $COIN_DENOM

printf "\n#3) Query delegation on x/kaiju by delegator, validator and $COIN_DENOM...\n\n"
kaijud query kaiju delegation $DEMO_WALLET_ADDRESS $VAL_ADDR $COIN_DENOM --home ./data/kaiju
