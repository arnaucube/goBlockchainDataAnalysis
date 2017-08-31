package main

import (
	"log"
	"strconv"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
	"gopkg.in/mgo.v2/bson"
)

func explorationContinue() {
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

func continuousExploration() {
	for {
		explorationContinue()
		time.Sleep(time.Second * 60)
	}
}
