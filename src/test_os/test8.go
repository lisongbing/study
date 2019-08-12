package main

//(f File).Read()这个是函数的指针来操作的,
// 属于FIlE的method,函数原型func (f *File) Read(b []byte) (n int, err error)输入读取的字节数，
// 返回字节的长度和error信息
import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 100) //设置读取的字节数
	f, _ := os.Open("10.txt")
	n, _ := f.Read(b)
	fmt.Println(n)
	fmt.Println(string(b[:n])) //输出内容 为什么是n而不直接输入100呢？底层这样实现的
	/*
		n, e := f.read(b)
		if n < 0 {
			n = 0
	}
		if n == 0 && len(b) > 0 && e == nil {
			return 0, io.EOF
		}
	*/
	//所以字节不足100就读取n
}
