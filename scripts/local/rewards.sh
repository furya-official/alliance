#!/bin/bash

VAL_WALLET_ADDRESS=$(kaijud --home ./data/kaiju keys show val1 --keyring-backend test -a)
DEMO_WALLET_ADDRESS=$(kaijud --home ./data/kaiju keys show demowallet1 --keyring-backend test -a)

VAL_ADDR=$(kaijud query staking validators --output json | jq .validators[0].operator_address --raw-output)
TOKEN_DENOM=ukaijux

printf "\n\n#1) Query wallet balances...\n\n"
kaijud query bank balances $DEMO_WALLET_ADDRESS --home ./data/kaiju

#printf "\n\n#2) Query rewards x/kaiju...\n\n"
#kaijud query kaiju rewards $DEMO_WALLET_ADDRESS $VAL_ADDR $TOKEN_DENOM --home ./data/kaiju

#printf "\n\n#3) Query native staked rewards...\n\n"
#kaijud query distribution rewards $DEMO_WALLET_ADDRESS $VAL_ADDR --home ./data/kaiju

#printf "\n\n#4) Claim rewards from validator...\n\n"
#kaijud tx distribution withdraw-rewards $VAL_ADDR --from=demowallet1 --home ./data/kaiju --keyring-backend=test --broadcast-mode=block --gas 1000000 -y #> /dev/null 2>&1

printf "\n\n#2) Claim rewards from x/kaiju $TOKEN_DENOM...\n\n"
kaijud tx kaiju claim-rewards $VAL_ADDR $TOKEN_DENOM --from=demowallet1 --home ./data/kaiju --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

#printf "\n\n#6) Query rewards x/kaiju...\n\n"
#kaijud query kaiju rewards $DEMO_WALLET_ADDRESS $VAL_ADDR $TOKEN_DENOM --home ./data/kaiju

#printf "\n\n#7) Query native staked rewards...\n\n"
#kaijud query distribution rewards $DEMO_WALLET_ADDRESS $VAL_ADDR --home ./data/kaiju

printf "\n\n#3) Query wallet balances after claim...\n\n"
kaijud query bank balances $DEMO_WALLET_ADDRESS --home ./data/kaiju
