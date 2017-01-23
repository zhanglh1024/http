package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloDemo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello My ")
	fmt.Fprintln(w, "hello My Demo")
	fmt.Fprintln(w, "the world is better!!")
}
func main() {
	http.HandleFunc("/", helloDemo)          //set the routes to be accessed
	err := http.ListenAndServe(":9080", nil) //set the listening port
	if err != nil {
		log.Fatal("ListenAndServer", nil)
	}
}
