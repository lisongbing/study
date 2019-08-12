package main
//(f *File).Readdir()函数原型func (f *File) Readdir(n int) (fi []FileInfo, err error)，我们要打开一个文件夹，
// 然后设置读取文件夹文件的个数，返回的是文件的fileinfo信息
import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("E:/study/server")    //打开一个目录
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	ff, _ := f.Readdir(0)    //设置读取的数量 <=0是读取所有的文件 返回的[]fileinfo
	for i, fi := range ff {
		fmt.Printf("filename %d: %+v\n", i, fi.Name())  //我们输出文件的名称
	}
}
