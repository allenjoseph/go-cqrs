package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/allenjoseph/go-cqrs/db"
	"github.com/allenjoseph/go-cqrs/messaging"
	"github.com/allenjoseph/go-cqrs/model"
	"github.com/allenjoseph/go-cqrs/util"
	uuid "github.com/satori/go.uuid"
)

func woofsHandler(w http.ResponseWriter, r *http.Request) {
	req := model.WoofRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	}
	defer r.Body.Close()

	id, err := uuid.NewV4()
	if err != nil {
		util.ResponseError(w, http.StatusBadRequest, "Failed to generate woof ID")
		return
	}

	woof := model.Woof{
		ID:        id.String(),
		Body:      req.Message,
		CreatedAt: time.Now().UTC(),
	}

	// Create woof
	err = db.InsertWoof(r.Context(), woof)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Failed to create woof")
		return
	}

	// Publish woof
	err = messaging.PublishWoofMessage(woof)
	if err != nil {
		log.Println(err)
	}

	util.ResponseOk(w, model.WoofResponse{
		ID: woof.ID,
	})
}
