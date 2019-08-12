package main

import (
	"fmt"
	"runtime"
)

func main() {

	defer rrr()

		var array [3]int
		for i := 0; i < 4; i++ {
			fmt.Println("dfasdf", array[i])
		}
	}

func rrr(){
		if err := recover(); err != nil {

			var stack string
			fmt.Println(fmt.Sprintf("Handler crashed with error", err))

			for i := 1; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				fmt.Println(fmt.Sprintf("%s:%d", file, line))

				stack = stack + fmt.Sprintln(fmt.Sprintf("%s:%d", file, line)) + "\r\n"
			}
			fmt.Println(stack)
		}

}