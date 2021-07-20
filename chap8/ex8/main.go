// 時刻を定期的に書き出すTCP Server
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) // 1度に1つの接続を処理
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	called := make(chan struct{})
	go func(c net.Conn, called <-chan struct{}) {
		for {
			aborted := time.After(10 * time.Second)
			select {
			case _, ok := <-called:
				if !ok {
					return
				}
			case <-aborted:
				c.Close()
				fmt.Println("connection closed")
			}
		}
	}(c, called)

	for input.Scan() {
		called <- struct{}{}
		go echo(c, input.Text(), 1*time.Second)
	}
	// ignore the potential error from input.Err()
	close(called)
	c.Close()
}
