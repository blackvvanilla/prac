package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Started, Listening on 8080")
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
