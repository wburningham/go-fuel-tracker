package main

import (
	"time"
)

// TODO do I need a mutex here for an in memory DAO

// TODO should and why are interfaces capital? for `var entry Entry`?
type EntryDAO interface {
	GetById(id string) (Entry, error)
	// GetAll() []Entry..?, error
	AddEntry(entry Entry) (string, error)
	// DeleteEntry(is string) error
}

type InMemoryFuelEntryDAO struct{}

func (f *InMemoryFuelEntryDAO) GetById(id string) (Entry, error) {
	time.Sleep(time.Duration(1 * time.Second))
	// TODO why is the below not &Entry{...}
	return Entry{Id: "GUID", Gallons: 12, Distance: 248}, nil
}

func (f *InMemoryFuelEntryDAO) AddEntry(entry Entry) (string, error) {
	time.Sleep(time.Duration(1 * time.Second))
	return "New GUID", nil
}
