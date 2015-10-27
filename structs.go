package main

import (
	"fmt"
)

type Entry struct {
	Id       string `json:"id"`
	Gallons  int    `json:"gallons"`
	Distance int    `json:"distance"`
}

type EntryDb interface {
	GetById(id string) (*Entry, error)
	// GetAll() []entry..?, error
	AddEntry(entry *Entry) (string, error)
	DeleteEntry(is string) error
}

func (e *Entry) ToString() string {
	return fmt.Sprintf("Id: %s, Gallons: %d, Distance: %d", e.Id, e.Gallons, e.Distance)
}
