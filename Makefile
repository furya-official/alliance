#!/usr/bin/make -f

localnet-kaiju-rmi:
	docker rmi terra-money/localnet-kaiju 2>/dev/null; true

localnet-build-env: localnet-kaiju-rmi
	docker build --tag terra-money/localnet-kaiju -f scripts/containers/Dockerfile \
    		$(shell git rev-parse --show-toplevel)
	
localnet-build-nodes:
	docker run --rm -v $(CURDIR)/.testnets:/kaiju terra-money/localnet-kaiju \
		testnet init-files --v 3 -o /kaiju --starting-ip-address 192.168.5.20 --keyring-backend=test --chain-id=kaiju-testnet-1
	docker-compose up -d

localnet-stop:
	docker-compose down

localnet-start: localnet-stop localnet-build-env localnet-build-nodes

.PHONY: localnet-start localnet-stop localnet-build-env localnet-build-nodes
