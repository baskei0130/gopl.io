package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/baskei0130/gopl.io/chap5/links"
)

// token: 20個の平行なリクエストという限界を
// 強制するために使われる計数セマフォ
var tokens = make(chan struct{}, 20)
var maxDepth = flag.Int("depth", 0, "Crawl Depth")

func crawl(depth int, url string) *node {
	if depth >= *maxDepth {
		return &node{depth + 1, nil}
	}
	tokens <- struct{}{}
	list, err := links.Extract(url)
	<-tokens // release token
	if err != nil {
		log.Print(err)
	}
	return &node{depth + 1, list}
}

type node struct {
	depth int
	links []string
}

func main() {
	flag.Parse()
	worklist := make(chan *node) // list of url, including dup
	var n int

	n++
	go func() { worklist <- &node{depth: 0, links: flag.Args()} }()

	seen := make(map[string]bool)
	for ; n > 0; n-- {
		node := <-worklist
		for _, link := range node.links {
			if !seen[link] {
				seen[link] = true
				fmt.Println(link)
				n++
				go func(depth int, link string) {
					worklist <- crawl(depth, link)
				}(node.depth, link)
			}
		}
	}
}
