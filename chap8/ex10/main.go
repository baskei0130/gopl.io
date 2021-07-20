package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/baskei0130/gopl.io/chap8/ex10/links"
)

// token: 20個の平行なリクエストという限界を
// 強制するために使われる計数セマフォ
var tokens = make(chan struct{}, 20)
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // get token
	list, err := links.Extract(url, done)
	<-tokens // release token
	if err != nil {
		log.Print(err)
	}
	return list
}

func crawlUnseenLinks(unseenLinks <-chan string, worklist chan<- []string, wg *sync.WaitGroup) {
	defer func() { wg.Done() }()
	for {
		select {
		case <-done:
			return
		case link := <-unseenLinks:
			foundLinks := crawl(link)
			wg.Add(1)
			go func() {
				wg.Done()
				select {
				case <-done:
					return
				default:
					worklist <- foundLinks
				}
			}()
		}
	}
}

func main() {
	worklist := make(chan []string)  // list of url, including dup
	unseenLinks := make(chan string) // 重複していない URL

	var wg sync.WaitGroup

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
		log.Println("canceled")

		wg.Wait()
		close(worklist)
		close(unseenLinks)
	}()

	go func() { worklist <- os.Args[1:] }()

	// 未探索のリンクを取得するために20個のクローラのゴルーチンを生成
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go crawlUnseenLinks(unseenLinks, worklist, &wg)
	}

	// メインゴルーチンは worklist の項目の重複をなくし
	// 未探索の項目をクローラへ送る
	seen := make(map[string]bool)
	for {
		select {
		case list, ok := <-worklist:
			if !ok {
				return
			}
			for _, link := range list {
				select {
				case <-done:
					continue
				default:
					if !seen[link] {
						seen[link] = true
						unseenLinks <- link
					}
				}
			}
		case <-done:
			for range worklist {
				// flush worklist
			}
			for range unseenLinks {
				// flush unseenLinks
			}
		}
	}
}
