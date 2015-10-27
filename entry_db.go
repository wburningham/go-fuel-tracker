package main

import (
	"fmt"
	"sync"
	// "time"

	"github.com/satori/go.uuid"
)

// Ensure InMemoryFuelEntryDb conforms to the EntryDb interface.
var _ EntryDb = &InMemoryFuelEntryDb{}

type InMemoryFuelEntryDb struct {
	mu      sync.Mutex
	entries map[string]*Entry // Where we store the entries

}

func newInMemoryFuelEntryDb() *InMemoryFuelEntryDb {
	return &InMemoryFuelEntryDb{
		entries: make(map[string]*Entry),
	}
}

func (db *InMemoryFuelEntryDb) GetById(id string) (*Entry, error) {
	if id == "" {
		return nil, fmt.Errorf("Entry with unassigned ID passed into GetById")
	}

	db.mu.Lock()
	defer db.mu.Unlock()
	// time.Sleep(time.Duration(1 * time.Second))

	entry, ok := db.entries[id]
	if !ok {
		return nil, fmt.Errorf("Entry not found with ID %s", id)
	}

	return entry, nil
}

func (db *InMemoryFuelEntryDb) AddEntry(entry *Entry) (string, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	// time.Sleep(time.Duration(1 * time.Second))

	entry.Id = uuid.NewV4().String()
	db.entries[entry.Id] = entry

	return entry.Id, nil
}

func (db *InMemoryFuelEntryDb) DeleteEntry(id string) error {
	if id == "" {
		return fmt.Errorf("Entry with unassigned ID passed into DeleteEntry")
	}

	db.mu.Lock()
	defer db.mu.Unlock()
	// time.Sleep(time.Duration(1 * time.Second))

	if _, ok := db.entries[id]; !ok {
		return fmt.Errorf("Could not delete entry with ID %s, does not exist.", id)
	}
	delete(db.entries, id)
	return nil
}
