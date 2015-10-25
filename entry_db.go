package main

import (
	"time"
)

// TODO do I need a mutex here for an in memory Db

// TODO should and why are interfaces capital? for `var entry Entry`?
type EntryDb interface {
	GetById(id string) (Entry, error)
	// GetAll() []Entry..?, error
	AddEntry(entry Entry) (string, error)
	// DeleteEntry(is string) error
}

// Ensure InMemoryFuelEntryDb conforms to the EntryDb interface.
var _ EntryDb = &InMemoryFuelEntryDb{}

type InMemoryFuelEntryDb struct {
	entries map[int64]*Entry // Where we stor the entries

}

func newInMemoryFuelEntryDb() *InMemoryFuelEntryDb {
	return &InMemoryFuelEntryDb{
		entries: make(map[int64]*Entry),
		// internalId: 1,
	}
}

func (f *InMemoryFuelEntryDb) GetById(id string) (Entry, error) {
	time.Sleep(time.Duration(1 * time.Second))
	// TODO why is the below not &Entry{...}
	return Entry{Id: "GUID", Gallons: 12, Distance: 248}, nil
}

func (f *InMemoryFuelEntryDb) AddEntry(entry Entry) (string, error) {
	time.Sleep(time.Duration(1 * time.Second))
	return "New GUID", nil
}
