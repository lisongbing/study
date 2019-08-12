package main

//(f *File).Seek()这个函数大家一看就懂了，就是偏移指针的地址，函数的原型是func (f *File) Seek(offset int64, whence int) (ret int64, err error) 其中offset是文件指针的位置 whence为0时代表相对文件开始的位置，
// 1代表相对当前位置，2代表相对文件结尾的位置 ret返回的是现在指针的位置

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 10)
	f, _ := os.Open("10.txt")
	defer f.Close()
	f.Seek(1, 2)                //相当于开始位置偏移1
	n, _ := f.Read(b)
	fmt.Println(string(b[:n]))  //原字符package 输出ackage
}
