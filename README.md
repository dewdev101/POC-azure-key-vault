# How to call transaction in block chain

# Prepare state
```shell
 fill secret in .env and docker-compose.yml
```
```shell
 docker compose up -d
```
# Get call data
 encrypt method to use when call eth_sendTransaction in eth blockchain
```shell
 go run .
```
we call transferFrom function
(you will see the log Call data in terminal)

# Confirm EthSigner is up
```shell
curl -X GET http://127.0.0.1:8545/upcheck
```

## Confirm EthSigner passing requests to Besu
## Request the current block number using eth_blockNumber with the EthSigner JSON-RPC endpoint (8545 in this example):
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":51}' http://127.0.0.1:8545
```

## Check estimate gas for each function 
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_estimateGas","params":[{"from": ${your wallet address},"to": ${token contract address},"gas": "0xC350","gasPrice": "0xB2D05E00","data":"call data"}],"id":9}' http://127.0.0.1:8545
```
Note : each method will use different gas usage
 - transferFroms about 45,000 gas units