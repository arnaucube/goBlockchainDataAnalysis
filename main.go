package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/gorilla/handlers"
)

var statsCollection *mgo.Collection
var blockCollection *mgo.Collection
var txCollection *mgo.Collection
var addressCollection *mgo.Collection
var nodeCollection *mgo.Collection
var edgeCollection *mgo.Collection
var dateCountCollection *mgo.Collection
var hourCountCollection *mgo.Collection

func main() {
	savelog()
	//read goBlockchainDataAbalysis config
	readConfig("config.json")

	//connect with mongodb
	//readMongodbConfig("./mongodbConfig.json")
	session, err := getSession()
	check(err)
	statsCollection = getCollection(session, "stats")
	blockCollection = getCollection(session, "blocks")
	txCollection = getCollection(session, "txs")
	addressCollection = getCollection(session, "address")
	nodeCollection = getCollection(session, "nodes")
	edgeCollection = getCollection(session, "edges")
	dateCountCollection = getCollection(session, "dateCounts")
	hourCountCollection = getCollection(session, "hourCounts")

	if len(os.Args) > 1 {
		if os.Args[1] == "-explore" {
			// create new client instance
			client, err := rpcclient.New(&rpcclient.ConnConfig{
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
			log.Println("starting to explore blockchain")
			start := time.Now()
			explore(client, config.GenesisBlock)
			log.Println("blockchain exploration finished, time:")
			log.Println(time.Since(start))

			// Get the current block count.
			blockCount, err := client.GetBlockCount()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Block count: %d", blockCount)
		}
		if os.Args[1] == "-continue" {
			// create new client instance
			client, err := rpcclient.New(&rpcclient.ConnConfig{
				HTTPPostMode: true,
				DisableTLS:   true,
				Host:         config.Host + ":" + config.Port,
				User:         config.User,
				Pass:         config.Pass,
			}, nil)
			check(err)
			//get last block stored in mongodb
			lastBlock := BlockModel{}
			err = blockCollection.Find(bson.M{}).Sort("-$natural").One(&lastBlock)
			check(err)
			log.Println("Getting last block stored in MongoDB. Hash: " + string(lastBlock.Hash) + ", BlockHeight: " + strconv.FormatInt(lastBlock.Height, 10))
			log.Println("continuing blockchain exploration since last block in mongodb")
			start := time.Now()
			explore(client, string(lastBlock.Hash))
			log.Println("blockchain exploration finished, time:")
			log.Println(time.Since(start))
		}
	}
	//run thw webserver
	go webserver()

	//http server start
	//readServerConfig("./serverConfig.json")
	log.Println("server running")
	log.Print("port: ")
	log.Println(config.Server.ServerPort)
	router := NewRouter()

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	log.Fatal(http.ListenAndServe(":"+config.Server.ServerPort, handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	//log.Fatal(http.ListenAndServe(":"+serverConfig.ServerPort, router))

}

func webserver() {
	log.Println("webserver in port " + config.Server.WebServerPort)
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":"+config.Server.WebServerPort, nil)
}
