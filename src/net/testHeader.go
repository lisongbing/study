package main

import (
	"net/http"
	"fmt"
)

type Handler struct {

}
func(h *Handler)ServeHTTP(w http.ResponseWriter, r *http.Request){

}

func SayHello(w http.ResponseWriter,r *http.Request){
	w.Header().Add("content-type","text/html")
	w.Header().Add("expires","26 Jul 2019")//expires:31 Dec 2008
	w.Header().Add("charset","iso-8859-1`")
	w.Write([]byte("hello 你好"))
}

func main() {
	http.HandleFunc("/hello",SayHello)
	err := http.ListenAndServe("127.0.0.1:8080",nil)
	if err !=nil{
		fmt.Println(err.Error())
	}
}
