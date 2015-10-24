package main

import (
	"fmt"
	"time"
)

type FuelPriceSvc struct{}

func (f *FuelPriceSvc) get() (string, error) {
	fmt.Println("Before mock http request")
	// Simulates an HTTP request
	time.Sleep(time.Second)
	return fmt.Sprint("this is the price")
}
