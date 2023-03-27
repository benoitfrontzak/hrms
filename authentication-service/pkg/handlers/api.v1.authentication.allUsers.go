package handlers

import (
	"authentication/pkg/pg"
	"encoding/json"
	"log"
	"net/http"
)

// AllUsers returns all users from the DB
func (rep *Repository) AllUsers(w http.ResponseWriter, r *http.Request) {
	// Fetch all users from DB
	all, err := rep.App.Models.User.GetAll()
	if err != nil {
		log.Println("error while fetching all users", err)
	}

	type myJsonResp struct {
		Data []*pg.User
	}
	resp := myJsonResp{
		Data: all,
	}

	json.NewEncoder(w).Encode(&resp)

}
