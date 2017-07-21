package main

import (
	"log"

	"github.com/btcsuite/btcrpcclient"
)

func main() {

	readConfig("config.json")

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

	explore(client, config.GenesisBlock)

	// Get the current block count.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Block count: %d", blockCount)
}
