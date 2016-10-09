package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-stomp/stomp"
	sockjs "github.com/mordillo123/sockjs-go-client"
)

// const defaultPort = ":61613"

// var serverAddr = flag.String("server", "localhost:61613", "STOMP server endpoint")
var messageCount = flag.Int("count", 10, "Number of messages to send/receive")
var queueName = flag.String("queue", "/topic/start", "Start Topic queue")
var helpFlag = flag.Bool("help", false, "Print help text")
var stop = make(chan bool)

func main() {
	flag.Parse()

	if *helpFlag {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}

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

	st, sterr := stomp.Connect(ws)

	if sterr != nil {
		log.Printf("Error switching to STOMP Connection %s", sterr)
	}

	subscribed := make(chan bool)
	go recvMessages(st, subscribed)

	// wait until we know the receiver has subscribed
	<-subscribed

	go sendMessages(st)

	<-stop
	<-stop

}
