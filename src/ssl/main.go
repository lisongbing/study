package main

import (
	"crypto/tls"
	"log"
	"time"
	"crypto/rand"
	"fmt"
	"net"
)

func main(){
	crt, err := tls.LoadX509KeyPair("E:/study/src/ssl/server.csr", "E:/study/src/ssl/server.key")
	if err !=nil{
		log.Fatalln(err.Error())
	}
	tlsConfig :=&tls.Config{}
	tlsConfig.Certificates = []tls.Certificate{crt}

	tlsConfig.Time =time.Now

	tlsConfig.Rand = rand.Reader

	listener, err := tls.Listen("tcp", ":8888", tlsConfig)
	if err !=nil{
		log.Fatalln(err.Error())
	}
	
	for{
		conn, err := listener.Accept()
		if err!=nil{
			fmt.Println(err.Error())
			continue
		}else{
			go HandleClientConnect(conn)
		}
	}

}

func HandleClientConnect(conn net.Conn)  {
	defer conn.Close()
	fmt.Println("Receive Connect Request From",conn.RemoteAddr().String())
	buffer :=make([]byte,1024)
	for{
		len,err :=conn.Read(buffer)
		if err !=nil{
			log.Println(err.Error())
			break
		}
		fmt.Printf("Receive Data:%s\n",string(buffer[:len]))

		//发送 给客户端

		_,err =conn.Write([]byte("服务器收到数据:"+string(buffer[:len])))
		if err !=nil{
			break
		}
	}
	fmt.Println("Client"+conn.RemoteAddr().String()+"Connection closed")
}