package main
//(f *File).Chdir()修改工作目录，函数原型func (f *File) Chdir() error，
// 这个时候f必须是目录了,但是吧这个不支持windows

import (
	"fmt"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
	f, _ := os.Open("10.txt")
	err := f.Chdir()
	if err != nil {
		fmt.Println(err)
	}
	dir1, _ := os.Getwd()
	fmt.Println(dir1)
}