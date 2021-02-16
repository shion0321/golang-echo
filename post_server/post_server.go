package main

import (
	"fmt"
	"log"
	"net/http"
	"local.packages/utils"
)

func postHandler(writer http.ResponseWriter, request *http.Request){
	placeholder := []byte("signature list goes here")
	_,err := writer.Write(placeholder)
	utils.Check(err)
}

func main()  {
	http.HandleFunc("/post",postHandler)
	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}