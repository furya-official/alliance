#!/bin/bash

COIN_DENOM=ibc/4627AD2524E3E0523047E35BB76CC90E37D9D57ACF14F0FCBCEB2480705F3CB8

printf "#1) Submit proposal to create $COIN_DENOM Kaiju...\n\n"
kaijud tx gov submit-legacy-proposal create-kaiju $COIN_DENOM 0.5 0.5 --from=aztestval --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1 --broadcast-mode=block -y
PROPOSAL_ID=$(kaijud query gov proposals --count-total --output json --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1 | jq .proposals[0].id -r)

printf "\n#2) Deposit funds to proposal $PROPOSAL_ID...\n\n"
kaijud tx gov deposit $PROPOSAL_ID 10000000stake --from=aztestval --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1 --broadcast-mode=block -y

printf "\n#3) Vote to pass the proposal...\n\n"
kaijud tx gov vote $PROPOSAL_ID yes --from=aztestval --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1 --broadcast-mode=block -y
kaijud tx gov vote $PROPOSAL_ID yes --from=aztestval2 --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1 --broadcast-mode=block -y
kaijud tx gov vote $PROPOSAL_ID yes --from=aztestval3 --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1 --broadcast-mode=block -y

printf "\n#4) Query proposals...\n\n"
kaijud query gov proposal $PROPOSAL_ID --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1

printf "\n#5) Query kaijus...\n\n"
kaijud query kaiju kaijus --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1

printf "\n#6) Waiting for gov proposal to pass...\n\n"
sleep 8

printf "\n#7) Query kaijus after proposal passed...\n\n"
kaijud query kaiju kaijus --node=tcp://3.75.187.158:26657 --chain-id=kaiju-testnet-1