package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "helllo nice work!!")
	fmt.Fprintln(w, "how to build a web server!!")
}
func main() {
	http.HandleFunc("/", sayHelloName)       //配置路由
	err := http.ListenAndServe(":9090", nil) //监听端口
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
