package main

import (
	"fmt"
	"linkingsign/router"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
)

func main() {
	r := router.Router()

	corsRouter := handlers.CORS()(r)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Starting server on the port " + port + "...")

	log.Fatal(http.ListenAndServe(":"+port, corsRouter))
}
