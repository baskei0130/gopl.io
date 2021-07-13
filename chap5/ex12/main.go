package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// n から始まるツリー内の個々のノードxに対して
// 関数 pre(x), post(x) を呼び出す. その2つの関数はオプション
// pre は子ノードを訪れる前に呼び出され
// post は子ノードを訪れた後に呼び出される
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

/*
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
	}
}
*/

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("failed to get(%s): %v\n", url, err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Printf("failed to parse: %v\n", err)
	}

	var depth int

	startElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}

	endElement := func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

	forEachNode(doc, startElement, endElement)
}
