package route

import (
	"fmt"
	"net/http"
)

func Get(url string, value string) {
	http.HandleFunc(url, func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(value)
		res.Write([]byte(value))
	})
}

func Post(url string, value string) {
	http.HandleFunc(url, func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(value)
		res.Write([]byte(value))
	})
}
