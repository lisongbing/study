package main

import (
	"os"
	"fmt"
)

func main() {
	r, w, _ := os.Pipe()

	w.Write([]byte("hello"))
	var buf []byte = make([]byte,1024)
	n, err := r.Read(buf)
	if err !=nil{
		fmt.Println(err.Error())
	}
	fmt.Println("de:",string(buf[:n]))
}
