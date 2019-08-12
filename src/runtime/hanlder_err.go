package main

import (
	"fmt"
	"runtime"
)

func main() {
		defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)       //这里可以打印到日志
			var stack string
			for i := 0; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				fmt.Println(fmt.Sprintf("%s:%d", file, line))  //这里可以打印到日志

				stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line)) + "\r\n"
			}
		}
	}()
	do := make(chan bool)
	go func() {
		for {
			fmt.Println("hello world")
			do <- true
		}
	}()
	a := []string{"red", "blue", "yellow"}
	go readSliceOneByOne(a)
	for i := 0; i < 10; i++ {
		<-do
	}
}

func readSliceOneByOne(a []string) {
/*	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) //这里可以打印到日志
			var stack string
			for i := 0; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				fmt.Println(fmt.Sprintf("%s:%d", file, line)) //这里可以打印到日志

				stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line)) + "\r\n"
			}
		}
	}()*/
	for _, v := range a {
		fmt.Println(v)
	}
	a[3] = "grey"
}
