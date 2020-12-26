package controllers

import (
	"log"
	"net/http"
)

func Test(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello world!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
