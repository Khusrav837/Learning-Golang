package main

import (
	"fmt"
	"net/http"
	"sync"
	"log"
)

var count int
var mu sync.Mutex

func main(){
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	log.Fatal(http.ListenAndServe("localhost:8000",nil))
}

func handler(w http.ResponseWriter, r *http.Request){

	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"%q\n",r.URL.Path)

}


func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	fmt.Fprintf(w,"Count: %d\n",count)
	mu.Unlock()

}