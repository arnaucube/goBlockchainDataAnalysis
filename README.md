# goBlockchainDataAnalysis
blockchain data analysis, written in Go

#### Not finished - ToDo list
To Do list in DevelopmentNotes.md https://github.com/arnaucode/goBlockchainDataAnalysis/blob/master/DevelopmentNotes.md

## Instructions

### Install
1. Nodejs & NPM https://nodejs.org/ --> to get npm packages for the web
2. MongoDB https://www.mongodb.com/
3. Faircoin wallet https://download.faircoin.world/, or the Cryptocurrency desired wallet
4. goBlockchainDataAnalysis https://github.com/arnaucode/goBlockchainDataAnalysis

### Configure
- Wallet /home/user/.faircoin2/faircoin.conf:
```
rpcuser=faircoinrpc
rpcpassword=password
rpcport=3021
rpcworkqueue=2000
server=1
rpcbind=127.0.0.1
rpcallowip=127.0.0.1
```

- goBlockchainDataAnalysis/config.json:
```json
{
    "user": "faircoinrpc",
    "pass": "password",
    "host": "127.0.0.1",
    "port": "3021",
	"genesisTx": "7c27ade2c28e67ed3077f8f77b8ea6d36d4f5eba04c099be3c9faa9a4a04c046",
	"genesisBlock": "beed44fa5e96150d95d56ebd5d2625781825a9407a5215dd7eda723373a0a1d7",
    "startFromBlock": 0,
    "server": {
        "serverIP": "127.0.0.1",
        "serverPort": "3014",
        "webServerPort": "8080",
        "allowedIPs": [
            "127.0.0.1"
        ],
        "blockedIPs": []
    },
    "mongodb": {
        "ip": "127.0.0.1",
        "database": "goBlockchainDataAnalysis"
    }
}
```

### Run

1. Start MongoDB
```
sudo service mongod start
```

2. Start wallet
```
./faircoind -txindex -reindex-chainstate
```
Wait until the entire blockchain is downloaded.

3. Run explorer, to fill the database
```
./goBlockchainDataAnalysis -explore
```

    3.1. The next runs, once the database have data and just need to add last blocks added in the blockchain, can just run:
```
./goBlockchainDataAnalysis -continue
```

    3.2. If don't want to fill the database, can just run:
```
./goBlockchainDataAnalysis
```

Webapp will run on 127.0.0.1:8080

4. ADDITIONAL - Run the webserver, directly from the /web directory
This can be useful if need to deploy the API server in one machine and the webserver in other.
In the /web directory:
```
npm start
```
Webapp will run on 127.0.0.1:8080



### Additional info
- Backend
    - Go lang https://golang.org/
    - MongoDB https://www.mongodb.com/
- Frontend
    - AngularJS https://angularjs.org/
    - Angular-Bootstrap-Material https://tilwinjoy.github.io/angular-bootstrap-material


### Some screenshots
Some screenshots can be old, and can contain errors.

![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/screenshots/goBlockchainDataAnalysis00.png "goBlockchainDataAnalysis")

![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/screenshots/goBlockchainDataAnalysis01.png "goBlockchainDataAnalysis")

![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/screenshots/goBlockchainDataAnalysis02.png "goBlockchainDataAnalysis")

![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/screenshots/goBlockchainDataAnalysis03.gif "goBlockchainDataAnalysis")
