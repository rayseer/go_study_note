package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	//Set a Http header
	res.Header().Set("Content-Type", "application/json")

	//String To Byte use this function
	//byteArray := []byte(string)

	//Byte To String
	//str := string(byte)
	res.Write([]byte("This is an example server.\n"))
}

func testGetParams(res http.ResponseWriter, req *http.Request) {
	/*
		futurelabdeMac-mini:base futurelab$ curl localhost:3000/get_params?id=name
		/get_params?id=name
	*/
	if req.Method != "GET" {
		panic("Request method error")
	}
	url := req.URL.String()
	res.Write([]byte(url))
	res.Write([]byte("\n"))

	//return http.URL.value type
	q := req.URL.Query()
	//return string type
	id := q.Get("id")
	res.Write([]byte("ID：" + id))
	res.Write([]byte("\n"))
}

func testPostParams(res http.ResponseWriter, req *http.Request) {
	/*
		http.URL.PostForm 与 http.URL.Form  区别
	*/
	//return http.URL.values
	if req.Method != "POST" {
		panic("Request method error")
	}
	if err := req.ParseForm(); err != nil {
		log.Fatal(err)
	}
	//curl -d "name=smith" "localhost:3000/post_params"
	//return map[name:[smith]] type
	q := req.Form
	fmt.Println(q)
	name := q.Get("name")
	res.Write([]byte("Name：" + name + "\n"))
}

func testPutParams(res http.ResponseWriter, req *http.Request) {
	//curl -X PUT -d "name=smith" "localhost:3000/put_params"
	if req.Method != "PUT" {
		panic("Request method error")
	}
	if err := req.ParseForm(); err != nil {
		log.Fatal(err)
	}
	q := req.Form
	fmt.Println(q)
	name := q.Get("name")
	res.Write([]byte("Name：" + name + "\n"))
	res.Write([]byte("This is put params.\n"))
}

func testDeleteParams(res http.ResponseWriter, req *http.Request) {
	//curl -X DELETE "localhost:3000/delete_params"

	if req.Method != "DELETE" {
		panic("Request method error")
	}

	res.Write([]byte("This is delete params.\n"))
}

func main() {
	fmt.Println("Start port 3000...")

	//TODO: Use Restful API to send GET/POST/PUT/DELETE Method
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/get_params", testGetParams)
	http.HandleFunc("/post_params", testPostParams)
	http.HandleFunc("/put_params", testPutParams)
	http.HandleFunc("/delete_params", testDeleteParams)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
