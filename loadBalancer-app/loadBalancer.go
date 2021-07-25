package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	counter int

	listenAddr = "localhost:8080"

	servers = []string{
		"localhost:5001",
		"localhost:5002",
		"localhost:5003",
	}
)

func main() {
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("dinleyici işlemi hatalı: %s", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("dinleyici bağlantıyı kabul edemedi: %s", err)
		}

		backend := chooseBackend()
		fmt.Printf("counter=%d backend=%s\n", counter, backend)
		go func(){
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("UYARI: proxy kapandı: %s", err)
			}
		}()
	}
}

func proxy(backend string, c net.Conn) error {
	bc, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("arka arayüz bağlanmayı kabul edemedi %s: %v", backend, err)
	}

	go io.Copy(bc,c)
	go io.Copy(c,bc)

	return nil
}

func chooseBackend() string {
	s := servers[counter%len(servers)]
	counter++
	return s
}