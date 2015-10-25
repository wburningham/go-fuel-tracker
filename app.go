package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// TODO should and why are structs capital? for `var entry Entry`?
type Entry struct {
	// TODO should and why to make these Capital?
	Id       string `json:"id"`
	Gallons  int    `json:"gallons"`
	Distance int    `json:"distance"`
}

func (e *Entry) ToString() string {
	return fmt.Sprintf("Id: %s, Gallons: %d, Distance: %d", e.Id, e.Gallons, e.Distance)

}

var (
	entryDb EntryDb
)

func init() {
	entryDb = newInMemoryFuelEntryDb()
}

// TODO how to use a logger
// TODO how to use an appError

func main() {

	router := mux.NewRouter()

	// Catch all handler
	router.NotFoundHandler = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(fmt.Sprintf("%s %s not found\n", req.Method, req.URL)))
	})

	apiRouter := router.
		Headers("Content-Type", "application/json").
		PathPrefix("/api").
		Subrouter()
	// apiRouter.HandleFunc("/entries", getEntries).Methods("GET")
	apiRouter.HandleFunc("/entries/{id}", getEntry).Methods("GET")
	apiRouter.HandleFunc("/entries", createEntry).Methods("POST")
	// apiRouter.HandleFunc("/entries/{id}", deleteEntry).Methods("DELETE")

	// TODO uiRouter

	// TODO how to read in a port?
	log.Printf("serving on port 8080")
	http.ListenAndServe(":8080", router)
}
