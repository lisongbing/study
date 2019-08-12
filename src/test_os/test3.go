package main

import (
	"os"
	"fmt"
)
//file.Fd()  返回句柄
func main() {
	file, _ := os.Open("10.txt")
	fmt.Println(file.Fd())
}
