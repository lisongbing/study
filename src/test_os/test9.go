package main
//(f *File).ReadAt()这个函数的原型是
// func (f *File) ReadAt(b []byte, off int64) (n int, err error)加入了下标，
// 可以自定义读取多少
import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("10.txt")
	b := make([]byte, 20)
	n, _ := f.ReadAt(b, 1)
	fmt.Println(n)
	fmt.Println(string(b[:n]))
}
