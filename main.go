package main

import (
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/btcsuite/btcrpcclient"
	"github.com/fatih/color"
	"github.com/gorilla/handlers"
)

var blockCollection *mgo.Collection
var nodeCollection *mgo.Collection
var edgeCollection *mgo.Collection
var dateCountCollection *mgo.Collection
var hourCountCollection *mgo.Collection

func main() {
	savelog()
	//read goBlockchainDataAbalysis config
	readConfig("config.json")

	//connect with mongodb
	readMongodbConfig("./mongodbConfig.json")
	session, err := getSession()
	check(err)
	blockCollection = getCollection(session, "blocks")
	nodeCollection = getCollection(session, "nodes")
	edgeCollection = getCollection(session, "edges")
	dateCountCollection = getCollection(session, "dateCounts")
	hourCountCollection = getCollection(session, "hourCounts")

	if len(os.Args) > 1 {
		if os.Args[1] == "-explore" {
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
			color.Blue("starting to explore blockchain")
			explore(client, config.GenesisBlock)

			// Get the current block count.
			blockCount, err := client.GetBlockCount()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Block count: %d", blockCount)
		}
	}

	//http server start
	readServerConfig("./serverConfig.json")
	log.Println("server running")
	log.Print("port: ")
	log.Println(serverConfig.ServerPort)
	router := NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+serverConfig.ServerPort, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	//log.Fatal(http.ListenAndServe(":"+serverConfig.ServerPort, router))

}
