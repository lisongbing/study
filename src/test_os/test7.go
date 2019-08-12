package main

import (
	"fmt"
	"os"
)

//(f *File).Name()这个函数是返回文件的名称，
// 函数原型func (f *File) Name() string要文件的指针操作，
// 返回字符串，感觉比较鸡助的方法底层实现
func main() {
	f, _ := os.Open("10.txt")
	fmt.Println(f.Name()) //输出1.go
}
