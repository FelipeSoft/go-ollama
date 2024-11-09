package main

import (
	"awesomeProject/rest"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := rest.InitRoutes()

	fmt.Println("HTTP server listening on port 1818")
	log.Fatal(http.ListenAndServe("localhost:1818", r))
}
