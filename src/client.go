package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"flag"
	"time"
)

var (
	chanQuit = make(chan bool, 0)
)

func CHandler(err error, why string) {

	if err != nil {
		fmt.Println(why, err)
		os.Exit(1)
	}
}
func main() {
	var name string
	//todo:在命令行参数中携带昵称
	flag.StringVar(&name,"name","","昵称")
	flag.Parse()
	//拨号链接，获得connection
	conn, err := net.Dial("tcp", "112.74.45.146:8000")
	//fmt.Println("ssse")
	defer conn.Close()

	CHandler(err, "net.Dial")
	//在一条独立的携程中，接收输入，并发送消息
	go handlerSend(conn,name)
	//在一条独立的携程中接受服务端的消息
	go handleRecive(conn,name)
	//添加优雅退出逻辑
	<-chanQuit

}
func handlerSend(conn net.Conn,name string) {
	//fmt.Println("ssse")
	_, err := conn.Write([]byte(name))
	CHandler(err,"conn.Write")
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := reader.ReadLine()
		_, err := conn.Write(line)
		CHandler(err, "handlerSend")

		if string(line) =="exit"{
			os.Exit(0)
		}
	}

}
func handleRecive(conn net.Conn,name string) {
	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != io.EOF {
			CHandler(err, "handleRecive")
		}
		if n > 0 {
			msg := string(buffer[0:n])
			if msg != "" {
				fmt.Println("接收到服务器的信息:", msg)
				saveTalkToFile(msg,name)
			}
		}
	}
}

func saveTalkToFile(msg string,name string)  {
	talkRecond := fmt.Sprintln(time.Now().Format("2006-01-02 15:04:05"), msg)
	fmt.Println(talkRecond)
	file, err := os.OpenFile("E:/mygo/study/src/"+name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0x666)
	defer file.Close()
	if err !=nil{
		fmt.Println("file:",err)
		return
	}
	file.Write([]byte(talkRecond))
}