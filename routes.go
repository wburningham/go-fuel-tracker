package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO my handler names are lowerCase to be private,
// but all interfaces and structs are UpperCase
func getEntry(res http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(res, "id missing in URL path", http.StatusBadRequest)
		return
	}

	entry, err := entryDb.GetById(id)
	if err != nil {
		http.Error(res, "id missing in URL path", http.StatusInternalServerError)
		return
	}

	// TODO write to JSON
	fmt.Fprintf(res, "Entry for: %s", entry.ToString())
}

func createEntry(res http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)

	var entry Entry
	err := decoder.Decode(&entry)

	if err != nil {
		http.Error(res, "Invalid POST body", http.StatusBadRequest)
		return
	}
	// TODO use a Db

	// TODO write to JSON
	fmt.Fprintf(res, "New entry ID: NEW GUID")
}
