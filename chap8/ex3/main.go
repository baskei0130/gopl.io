package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.DialTCP("tcp", &net.TCPAddr{}, &net.TCPAddr{Port: 8000})
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{} // メインゴルーチンへ通知
	}()
	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done // バックグラウンドのゴルーチンが完了するのを待つ
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
