package main

import (
	"fmt"
	"log"

	"github.com/fatih/color"

	"gopkg.in/mgo.v2/bson"
)

func upTree(address string, network NetworkModel) NetworkModel {
	var upNetwork NetworkModel

	//add address node to network
	node := NodeModel{}
	err := nodeCollection.Find(bson.M{"id": address}).One(&node)
	check(err)
	if nodeInNodes(network.Nodes, node) == false {
		network.Nodes = append(network.Nodes, node)
	}

	//get edges before the address
	edges := []EdgeModel{}
	err = edgeCollection.Find(bson.M{"to": address}).All(&edges)
	check(err)
	for _, e := range edges {
		if edgeInEdges(network.Edges, e) == false {
			network.Edges = append(network.Edges, e)
		}
		//get the nodes involved in edges
		/*
			nodes := []NodeModel{}
			err := nodeCollection.Find(bson.M{"id": }).All(&edges)
			check(err)
		*/
		endBranch := false
		edgeUpCheck := EdgeModel{}
		err := edgeCollection.Find(bson.M{"to": e.From}).One(&edgeUpCheck)
		if err != nil {
			log.Println(err)
			color.Blue("not found")
			endBranch = true
		}

		//need to be fixed when there is a bucle between the addresses (A-->B, B-->C, C-->A)
		fmt.Println(e.From + " - " + e.To)
		if edgeInEdges(network.Edges, edgeUpCheck) == false && endBranch == false && edgeUpCheck.BlockHeight <= e.BlockHeight && e.To != edgeUpCheck.From {
			upNetwork = upTree(e.From, network)
			for _, upN := range upNetwork.Nodes {
				if nodeInNodes(network.Nodes, upN) == false {
					network.Nodes = append(network.Nodes, upN)
				}
			}
			for _, upE := range upNetwork.Edges {
				if edgeInEdges(network.Edges, upE) == false {
					network.Edges = append(network.Edges, upE)
				}
			}
		}
	}

	return network
}
func addressTree(address string) NetworkModel {
	var network NetworkModel

	network = upTree(address, network)
	return network

}
