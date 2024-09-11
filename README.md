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

## Confirm EthSigner passing requests to Besu : eth_blockNumber
Request the current block number using eth_blockNumber with the EthSigner JSON-RPC endpoint (8545 in this example):
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":51}' http://127.0.0.1:8545
```

## Check estimate gas for each function : eth_estimateGas
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_estimateGas","params":[{"from": ${your wallet address},"to": ${token contract address},"gas": "0xC350","gasPrice": "0xB2D05E00","data":"call data"}],"id":9}' http://127.0.0.1:8545
```
this function will help you to estimate how much gas should use in the function
Note : each method will use different gas usage
 - transferFrom = 45,000 units
 - name  = 24,000 uints

# TransferFrom token : eth_sendTransaction
```shell
curl -X POST --data '{"jsonrpc": "2.0","method": "eth_sendTransaction","params": [{"from": ${your wallet address},"to": ${token contract address},"gas": "0xC350","gasPrice": "0xB2D05E00","data":"call data"}],"id":9}' http://127.0.0.1:8545
```
Note you will not see destination address here, because it was encrpy to call data already