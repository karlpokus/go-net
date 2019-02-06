package main

import (
	"net"
	"os"
	"time"
	"log"
	"io"
)

func main() {
	l, err := net.Listen("tcp", "localhost:37042")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("listening")
	for {
		conn, err := l.Accept() // blocking
		if err != nil {
			log.Fatal(err)
		}
		//go copyOut(conn)
		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close()
	remote, err := net.Dial("tcp", "localhost:37043") // returns a net.Conn
	if err != nil {
		log.Println(err)
		return
	}
	defer remote.Close()
	go io.Copy(remote, conn) // goroutine should return properly
	io.Copy(conn, remote)
}

func copyOut(conn net.Conn) {
	defer conn.Close()
	//n, err := io.Copy(os.Stdout, conn)
	for {
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println(err)
			return
		}
		os.Stdout.Write(buf[:n])
	}

}
