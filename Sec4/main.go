package main

import (
	"fmt"
	"net/http"
)

// Simple HTTP server

/*func main(){
	mux := http.NewServeMux() // Simple multiplexer that takes the requests and routes it to the API
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("request received")
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe("localhost:3000", mux) // API ready listening for requests in port :3000, server localhost
}*/

// HTTP methods

func main(){
	mux := http.NewServeMux()  cd 
	mux.HandleFunc("/getgoing/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("request received")
		fmt.Println(r.Method) // To see the HTTP method
		w.Write([]byte("Hello world"))
	})
	http.ListenAndServe("localhost:3000", mux) 
}