package main

import (
	"log"
	"net/http"

	"github.com/allenjoseph/go-cqrs/db"
	"github.com/allenjoseph/go-cqrs/util"
	"github.com/gorilla/mux"
)

func newRouter() (router *mux.Router) {
	router = mux.NewRouter()
	router.HandleFunc("/woofs", listWoofsHandler).Methods("GET")
	return
}

func main() {
	defer db.Close()

	// Connect to Postgres
	addrDB := "postgres://woofer:woofwoof@postgres/woofer?sslmode=disable"
	r, _ := db.OpenConnection(addrDB)
	db.SetRepository(r)
	log.Println("Postgres connected")

	// Start Query service
	log.Println("Query service started")
	router := newRouter()
	err := http.ListenAndServe(":9090", router)
	util.FailOnError(err, "Failed to serve query service")
}
