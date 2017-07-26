package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoConfig stores the configuration of mongodb to connect
type MongoConfig struct {
	Ip       string `json:"ip"`
	Database string `json:"database"`
}

var mongoConfig MongoConfig

func readMongodbConfig(path string) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &mongoConfig)
}

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://" + mongoConfig.Ip)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
func getCollection(session *mgo.Session, collection string) *mgo.Collection {

	c := session.DB(mongoConfig.Database).C(collection)
	return c
}
func saveBlock(c *mgo.Collection, block BlockModel) {
	//first, check if the item already exists
	result := BlockModel{}
	err := c.Find(bson.M{"hash": block.Hash}).One(&result)
	if err != nil {
		//item not found, so let's add a new entry
		err = c.Insert(block)
		check(err)
	} else {
		err = c.Update(bson.M{"hash": block.Hash}, &block)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func getAllNodes() ([]NodeModel, error) {
	result := []NodeModel{}
	iter := nodeCollection.Find(bson.M{}).Limit(100).Iter()
	err := iter.All(&result)
	return result, err
}

func saveNode(c *mgo.Collection, node NodeModel) {
	//first, check if the node already exists
	result := NodeModel{}
	err := c.Find(bson.M{"id": node.Id}).One(&result)
	if err != nil {
		//node not found, so let's add a new entry
		err = c.Insert(node)
		check(err)
	} else {
		err = c.Update(bson.M{"id": node.Id}, &node)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getAllEdges() ([]EdgeModel, error) {
	result := []EdgeModel{}
	iter := edgeCollection.Find(bson.M{}).Limit(100).Iter()
	err := iter.All(&result)
	return result, err
}
func saveEdge(c *mgo.Collection, edge EdgeModel) {
	//first, check if the edge already exists
	result := EdgeModel{}
	err := c.Find(bson.M{"txid": edge.Txid}).One(&result)
	if err != nil {
		//edge not found, so let's add a new entry
		err = c.Insert(edge)
		check(err)
	} else {
		err = c.Update(bson.M{"txid": edge.Txid}, &edge)
		if err != nil {
			log.Fatal(err)
		}
	}
}
