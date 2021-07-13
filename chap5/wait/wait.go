package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	for _, url := range os.Args[1:] {
		if err := WaitFotServer(url); err != nil {
			fmt.Fprintf(os.Stderr, "Site is down: %v", err)
			os.Exit(1)
		}
	}
}

// URL のサーバへ接続を試みる
// 指数バックオフを使って1分間試みる
// 全ての試みが失敗したらエラーを報告
func WaitFotServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		log.Printf("server not responding (%s)", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
