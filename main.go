package main

import (
	"fmt"
	"os"
	"log"
)

func main() {
	fmt.Println("hello world")


	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	fmt.Println("Port:", portString)
}