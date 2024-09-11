# POC Azure key vault

## Table of Contents

- [API_Instruction](#API)
- [Call_block_chain](#Calling_Transactions_on_the_Blockchain)

## API

### Set up 

1. **Update Secrets**
   
   - Fill in the required secrets in your `.env`  files.

2. **Start Services**
   
   - Launch api:
     ```shell
     go run .
     ```
  
### How to get secret value

#### 1 . Get token for getting secret

```shell
  http://localhost:8080/token/secret
```
#### 2. Get secrete value

 - Fill token that you got from 1 and fill in the Authorization (Bearer Token)
 - Call api
```shell
 http://localhost:8080/secret
```  
Note : after you run this api, it will generate keyfile and save passwordFile in ether-singer folder
which is used to sign in blockchain
=====================================================================

## Calling_Transactions_on_the_Blockchain

### 1. Prepare the Environment

1. **Update Secrets**
   
   - Fill in the required secrets in your `.env` and `docker-compose.yml` files.

2. **Start Services**

   - Launch the Docker containers:
     ```shell
     docker compose up -d
     ```

### 2. Generate Call Data

To generate call data for the `transferFrom` function:
   ```shell
   go run .
```
(you will see the log Call data in terminal)

### 3. Confirm EthSigner is up

```shell
curl -X GET http://127.0.0.1:8545/upcheck
```

### 4. Confirm EthSigner passing requests to Besu : eth_blockNumber

Request the current block number using eth_blockNumber with the EthSigner JSON-RPC endpoint (8545 in this example):
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":51}' http://127.0.0.1:8545
```

### 5. Check estimate gas for each function : eth_estimateGas

```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_estimateGas","params":[{"from": "`your_wallet_address`","to":  "`token_contract_address`","gas": "0xC350","gasPrice": "0xB2D05E00","data":"call data"}],"id":9}' http://127.0.0.1:8545
```

example
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_estimateGas","params":[{"from": "0xe8f9f81cb78f6096b10515d9d2675xxxxxxx","to": "0x402349046AA8F7e6fC355E7c8xxxxxxxxx","gas": "0xC350","gasPrice": "0xB2D05E00","data":"0x23b872dd000000000000000000000000e8f9f81cb78f6096b10515d9d26750ebfeaffd5d0000000000000000000000000e792a695b2aee2a49f654a219bdfc1c4381fbc20000000000000000000000000000000000000000000000008ac7230489e80000"}],"id":9}' http://127.0.0.1:8545
```

This function will help you to estimate how much gas should use in the function  
Note : each method will use different gas usage
 - transferFrom = 45,000 units
 - name  = 24,000 uints

### 6. TransferFrom token : eth_sendTransaction

```shell
curl -X POST --data '{"jsonrpc": "2.0","method": "eth_sendTransaction","params": [{"from": "`your_wallet_address`","to": "`token_contract_address`","gas": "0xC350","gasPrice": "0xB2D05E00","data":"call data"}],"id":9}' http://127.0.0.1:8545
```
Note you will not see destination address here, because it was encrypt to call data already

example
```shell
curl -X POST --data '{"jsonrpc":"2.0","method":"eth_sendTransaction","params":[{"from": "0xe8f9f81cb78f6096b10515d9d2675xxxxxxx","to": "0x402349046AA8F7e6fC355E7c8xxxxxxxxx","gas": "0xC350","gasPrice": "0xB2D05E00","data":"xxxxxxx000000000000000000000000e8f9f81cb78f6096b10515d9d26xxxxxx0000000000000000000000000e792a695b2aee2a49f654a219bdfc1c4381fbc2000000000000000000000000000000000000000000000000xxxxxxx"}],"id":9}' http://127.0.0.1:8545
```