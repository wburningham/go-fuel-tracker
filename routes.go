package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getEntry(res http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(res, "id missing in URL path", http.StatusBadRequest)
		return
	}

	entry, err := entryDb.GetById(id)
	if err != nil {
		http.Error(res, "Unable to lookup entry", http.StatusNotFound)
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

	guid, err := entryDb.AddEntry(&entry)
	if err != nil {
		http.Error(res, "Unable to lookup entry", http.StatusInternalServerError)
		return
	}

	// TODO write to JSON
	fmt.Fprintf(res, "New entry Id: %s", guid)
}

func deleteEntry(res http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(res, "id missing in URL path", http.StatusBadRequest)
		return
	}

	err := entryDb.DeleteEntry(id)
	if err != nil {
		http.Error(res, "unable to delete entry", http.StatusInternalServerError)
		return
	}

	// TODO write to JSON
	fmt.Fprintf(res, "Deleted: %s", id)

}
