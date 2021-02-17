package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"local.packages/utils"
)

func postHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("views/view.html")
	utils.Check(err)
	err = html.Execute(writer, nil)
	utils.Check(err)
}

func main() {
	http.HandleFunc("/post", postHandler)
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
