package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	User           string `json:"user"`
	Pass           string `json:"pass"`
	Host           string `json:"host"`
	Port           string `json:"port"`
	GenesisTx      string `json:"genesisTx"`
	GenesisBlock   string `json:"genesisBlock"`
	StartFromBlock int64  `json:"startFromBlock"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
