#!/bin/bash

COIN_DENOM=ibc/4627AD2524E3E0523047E35BB76CC90E37D9D57ACF14F0FCBCEB2480705F3CB8
WALLET_ADDRESS=$(kaijud keys show aztestval -a)
VAL_ADDR=$(kaijud query staking validators --output json --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1 | jq .validators[0].operator_address --raw-output)

printf "\n\n#1) Query $COIN_DENOM kaiju rewards...\n\n"
kaijud query kaiju rewards $WALLET_ADDRESS $VAL_ADDR $COIN_DENOM --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1

printf "\n\n#2) Claim rewards from x/kaiju $COIN_DENOM...\n\n"
kaijud tx kaiju claim-rewards $VAL_ADDR $COIN_DENOM --from=aztestval --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1 --gas=auto --broadcast-mode=block -y

printf "\n\n#3) Query $COIN_DENOM kaiju rewards after claim...\n\n"
kaijud query kaiju rewards $WALLET_ADDRESS $VAL_ADDR $COIN_DENOM --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1
