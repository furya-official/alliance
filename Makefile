#!/usr/bin/make -f

localnet-furya-rmi:
	docker rmi furya-official/localnet-furya 2>/dev/null; true

localnet-build-env: localnet-furya-rmi
	docker build --tag furya-official/localnet-furya -f scripts/containers/Dockerfile \
    		$(shell git rev-parse --show-toplevel)
	
localnet-build-nodes:
	docker run --rm -v $(CURDIR)/.testnets:/furya furya-official/localnet-furya \
		testnet init-files --v 3 -o /furya --starting-ip-address 192.168.5.20 --keyring-backend=test --chain-id=furya-testnet-1
	docker-compose up -d

localnet-stop:
	docker-compose down

localnet-start: localnet-stop localnet-build-env localnet-build-nodes

.PHONY: localnet-start localnet-stop localnet-build-env localnet-build-nodes
