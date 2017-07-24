package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/btcsuite/btcrpcclient"
	"github.com/fatih/color"
)

var blockCollection *mgo.Collection
var nodeCollection *mgo.Collection
var edgeCollection *mgo.Collection

func main() {
	//read goBlockchainDataAbalysis config
	readConfig("config.json")

	//connect with mongodb
	readMongodbConfig("./mongodbConfig.json")
	session, err := getSession()
	check(err)
	blockCollection = getCollection(session, "blocks")
	nodeCollection = getCollection(session, "nodes")
	edgeCollection = getCollection(session, "edges")

	// create new client instance
	client, err := btcrpcclient.New(&btcrpcclient.ConnConfig{
		HTTPPostMode: true,
		DisableTLS:   true,
		Host:         config.Host + ":" + config.Port,
		User:         config.User,
		Pass:         config.Pass,
	}, nil)
	if err != nil {
		log.Fatalf("error creating new btc client: %v", err)
	}

	// list accounts
	accounts, err := client.ListAccounts()
	if err != nil {
		log.Fatalf("error listing accounts: %v", err)
	}
	// iterate over accounts (map[string]btcutil.Amount) and write to stdout
	for label, amount := range accounts {
		log.Printf("%s: %s", label, amount)
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "-explore" {
			fmt.Println("starting to explore blockchain")
			explore(client, config.GenesisBlock)
		}
	}

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)

	//http server start
	readServerConfig("./serverConfig.json")
	color.Green("server running")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":"+serverConfig.ServerPort, router))

}
