# Test call transaction in block chain

# Prepare state
```shell
 fill secret in .env and docker-compose.yml
```
```shell
 docker compose up -d
```

# Confirm EthSigner is up
```shell
curl -X GET http://127.0.0.1:8545/upcheck
```

## Confirm EthSigner passing requests to Besu
## Request the current block number using eth_blockNumber with the EthSigner JSON-RPC endpoint (8545 in this example):
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":51}' http://127.0.0.1:8545
```
