package main

import (
	"time"
)

type EntryDAO interface {
	getById(id string) (Entry, error)
	// getAll() []Entry..?, error
	addEntry(entry Entry) (string, error)
	// deleteEntry(is string) error
}

type FuelEntryDAO struct{}

func (f *FuelEntryDAO) getById(id string) (Entry, error) {
	time.Sleep(time.Duration(3*time.Second))
	return &Entry{Id: "GUID", Gallons: 12, Distance: 248}
}

func (f *FuelEntryDAO) addEntry(entry Entry) (string, error) {
	time.Sleep(time.Duration(3*time.Second))
	return "New GUID":
}