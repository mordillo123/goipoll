package main

import (
	"encoding/json"
	"log"
	"os"
)

const configFile = "config.json"

// Configuration contain
// Addess = Socket URL
// Topics = STOMP Topic to subscribe
type Configuration struct {
	Address string
	Topics  []string
}

// ReadConf Open the default configuration file and parse the conten on Configuration structure
func ReadConf() (Configuration, error) {
	log.Printf("Open Configuration on configfile %s", configFile)

	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Printf("No Configuration found error: %s", err)
		return conf, err
	}
	log.Printf("Configuration achieved.")

	return conf, err

}
