package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func sender(conn *net.TCPConn, name string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(conn, "%s: %s", name, str)
	}
}

func reader(conn *net.TCPConn) {
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			println("goodnight")
			os.Exit(0)
			return
		}
		fmt.Print(str)
	}
}

func main() {
	println("Escreva o endereço IP do servidor")
	read := bufio.NewReader(os.Stdin)

	ip, err := read.ReadString('\n')
	ip = ip[:len(ip)-1]

	println("Escreva o seu nome")
	name, err := read.ReadString('\n')
	name = name[:len(name)-1]

	tcpaddr, err := net.ResolveTCPAddr("tcp", ip+":2000")
	conn, err := net.DialTCP("tcp", nil, tcpaddr)

	if err != nil {
		println("deu ruim")
		os.Exit(0)
	}

	fmt.Fprintf(conn, "%s acabou de entrar\n", name)

	go reader(conn)
	sender(conn, name)

	conn.Close()

}
