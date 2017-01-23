package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func helloLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello my good login") //这个写入到w的时输出到客户端的
	r.ParseForm()                         //解析url传递的参数，对于POST则解析响应包的主体(request body)
	//注意：如果没有调用Parseorm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息时输出到服务端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm() //调用Parseorm解析form表单中的信息
		fmt.Println(r.Form)
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func main() {
	http.HandleFunc("/", helloLogin) //设置访问路由
	http.HandleFunc("/login", login) //设置访问路由
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
