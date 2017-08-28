package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"time"

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
		"Stats",
		"Get",
		"/stats",
		Stats,
	},
	Route{
		"AllAddresses",
		"Get",
		"/alladdresses",
		AllAddresses,
	},
	Route{
		"GetLastAddr",
		"Get",
		"/lastaddr",
		GetLastAddr,
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
		"Address",
		"GET",
		"/address/{hash}",
		Address,
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
		"GetTotalHourAnalysis",
		"Get",
		"/totalhouranalysis",
		GetTotalHourAnalysis,
	},
	Route{
		"GetLast24HourAnalysis",
		"Get",
		"/last24hour",
		GetLast24HourAnalysis,
	},
	Route{
		"GetLast7DayAnalysis",
		"Get",
		"/last7day",
		GetLast7DayAnalysis,
	},
	Route{
		"GetLast7DayHourAnalysis",
		"Get",
		"/last7dayhour",
		GetLast7DayHourAnalysis,
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
func Stats(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	stats := getStats()

	jsonResp, err := json.Marshal(stats)
	check(err)

	fmt.Fprintln(w, string(jsonResp))
}
func AllAddresses(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	nodes := []NodeModel{}
	iter := nodeCollection.Find(bson.M{"type": "address"}).Limit(10000).Iter()
	err := iter.All(&nodes)

	//convert []resp struct to json
	jsonNodes, err := json.Marshal(nodes)
	check(err)

	fmt.Fprintln(w, string(jsonNodes))
}
func GetLastAddr(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	addresses := []AddressModel{}
	err := addressCollection.Find(bson.M{}).Limit(10).Sort("-$natural").All(&addresses)
	check(err)

	//convert []resp struct to json
	jsonResp, err := json.Marshal(addresses)
	check(err)

	fmt.Fprintln(w, string(jsonResp))
}
func GetLastTx(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	txs := []TxModel{}
	err := txCollection.Find(bson.M{}).Limit(10).Sort("-$natural").All(&txs)
	check(err)

	//convert []resp struct to json
	jsonData, err := json.Marshal(txs)
	check(err)

	fmt.Fprintln(w, string(jsonData))
}
func AddressNetwork(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	vars := mux.Vars(r)
	address := vars["address"]
	if address == "undefined" {
		fmt.Fprintln(w, "not valid address")
	} else {
		network := addressTree(address)
		network.Nodes[0].Shape = "triangle"

		//convert []resp struct to json
		jNetwork, err := json.Marshal(network)
		check(err)

		fmt.Fprintln(w, string(jNetwork))
	}
}
func Address(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	vars := mux.Vars(r)
	hash := vars["hash"]
	if hash == "undefined" {
		fmt.Fprintln(w, "not valid hash")
	} else {
		address := AddressModel{}
		err := addressCollection.Find(bson.M{"hash": hash}).One(&address)

		txs := []TxModel{}
		err = txCollection.Find(bson.M{"$or": []bson.M{bson.M{"from": hash}, bson.M{"to": hash}}}).All(&txs)
		address.Txs = txs

		for _, tx := range address.Txs {
			blocks := []BlockModel{}
			err = blockCollection.Find(bson.M{"hash": tx.BlockHash}).All(&blocks)
			for _, block := range blocks {
				address.Blocks = append(address.Blocks, block)
			}
		}

		//convert []resp struct to json
		jsonResp, err := json.Marshal(address)
		check(err)

		fmt.Fprintln(w, string(jsonResp))
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

func GetTotalHourAnalysis(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	hourAnalysis := []ChartCountModel{}
	iter := hourCountCollection.Find(bson.M{}).Limit(10000).Iter()
	err := iter.All(&hourAnalysis)

	//sort by hour
	sort.Slice(hourAnalysis, func(i, j int) bool {
		return hourAnalysis[i].Elem < hourAnalysis[j].Elem
	})

	var resp ChartAnalysisResp
	for _, d := range hourAnalysis {
		resp.Labels = append(resp.Labels, strconv.Itoa(d.Elem))
		resp.Data = append(resp.Data, d.Count)
	}

	//convert []resp struct to json
	jsonResp, err := json.Marshal(resp)
	check(err)
	fmt.Fprintln(w, string(jsonResp))
}

func GetLast24HourAnalysis(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	fromDate := time.Now().AddDate(0, 0, -1)
	toDate := time.Now()

	txs := []TxModel{}
	err := txCollection.Find(bson.M{
		"datet": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
	}).Sort("-$natural").All(&txs)
	check(err)

	//generate map with 24 hours
	hourFrequencies := map24hours()
	for _, tx := range txs {
		hourFrequencies[tx.Date.Hour]++
	}
	var hourCount []ChartCountModel
	for hour, frequency := range hourFrequencies {
		hourCount = append(hourCount, ChartCountModel{hour, frequency})
	}
	//sort by hour
	sort.Slice(hourCount, func(i, j int) bool {
		return hourCount[i].Elem < hourCount[j].Elem
	})

	var resp ChartAnalysisResp
	for _, d := range hourCount {
		resp.Labels = append(resp.Labels, strconv.Itoa(d.Elem))
		resp.Data = append(resp.Data, d.Count)
	}

	//convert []resp struct to json
	jsonResp, err := json.Marshal(resp)
	check(err)
	fmt.Fprintln(w, string(jsonResp))
}
func GetLast7DayAnalysis(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	fromDate := time.Now().AddDate(0, 0, -7)
	toDate := time.Now()

	txs := []TxModel{}
	err := txCollection.Find(bson.M{
		"datet": bson.M{
			"$gt": fromDate,
			"$lt": toDate,
		},
	}).Sort("-$natural").All(&txs)
	check(err)

	//generate map with 24 hours
	//hourFrequencies := map24hours()
	dayFrequencies := make(map[int]int)
	for _, tx := range txs {
		dayFrequencies[tx.Date.Day]++
	}
	var dayCount []ChartCountModel
	for day, frequency := range dayFrequencies {
		dayCount = append(dayCount, ChartCountModel{day, frequency})
	}
	//sort by hour
	sort.Slice(dayCount, func(i, j int) bool {
		return dayCount[i].Elem < dayCount[j].Elem
	})

	var resp ChartAnalysisResp
	for _, d := range dayCount {
		resp.Labels = append(resp.Labels, strconv.Itoa(d.Elem))
		resp.Data = append(resp.Data, d.Count)
	}

	//convert []resp struct to json
	jsonResp, err := json.Marshal(resp)
	check(err)
	fmt.Fprintln(w, string(jsonResp))
}

func GetLast7DayHourAnalysis(w http.ResponseWriter, r *http.Request) {
	ipFilter(w, r)

	var resp ChartSeriesAnalysisResp

	for i := 0; i < 7; i++ {
		fromDate := time.Now().AddDate(0, 0, -i-1)
		toDate := time.Now().AddDate(0, 0, -i)
		txs := []TxModel{}
		err := txCollection.Find(bson.M{
			"datet": bson.M{
				"$gt": fromDate,
				"$lt": toDate,
			},
		}).Sort("-$natural").All(&txs)
		check(err)

		//generate map with 24 hours
		hourFrequencies := map24hours()
		for _, tx := range txs {
			hourFrequencies[tx.Date.Hour]++
		}
		var hourCount []ChartCountModel
		for hour, frequency := range hourFrequencies {
			hourCount = append(hourCount, ChartCountModel{hour, frequency})
		}
		//sort by hour
		sort.Slice(hourCount, func(i, j int) bool {
			return hourCount[i].Elem < hourCount[j].Elem
		})
		var dayData []int
		for _, d := range hourCount {
			dayData = append(dayData, d.Count)
		}
		if len(txs) > 0 {
			resp.Series = append(resp.Series, txs[0].Date.Day)
			resp.Data = append(resp.Data, dayData)
		}
	}
	hourLabels := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23"}
	resp.Labels = hourLabels

	//convert []resp struct to json
	jsonResp, err := json.Marshal(resp)
	check(err)
	fmt.Fprintln(w, string(jsonResp))
}
