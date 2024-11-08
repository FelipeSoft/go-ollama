package main

import (
	"awesomeProject/rest"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := rest.InitRoutes()
	fmt.Print("HTTP server listening on port 1818")
	log.Fatal(http.ListenAndServe(":1818", r))
}
