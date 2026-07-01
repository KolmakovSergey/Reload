package main

import (
	"fmt"
	"net/http"
)

func eventHandler(rw http.ResponseWriter, r *http.Request){
	
}

func main() {

	http.HandleFunc("/", eventHandler)

	if err := http.ListenAndServe("8080", _); err != nil{
		fmt.Println(err.Error())
	}
}
