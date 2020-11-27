package main

import (
	"log"
	"net/http"

	"github.com/allenjoseph/go-cqrs/db"
	"github.com/allenjoseph/go-cqrs/messaging"
	"github.com/allenjoseph/go-cqrs/util"
	"github.com/gorilla/mux"
)

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/woofs", woofsHandler).Methods("POST")
	return
}

func main() {
	defer db.Close()
	defer messaging.Close()

	// Connect to Postgres
	addrDB := "postgres://woofer:woofwoof@postgres/woofer?sslmode=disable"
	r, _ := db.OpenConnection(addrDB)
	db.SetRepository(r)
	log.Println("Postgres connected")

	// Connect to Nats
	addrNATS := "nats://nats:4222"
	es, _ := messaging.OpenConnection(addrNATS)
	messaging.SetEventStore(es)
	log.Println("NATS connected")

	// Start Woof service
	log.Println("Woof service started")
	router := newRouter()
	err := http.ListenAndServe(":8080", router)
	util.FailOnError(err, "Failed to serve woof service")
}
