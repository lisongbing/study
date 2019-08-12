package main

//(f *File).Readdirnames这个函数的作用是读取目录内的文件名，其实上一个函数我们已经实现了这个函数的功能，
// 函数的原型func (f *File) Readdirnames(n int) (names []string, err error)，
// 跟上边一下只不过返回的是文件名 []string的slice
import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("E:/study/server")
	names, err := f.Readdirnames(0)
	if err != nil {
		fmt.Println(err)
	}
	for i, name := range names {
		fmt.Printf("filename %d: %s\n", i, name)
	}
}