package main

import (
	"fmt"
	"time"
)

func httpEvenNumToString(num int) (string, error) {
	fmt.Println("Before mock http request")
	// Simulates an HTTP request
	time.Sleep(time.Duration(num) * time.Second)
	if num%2 == 0 {
		return fmt.Sprintf("EVEN: %d", num), nil
	} else {
		return "", fmt.Errorf("ODD: %d", num)
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
		fmt.Println(err)
	}
	fmt.Println(evenStr)

	oddStr, err := oddToString()
	if err != nil {
		fmt.Printf("ERROR %s\n", err)
	} else {
		fmt.Println(oddStr)
	}

}

func concurrentPrinting() {
	evenChn := make(chan string)

	go func() {
		serialPrinting()
	}()
}

func main() {
	serialPrinting()
}
