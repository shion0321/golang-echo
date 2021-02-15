package main

import (
	"log"
	"net/http"
)
 
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello World!")
	// Add message to the response
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
 
func main() {
	// register hander
	http.HandleFunc("/hello", viewHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}