package main

import (
	"fmt"
	"log"

	"go-cqrs/messaging"
	"go-cqrs/util"
)

func main() {
	defer messaging.Close()

	// Connect to Nats
	addrNATS := "nats://nats:4222"
	es, _ := messaging.OpenConnection(addrNATS)
	messaging.SetEventStore(es)
	log.Println("NATS connected")

	woofListener, err := messaging.SubscribeWoofMessage()
	if err != nil {
		util.FailOnError(err, "Failed to subscribe to Woof messages")
	}
	for woof := range woofListener {
		fmt.Println(woof)
	}
}
