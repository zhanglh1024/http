package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		helloMyMux(w, r)
		return
	}
	http.NotFound(w, r)
	return
}
func helloMyMux(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello myMux you are perfect!!")
}
func main() {
	mux := &MyMux{}
	http.ListenAndServe(":8089", mux)
}
