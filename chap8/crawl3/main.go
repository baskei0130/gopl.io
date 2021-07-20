package main

import (
	"fmt"
	"log"
	"os"

	"github.com/baskei0130/gopl.io/chap5/links"
)

// token: 20個の平行なリクエストという限界を
// 強制するために使われる計数セマフォ
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // get token
	list, err := links.Extract(url)
	<-tokens // release token
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // list of url, including dup
	unseenLinks := make(chan string) // 重複していない URL

	go func() { worklist <- os.Args[1:] }()

	// 未探索のリンクを取得するために20個のクローラのゴルーチンを生成
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// メインゴルーチンは worklist の項目の重複をなくし
	// 未探索の項目をクローラへ送る
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}
