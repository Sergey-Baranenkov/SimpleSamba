package main

import (
	"fmt"
	"log"
	"net/http"
)

func main (){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "static.html")
	})

	fmt.Println("Server is listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal("Cannot start the server because of ", err)
	}
}
