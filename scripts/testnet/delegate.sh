#!/bin/bash

COIN_DENOM=ibc/4627AD2524E3E0523047E35BB76CC90E37D9D57ACF14F0FCBCEB2480705F3CB8
WALLET_ADDRESS=$(furyad keys show aztestval -a)
VAL_ADDR=$(furyad query staking validators --output json --node=tcp://3.75.187.158:26657 --chain-id=furya-testnet-1 | jq .validators[0].operator_address --raw-output)

printf "#1) Delegating 100000000000$COIN_DENOM thru x/furya...\n\n"
furyad tx furya delegate $VAL_ADDR 100000000000$COIN_DENOM --from=aztestval --node=tcp://3.75.187.158:26657 --chain-id=furya-testnet-1 --broadcast-mode=block -y

printf "\n#2) Delegations groupped by furya $COIN_DENOM...\n\n"
furyad query furya furya $COIN_DENOM --node=tcp://3.75.187.158:26657 --chain-id=furya-testnet-1

printf "\n#3) Delegation by wallet: $WALLET_ADDRESS...\n\n"
furyad query furya delegations-by-delegator $WALLET_ADDRESS --node=tcp://3.75.187.158:26657 --chain-id=furya-testnet-1

printf "\n#4) Delegations by wallet: $WALLET_ADDRESS and validator: $VAL_ADDR...\n\n"
furyad query furya delegations-by-delegator-and-validator $WALLET_ADDRESS $VAL_ADDR --node=tcp://3.75.187.158:26657 --chain-id=furya-testnet-1

printf "\n#5) Delegation by wallet: $WALLET_ADDRESS, validator: $VAL_ADDR and denom: $COIN_DENOM...\n\n"
furyad query furya delegation $WALLET_ADDRESS $VAL_ADDR $COIN_DENOM --node=tcp://3.75.187.158:26657 --chain-id=furya-testnet-1
