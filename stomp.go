package main

import (
	"fmt"

	"github.com/go-stomp/stomp"
)

// these are the default options that work with RabbitMQ
var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	stomp.ConnOpt.Login("guest", "guest"),
	stomp.ConnOpt.Host("/"),
}

func sendMessages(conn *stomp.Conn) {
	defer func() {
		stop <- true
	}()

	// conn, err := stomp.Dial("tcp", *serverAddr, options...)
	// if err != nil {
	// 	println("cannot connect to server", err.Error())
	// 	return
	// }

	for i := 1; i <= *messageCount; i++ {
		text := fmt.Sprintf("Message #%d", i)
		err := conn.Send(*queueName, "text/plain",
			[]byte(text), nil)
		if err != nil {
			println("failed to send to server", err)
			return
		}
	}
	println("sender finished")
}

func recvMessages(conn *stomp.Conn, subscribed chan bool) {
	defer func() {
		stop <- true
	}()

	// conn, err := stomp.Dial("tcp", *serverAddr, options...)

	// if err != nil {
	// 	println("cannot connect to server", err.Error())
	// 	return
	// }

	sub, err := conn.Subscribe(*queueName, stomp.AckAuto)
	if err != nil {
		println("cannot subscribe to", *queueName, err.Error())
		return
	}
	close(subscribed)

	for i := 1; i <= *messageCount; i++ {
		msg := <-sub.C
		expectedText := fmt.Sprintf("Message #%d", i)
		actualText := string(msg.Body)
		if expectedText != actualText {
			println("Expected:", expectedText)
			println("Actual:", actualText)
		}
	}
	println("receiver finished")

}
