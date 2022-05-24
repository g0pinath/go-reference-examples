package main

import "net/http" 
import "fmt"


func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("received..")
		w.Write([]byte("hello..."))
	})
    http.ListenAndServe("localhost:3000", mux)
}