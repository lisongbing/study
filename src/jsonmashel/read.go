package main

import (
	"encoding/json"
	"os"
	"fmt"
)

func main() {
	file, err := os.Open("test.json")
	if err !=nil{
		fmt.Println(err)
		os.Exit(-1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.Decode()
}