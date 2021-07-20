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
	worklist := make(chan []string)
	var n int // worklist への送信待ちの数

	// コマンドラインの引数で開始
	n++
	go func() { worklist <- os.Args[1:] }()

	// crawl web by parallel
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}
