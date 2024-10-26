run-chain:
	kurtosis run github.com/ethpandaops/ethereum-package --args-file ./chain/network_params.yml --image-download always --enclave test-chain

stop-chain:
	kurtosis enclave stop test-chain

remove-chain:
	kurtosis enclave rm test-chain

run-service:
	docker compose up -d

connect-containers:
	@containers=$$(docker ps --filter "name=el-1-geth-teku-" --format "{{.ID}}"); \
	for container in $$containers; do \
		new_name="e1-1-geth"; \
		echo "Renaming container $$container to $$new_name"; \
		docker rename $$container $$new_name; \
		echo "Connecting container $$new_name to network go-eth_bridge-network"; \
		docker network connect go-eth_bridge-network $$new_name; \
	done

stop-service:
	docker compose down