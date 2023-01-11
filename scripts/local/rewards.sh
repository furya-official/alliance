#!/bin/bash

VAL_WALLET_ADDRESS=$(furyad --home ./data/furya keys show val1 --keyring-backend test -a)
DEMO_WALLET_ADDRESS=$(furyad --home ./data/furya keys show demowallet1 --keyring-backend test -a)

VAL_ADDR=$(furyad query staking validators --output json | jq .validators[0].operator_address --raw-output)
TOKEN_DENOM=ufuryx

printf "\n\n#1) Query wallet balances...\n\n"
furyad query bank balances $DEMO_WALLET_ADDRESS --home ./data/furya

#printf "\n\n#2) Query rewards x/furya...\n\n"
#furyad query furya rewards $DEMO_WALLET_ADDRESS $VAL_ADDR $TOKEN_DENOM --home ./data/furya

#printf "\n\n#3) Query native staked rewards...\n\n"
#furyad query distribution rewards $DEMO_WALLET_ADDRESS $VAL_ADDR --home ./data/furya

#printf "\n\n#4) Claim rewards from validator...\n\n"
#furyad tx distribution withdraw-rewards $VAL_ADDR --from=demowallet1 --home ./data/furya --keyring-backend=test --broadcast-mode=block --gas 1000000 -y #> /dev/null 2>&1

printf "\n\n#2) Claim rewards from x/furya $TOKEN_DENOM...\n\n"
furyad tx furya claim-rewards $VAL_ADDR $TOKEN_DENOM --from=demowallet1 --home ./data/furya --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

#printf "\n\n#6) Query rewards x/furya...\n\n"
#furyad query furya rewards $DEMO_WALLET_ADDRESS $VAL_ADDR $TOKEN_DENOM --home ./data/furya

#printf "\n\n#7) Query native staked rewards...\n\n"
#furyad query distribution rewards $DEMO_WALLET_ADDRESS $VAL_ADDR --home ./data/furya

printf "\n\n#3) Query wallet balances after claim...\n\n"
furyad query bank balances $DEMO_WALLET_ADDRESS --home ./data/furya
