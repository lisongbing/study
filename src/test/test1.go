package main

import "fmt"

func main() {
	type mm struct{
		name string
	}

	data :=make([]*mm,3)
	for i :=range data{
		fmt.Println(data[i]==nil)
	}
}
