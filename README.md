# goBlockchainDataAnalysis
blockchain data analysis, written in Go

#### Not finished - ToDo list
- Backend
    - Network Address generation avoiding infinite relation loops
    - Sankey Address generation without loops
- Frontend
    - After Sankey visualization, go to Network Address visualization and render without Sankey dots


### Install
1. Nodejs & NPM https://nodejs.org/ --> to serve the web, not necessary if the web files are in a webserver
2. MongoDB https://www.mongodb.com/
3. Faircoin wallet https://download.faircoin.world/, or the Cryptocurrency desired wallet
4. goBlockchainDataAnalysis https://github.com/arnaucode/goBlockchainDataAnalysis

### Configure
- Wallet /home/user/.faircoin2/faircoin.conf:
```
rpcuser=usernamerpc
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
    "user": "usernamerpc",
    "pass": "password",
    "host": "127.0.0.1",
    "port": "3021",
    "genesisTx": "7c27ade2c28e67ed3077f8f77b8ea6d36d4f5eba04c099be3c9faa9a4a04c046",
    "genesisBlock": "beed44fa5e96150d95d56ebd5d2625781825a9407a5215dd7eda723373a0a1d7"
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

    3.1. The next runs, once the database have data, can just run:
```
./goBlockchainDataAnalysis
```

4. Run the webserver, in the /web directory
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

![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/goBlockchainDataAnalysis00.png "goBlockchainDataAnalysis")

![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/goBlockchainDataAnalysis06.png "goBlockchainDataAnalysis")

![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/goBlockchainDataAnalysis05.png "goBlockchainDataAnalysis")


![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/goBlockchainDataAnalysis01.png "goBlockchainDataAnalysis")


![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/goBlockchainDataAnalysis02.png "goBlockchainDataAnalysis")


![goBlockchainDataAnalysis](https://raw.githubusercontent.com/arnaucode/goBlockchainDataAnalysis/master/goBlockchainDataAnalysis04.png "goBlockchainDataAnalysis")
