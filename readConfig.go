package main

import (
	"encoding/json"
	"io/ioutil"
)

//MongoConfig stores the configuration of mongodb to connect
type MongoConfig struct {
	IP       string `json:"ip"`
	Database string `json:"database"`
}

//ServerConfig reads the server configuration
type ServerConfig struct {
	ServerIP      string   `json:"serverIP"`
	ServerPort    string   `json:"serverPort"`
	WebServerPort string   `json:"webserverPort"`
	AllowedIPs    []string `json:"allowedIPs"`
	BlockedIPs    []string `json:"blockedIPs"`
}

//Config reads the config
type Config struct {
	User           string       `json:"user"`
	Pass           string       `json:"pass"`
	Host           string       `json:"host"`
	Port           string       `json:"port"`
	GenesisTx      string       `json:"genesisTx"`
	GenesisBlock   string       `json:"genesisBlock"`
	StartFromBlock int64        `json:"startFromBlock"`
	Server         ServerConfig `json:"server"`
	Mongodb        MongoConfig  `json:"mongodb"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}

/*
var mongoConfig MongoConfig

func readMongodbConfig(path string) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &mongoConfig)
}

var serverConfig ServerConfig

func readServerConfig(path string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("error: ", err)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &serverConfig)
}
*/
