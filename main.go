package main

import (
	"fmt"
	"linkingsign/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	port := "8080"
	fmt.Println("Starting server on the port " + port + "...")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
