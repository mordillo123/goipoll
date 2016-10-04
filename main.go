package main

import (
	"log"

	sockjs "github.com/lavab/sockjs-go-client"
	// "github.com/gmallard/stompngo"
)

func main() {
	log.Printf("Start")

	conf, err := ReadConf()

	if err != nil {
		return
	}

	ws, wserr := sockjs.NewClient(conf.Address)

	if wserr != nil {
		log.Printf("Error open WebSocket SJS %s", wserr)
	}

	wsinfo, wserr := ws.Info()

	log.Printf("Info SockJS WebSocket: %t\n", wsinfo.WebSocket)

}
