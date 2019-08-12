package main

//(f *File) WriteAt()在偏移位置多少的地方写入，
// 函数原型是func (f *File) WriteAt(b []byte, off int64) (n int, err error)返回值是一样的
import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.OpenFile("1.go", os.O_RDWR, os.ModePerm)
	f.WriteAt([]byte("widuu"), 10) //在偏移10的地方写入
	b := make([]byte, 20)
	d, _ := f.ReadAt(b, 10)    //偏移10的地方开始读取
	fmt.Println(string(b[:d])) //widuudhellowordhello
}
