package main

import (
	"fmt"
	"log"
	"os"

	"github.com/baskei0130/gopl.io/chap5/links"
)

// breadthFirst: worklist 内の個々の項目に対して f を呼び出す.
// f から返されたすべての項目は worklist へ追加される.
// f はそれぞれの項目に対して高々1度しか呼び出されない.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
