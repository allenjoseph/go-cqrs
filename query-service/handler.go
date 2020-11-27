package main

import (
	"net/http"
	"strconv"

	"github.com/allenjoseph/go-cqrs/db"
	"github.com/allenjoseph/go-cqrs/util"
)

func listWoofsHandler(w http.ResponseWriter, r *http.Request) {
	offset := uint64(0)
	limit := uint64(10)
	var err error

	offsetStr := r.FormValue("offset")
	if len(offsetStr) != 0 {
		offset, err = strconv.ParseUint(offsetStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid offset parameter")
			return
		}
	}

	limitStr := r.FormValue("limit")
	if len(limitStr) != 0 {
		limit, err = strconv.ParseUint(limitStr, 10, 64)
		if err != nil {
			util.ResponseError(w, http.StatusBadRequest, "Invalid limit parameter")
			return
		}
	}

	ctx := r.Context()
	woofs, err := db.ListWoofs(ctx, offset, limit)
	if err != nil {
		util.ResponseError(w, http.StatusInternalServerError, "Could not fetch woofs")
		return
	}

	util.ResponseOk(w, woofs)
}
