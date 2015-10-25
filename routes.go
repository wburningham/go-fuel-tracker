package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// var (
// 	entryDAO FuelEntryDAO
// )

func getEntry(res http.ResponseWriter, req *http.Request) {
	id, err := mux.Vars(req)["id"]
	if !err {
		http.Error(res, "id missing in URL path", http.StatusBadRequest)
		return
	}
	// TODO use a DAO
	fmt.Fprintf(res, "Entry for: %s", id)
}

func createEntry(res http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var entry Entry
	err := decoder.Decode(&entry)

	if err != nil {
		http.Error(res, "Invalid POST body", http.StatusBadRequest)
		return
	}
	// TODO use a DAO

	fmt.Fprintf(res, "New entry ID: NEW GUID")
}
