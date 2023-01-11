#!/bin/bash

printf "#1) Submit proposal to create ufuryx Furya...\n\n"
furyad tx gov submit-legacy-proposal create-furya ufuryx 0.5 0.00005 1 0 --from=demowallet1 --home ./data/furya --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

PROPOSAL_ID=$(furyad query gov proposals --count-total --output json --home ./data/furya | jq .pagination.total -r)

printf "\n#2) Deposit funds to proposal $PROPOSAL_ID...\n\n"
furyad tx gov deposit $PROPOSAL_ID 10000000stake --from=demowallet1 --home ./data/furya --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

printf "\n#3) Vote to pass the proposal...\n\n"
furyad tx gov vote $PROPOSAL_ID yes --from=val1 --home ./data/furya --keyring-backend=test --broadcast-mode=block --gas 1000000 -y > /dev/null 2>&1

printf "\n#4) Query proposals...\n\n"
furyad query gov proposal $PROPOSAL_ID --home ./data/furya

printf "\n#5) Query furyas...\n\n"
furyad query furya furyas --home ./data/furya

printf "\n#6) Waiting for gov proposal to pass...\n\n"
sleep 8

printf "\n#7) Query furyas after proposal passed...\n\n"
furyad query furya furyas --home ./data/furya