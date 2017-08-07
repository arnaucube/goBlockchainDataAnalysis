package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
)

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AllAddresses",
		"Get",
		"/alladdresses",
		AllAddresses,
	},
	Route{
		"GetLastTx",
		"Get",
		"/lasttx",
		GetLastTx,
	},
	Route{
		"AddressNetwork",
		"GET",
		"/address/network/{address}",
		AddressNetwork,
	},
	Route{
		"AddressSankey",
		"GET",
		"/address/sankey/{address}",
		AddressSankey,
	},
	Route{
		"NetworkMap",
		"Get",
		"/map",
		NetworkMap,
	},
	Route{
		"GetHourAnalysis",
		"Get",
		"/houranalysis",
		GetHourAnalysis,
	},
}

//ROUTES

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ask for recommendations in /r")
	//http.FileServer(http.Dir("./web"))
}

/*
func NewUser(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)
	decoder := json.NewDecoder(r.Body)
	var newUser UserModel
	err := decoder.Decode(&newUser)
	check(err)
	defer r.Body.Close()

	saveUser(userCollection, newUser)

	fmt.Println(newUser)
	fmt.Fprintln(w, "new user added: ", newUser.ID)
}
*/
func AllAddresses(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	nodes := []NodeModel{}
	iter := nodeCollection.Find(bson.M{}).Limit(10000).Iter()
	err := iter.All(&nodes)

	//convert []resp struct to json
	jsonNodes, err := json.Marshal(nodes)
	check(err)

	fmt.Fprintln(w, string(jsonNodes))
}
func GetLastTx(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	nodes := []NodeModel{}
	err := nodeCollection.Find(bson.M{}).Limit(10).Sort("-$natural").All(&nodes)
	check(err)

	//convert []resp struct to json
	jNodes, err := json.Marshal(nodes)
	check(err)

	fmt.Fprintln(w, string(jNodes))
}
func AddressNetwork(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	vars := mux.Vars(r)
	address := vars["address"]
	if address == "undefined" {
		fmt.Fprintln(w, "not valid address")
	} else {
		network := addressTree(address)

		//convert []resp struct to json
		jNetwork, err := json.Marshal(network)
		check(err)

		fmt.Fprintln(w, string(jNetwork))
	}
}
func AddressSankey(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	vars := mux.Vars(r)
	address := vars["address"]
	if address == "undefined" {
		fmt.Fprintln(w, "not valid address")
	} else {
		network := addressTree(address)
		var sankey SankeyModel

		fmt.Println("network generated")
		mapNodeK := make(map[string]int)
		for k, n := range network.Nodes {
			var sankeyNode SankeyNodeModel
			//sankeyNode.StringNode = n.Id
			sankeyNode.Node = k
			sankeyNode.Name = n.Id
			sankey.Nodes = append(sankey.Nodes, sankeyNode)
			mapNodeK[n.Id] = k
		}
		for _, e := range network.Edges {
			var sankeyLink SankeyLinkModel
			//sankeyLink.StringSource = e.From
			sankeyLink.Source = mapNodeK[e.From]
			//sankeyLink.StringTarget = e.To
			sankeyLink.Target = mapNodeK[e.To]
			sankeyLink.Value = e.Label
			sankey.Links = append(sankey.Links, sankeyLink)
		}
		fmt.Println("Sankey generated")

		//convert []resp struct to json
		jsonSankey, err := json.Marshal(sankey)
		check(err)

		fmt.Fprintln(w, string(jsonSankey))
	}
}
func NetworkMap(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	nodes, err := getAllNodes()
	check(err)
	edges, err := getAllEdges()
	check(err)

	var network NetworkModel
	network.Nodes = nodes
	network.Edges = edges

	//convert []resp struct to json
	jNetwork, err := json.Marshal(network)
	check(err)

	fmt.Fprintln(w, string(jNetwork))
}

/*
func SelectItem(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)
	vars := mux.Vars(r)
	userid := vars["userid"]
	itemid := vars["itemid"]
	//find item
	item, err := getItemById(itemid)
	if err != nil {
		fmt.Fprintln(w, "item "+itemid+" not found")
	}

	//find user
	user, err := getUserById(userid)
	if err != nil {
		fmt.Fprintln(w, "user "+userid+" not found")
	}

	//increase TActed in item
	item.TActed = item.TActed + 1

	//save item
	item, err = updateItem(item)
	check(err)
	fmt.Println(item)

	//add item to []Actions of user
	user.Actions = append(user.Actions, itemid)

	//save user
	user, err = updateUser(user)
	check(err)
	fmt.Println(user)

	fmt.Fprintln(w, "user: "+userid+", selects item: "+itemid)
}
*/
func GetHourAnalysis(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	hourAnalysis := []HourCountModel{}
	iter := hourCountCollection.Find(bson.M{}).Limit(10000).Iter()
	err := iter.All(&hourAnalysis)

	//sort by hour
	sort.Slice(hourAnalysis, func(i, j int) bool {
		return hourAnalysis[i].Hour < hourAnalysis[j].Hour
	})

	var resp HourAnalysisResp
	for _, d := range hourAnalysis {
		resp.Labels = append(resp.Labels, d.Hour)
		resp.Data = append(resp.Data, d.Count)
	}

	//convert []resp struct to json
	jsonResp, err := json.Marshal(resp)
	check(err)
	fmt.Fprintln(w, string(jsonResp))
}
