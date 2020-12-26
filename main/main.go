package main

import (
	"gowebapi/config"
	"log"
	"net/http"
)

func main() {
	routes := make(map[string]func(http.ResponseWriter, *http.Request))
	config.Routes(&routes)
	for k, v := range routes {
		http.HandleFunc(k, v)
	}
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
