package main

import (
	"fmt"
	"log"

	// "net/http"
	"os"
)

func main() {
	fmt.Println("Hello world")
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not found in the enviroment.")
		// log.Fatal(http.ListenAndServe(":8000", nil))
	} else {
		fmt.Println("Port:", portString)
	}
}
