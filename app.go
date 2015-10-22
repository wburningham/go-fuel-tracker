package main

import (
	structs "./structs"
	"fmt"
	"log"
)

func main() {
	fmt.Println("hello world")
	e := &structs.AppError{Message: "foo message", Code: 500}
	log.Printf("Handler error: status code: %d, message: %s",
		e.Code, e.Message)

}
