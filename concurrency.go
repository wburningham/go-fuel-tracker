package main

import (
	"fmt"
	"log"
	"time"
)

func httpEvenNumToString(num int) (string, error) {
	log.Println("Before mock http request")
	// Simulates an HTTP request
	time.Sleep(time.Duration(num) * time.Second)
	if num%2 == 0 {
		return fmt.Sprintf("EVEN: %d", num), nil
	} else {
		return "", fmt.Errorf("ERROR ODD: %d", num)
	}
}

func evenToString() (string, error) {
	return httpEvenNumToString(2)
}

func oddToString() (string, error) {
	return httpEvenNumToString(3)
}

func serialPrinting() {
	evenStr, err := evenToString()
	if err != nil {
		log.Println(err)
	}
	log.Println(evenStr)

	oddStr, err := oddToString()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(oddStr)
	}

}

func concurrentPrinting() {
	var evenStr, oddStr string

	// important to use a buffered channel so that the routines are
	// guaranteed not to block.
	errorChan := make(chan error, 2)

	go func() {
		var err error
		evenStr, err = evenToString()
		errorChan <- err
	}()

	go func() {
		var err error
		oddStr, err = oddToString()
		errorChan <- err
	}()

	for evenStr == "" || oddStr == "" {
		select {
		case err := <-errorChan:
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	log.Println(evenStr)
	log.Println(oddStr)

}

func main() {
	// concurrentPrinting()
	serialPrinting()
}
