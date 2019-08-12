package main

import (
	"fmt"
	"net"
)

func main() {
	/*	name := "192.168.1.97"

		ip := net.ParseIP(name)

		if ip == nil {
			fmt.Fprintf(os.Stderr, "Err:无效的地址")
			return
		}

		fmt.Fprintf(os.Stdout, "IP: %s %s\n", ip, ip.String())
		defaultMask := ip.DefaultMask()
		fmt.Fprintf(os.Stdout, "DefaultMask: %s %s\n", defaultMask, defaultMask.String())

		ones, bits := defaultMask.Size()
		fmt.Fprintf(os.Stdout, "ones: %d bits: %d\n", ones, bits)*/
	//net.Dial("tcp","golang.org:80")

	port, _ := net.LookupPort("tcp", "ftp")
	fmt.Println(port)
	ips, _ := net.LookupTXT("www.baidu.com")
	fmt.Println(ips)
}
