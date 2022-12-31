#!/bin/bash

printf "#1) Submit proposal to create ukaijux Kaiju...\n\n"
kaijud tx gov submit-legacy-proposal create-kaiju ukaijux 0.5 0.00005 1 0 --from=demowallet1 --home ./data/kaiju --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

PROPOSAL_ID=$(kaijud query gov proposals --count-total --output json --home ./data/kaiju | jq .pagination.total -r)

printf "\n#2) Deposit funds to proposal $PROPOSAL_ID...\n\n"
kaijud tx gov deposit $PROPOSAL_ID 10000000stake --from=demowallet1 --home ./data/kaiju --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

printf "\n#3) Vote to pass the proposal...\n\n"
kaijud tx gov vote $PROPOSAL_ID yes --from=val1 --home ./data/kaiju --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

printf "\n#4) Query proposals...\n\n"
kaijud query gov proposal $PROPOSAL_ID --home ./data/kaiju

printf "\n#5) Query kaijus...\n\n"
kaijud query kaiju kaijus --home ./data/kaiju

printf "\n#6) Waiting for gov proposal to pass...\n\n"
sleep 8

printf "\n#7) Query kaijus after proposal passed...\n\n"
kaijud query kaiju kaijus --home ./data/kaiju