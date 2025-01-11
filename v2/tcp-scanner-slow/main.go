package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// порт закрыт или отфильтрован
			// fmt.Printf("%d close\n", i)
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
