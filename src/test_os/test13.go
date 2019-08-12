package main

//(f *File) Write像文件中写入内容，
// 函数原型func (f *File) Write(b []byte) (n int, err error)返回的是n写入的字节数
import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("1.go", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755) //以追加和读写的方式去打开文件
	n, _ := f.Write([]byte("helloword"))                     //我们写入hellword
	fmt.Println(n)                                           //打印写入的字节数
	b := make([]byte, 20)
	f.Seek(0, 0)                                            //指针返回到0
	data, _ := f.Read(b)
	fmt.Println(string(b[:data]))                           //输出了packagehelloword
}
