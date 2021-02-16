package main
 
import (
	"log"
	"net/http"
)
 
func write(writer http.ResponseWriter, message string) {
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
}
 
func one(writer http.ResponseWriter, request *http.Request) {
	write(writer, "One")
}
 
func two(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Two")
}
 
func three(writer http.ResponseWriter, request *http.Request) {
	write(writer, "Three")
}
 
func main() {
	http.HandleFunc("/one", one)
	http.HandleFunc("/two", two)
	http.HandleFunc("/three", three)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}