package main

import (
	"log"
	"runtime"
	"fmt"
)

func main() {
	test()
}

func test() {
	test2()
}

func test2() {
	b :=make([]uintptr,10)
	callers := runtime.Callers(0, b)
	fmt.Println(callers,b)
	pc, file, line, ok := runtime.Caller(2)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f := runtime.FuncForPC(pc)
	log.Println(f.Name())

	pc, file, line, ok = runtime.Caller(0)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f = runtime.FuncForPC(pc)
	log.Println(f.Name())

	pc, file, line, ok = runtime.Caller(1)
	log.Println(pc)
	log.Println(file)
	log.Println(line)
	log.Println(ok)
	f = runtime.FuncForPC(pc)
	log.Println(f.Name())
}
