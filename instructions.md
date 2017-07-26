# Instructions

#### Install FairCoin Wallet and configure
now, need to configure wallet:
.FairCoin/
	FairCoin.conf

```
rpcuser=usernamerpc
rpcpassword=password
rpcport=3021
rpcworkqueue=2000
server=1
rpcbind=127.0.0.1
rpcallowip=127.0.0.1
```

execute wallet:
./faircoind -txindex -reindex-chainstate

### Configure

```json
{
    "user": "usernamerpc",
    "pass": "password",
    "host": "127.0.0.1",
    "port": "3021",
    "genesisTx": "7c27ade2c28e67ed3077f8f77b8ea6d36d4f5eba04c099be3c9faa9a4a04c046",
    "genesisBlock": "beed44fa5e96150d95d56ebd5d2625781825a9407a5215dd7eda723373a0a1d7"
}
```
