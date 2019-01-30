package main

import (
	    "fmt"
		"net/http"
		"log"
	)


func main(){
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"%q\n",r.URL.Path)
}