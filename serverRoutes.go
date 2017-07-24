package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	/*	Route{
			"Recommendations",
			"GET",
			"/r/{userid}/{nrec}",
			Recommendations,
		},
	*/
	Route{
		"NetworkMap",
		"Get",
		"/map",
		NetworkMap,
	},
	/*
		Route{
			"SelectItem",
			"GET",
			"/selectItem/{userid}/{itemid}",
			SelectItem,
		},
	*/
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
