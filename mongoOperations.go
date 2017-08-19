package main

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://" + config.Mongodb.IP)
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

	c := session.DB(config.Mongodb.Database).C(collection)
	return c
}

/*
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
*/

func getAllNodes() ([]NodeModel, error) {
	result := []NodeModel{}
	iter := nodeCollection.Find(bson.M{}).Limit(10000).Iter()
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
		/*err = c.Update(bson.M{"id": node.Id}, &node)
		if err != nil {
			log.Fatal(err)
		}
		*/
	}
}

func getAllEdges() ([]EdgeModel, error) {
	result := []EdgeModel{}
	iter := edgeCollection.Find(bson.M{}).Limit(10000).Iter()
	err := iter.All(&result)
	return result, err
}
func saveEdge(c *mgo.Collection, edge EdgeModel) {
	//first, check if the edge already exists
	result := EdgeModel{}
	err := c.Find(bson.M{"txid": edge.Txid, "to": edge.To, "from": edge.From, "blockheight": edge.BlockHeight, "label": edge.Label}).One(&result)
	if err != nil {
		//edge not found, so let's add a new entry
		err = c.Insert(edge)
		check(err)
	} else {
		err = c.Update(bson.M{"txid": edge.Txid, "to": edge.To, "from": edge.From, "blockheight": edge.BlockHeight, "label": edge.Label}, &edge)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func nodeInNodes(nodes []NodeModel, node NodeModel) bool {
	for _, n := range nodes {
		if n.Id == node.Id {
			return true
		}
	}
	return false
}
func edgeInEdges(edges []EdgeModel, edge EdgeModel) bool {
	for _, e := range edges {
		if e.From == edge.From && e.To == edge.To && e.Label == edge.Label && e.BlockHeight == edge.BlockHeight {
			return true
		}
	}
	return false
}
func saveAddress(address AddressModel) {

	result := AddressModel{}
	err := addressCollection.Find(bson.M{"hash": address.Hash}).One(&result)
	if err != nil {
		//address not found, so let's add a new entry
		err = addressCollection.Insert(address)
		check(err)
	}
}
func saveTx(tx TxModel) {

	result := TxModel{}
	err := txCollection.Find(bson.M{"txid": tx.Txid}).One(&result)
	if err != nil {
		//tx not found, so let's add a new entry
		err = txCollection.Insert(tx)
		check(err)
	}
}

func saveBlock(block BlockModel) {

	result := BlockModel{}
	err := blockCollection.Find(bson.M{"hash": block.Hash}).One(&result)
	if err != nil {
		//block not found, so let's add a new entry
		err = blockCollection.Insert(block)
		check(err)
	}
}
