package main

import (
	"fmt"
	"os"
)

//os.NewFile()函数原型是func NewFile(fd uintptr, name string) *File 第一个传入的是句柄，
// 然后是文件名称，这个函数并不是真的创建了一个文件，是新建一个文件不保存，然后返回文件的指针
func main() {
	f, _ := os.Open("10.txt")
	defer f.Close()
	f1 := os.NewFile(f.Fd(), "ceshi.go") //输如ini.go的句柄
	defer f1.Close()
	fd, _ := f1.Stat()
	fmt.Println(fd.ModTime()) //返回的是ini.go的创建时间2013-11-27 09:11:50.2793737 +0800 CST

}
