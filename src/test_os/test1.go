package main

import (
	"os"
	"log"
	"fmt"
	"unsafe"
	"syscall"
)

func main() {
	file, err := os.Create("golang.txt")
	if err !=nil{
		log.Fatal(err.Error())
	}
	defer file.Close()
	fmt.Println(unsafe.Pointer(file))
	file.Write([]byte("myself"))
	fmt.Printf("%b",syscall.InvalidHandle)
}



