# Calling Transactions on the Blockchain

## 1. Prepare the Environment

1. **Update Secrets**
   - Fill in the required secrets in your `.env` and `docker-compose.yml` files.

2. **Start Services**
   - Launch the Docker containers:
     ```shell
     docker compose up -d
     ```

## 2. Generate Call Data
To generate call data for the `transferFrom` function:
   ```shell
   go run .
```
(you will see the log Call data in terminal)

## 3. Confirm EthSigner is up
```shell
curl -X GET http://127.0.0.1:8545/upcheck
```

## 4. Confirm EthSigner passing requests to Besu : eth_blockNumber
Request the current block number using eth_blockNumber with the EthSigner JSON-RPC endpoint (8545 in this example):
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":51}' http://127.0.0.1:8545
```

## 5. Check estimate gas for each function : eth_estimateGas
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_estimateGas","params":[{"from": "`your_wallet_address`","to":  "`token_contract_address`","gas": "0xC350","gasPrice": "0xB2D05E00","data":"call data"}],"id":9}' http://127.0.0.1:8545
```
This function will help you to estimate how much gas should use in the function
Note : each method will use different gas usage
 - transferFrom = 45,000 units
 - name  = 24,000 uints

# 6. TransferFrom token : eth_sendTransaction
```shell
curl -X POST --data '{"jsonrpc": "2.0","method": "eth_sendTransaction","params": [{"from": ${your wallet address},"to": ${token contract address},"gas": "0xC350","gasPrice": "0xB2D05E00","data":"call data"}],"id":9}' http://127.0.0.1:8545
```
Note you will not see destination address here, because it was encrpy to call data already