package main

import (
	"encoding/json"
	"log"
	"os"

	sockjs "github.com/lavab/sockjs-go-client"
)

const configFile = "config.json"

// Configuration contain
// Addess = Socket URL
// Topics = STOMP Topic to subscribe
type Configuration struct {
	Address string
	Topics  []string
}

func main() {
	log.Printf("Start")

	log.Printf("Read configfile")

	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		log.Printf("error: %s", err)
	}
	log.Println(conf.Address)

	ws, wserr := sockjs.NewClient(conf.Address)

	if wserr != nil {
		log.Printf("Error open WebSocket SJS %s", wserr)
	}

	wsinfo, wserr := ws.Info()

	log.Printf("Info: %t\n", wsinfo.WebSocket)

}
