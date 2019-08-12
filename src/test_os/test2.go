package main

import (
	"os"
	"log"
)

func main() {
	file, err := os.OpenFile("10.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err !=nil{
		log.Fatalln(err.Error())
	}
	defer  file.Close()
	log.Println(file.Stat())
}
