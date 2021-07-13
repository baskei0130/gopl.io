package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var hit []*html.Node

	forEachNode(doc, func(n *html.Node) {

		if n.Type == html.ElementNode {
			for _, nm := range name {
				if n.Data == nm {
					hit = append(hit, n)
					break
				}
			}
		}

	}, nil)

	return hit
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func visit(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println("img: ", ElementsByTagName(doc, "img"))
	fmt.Println("h: ", ElementsByTagName(doc, "h1", "h2", "h3", "h4"))
}

func main() {
	for _, url := range os.Args[1:] {
		visit(url)
	}
}
