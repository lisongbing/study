package main

import (
	"fmt"
	"os"
	"diabloserver/res/def"
)

func main() {
/*	err := def.LoadAll()
	fmt.Println(err)*/
	file, e := os.OpenFile("E:/workproject/src/diabloserver/res/loadTips.xlsx",os.O_RDONLY,0666)
	defer file.Close()
	if e !=nil{
		fmt.Println(e)
		return
	}
	fmt.Println(file)
	data := def.LoadSignInData("")
	fmt.Println(data)
}

