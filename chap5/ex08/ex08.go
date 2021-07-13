package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var (
	url = flag.String("url", "", "url")
	id  = flag.String("id", "", "id")
)

func main() {
	flag.Parse()

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	n := ElementByID(doc, *id)

	if n != nil {
		var attr string
		for _, a := range n.Attr {
			attr += fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
		fmt.Printf("<%s%s>\n", n.Data, attr)
	}
}

// n から始まるツリー内の個々のノードxに対して
// 関数 pre(x), post(x) を呼び出す. その2つの関数はオプション
// pre は子ノードを訪れる前に呼び出され
// post は子ノードを訪れた後に呼び出される
func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if cn := forEachNode(c, pre, post); cn != nil {
			return cn
		}
	}

	if post != nil {
		if !post(n) {
			return n
		}
	}
	return nil
}

var depth int

// 指定された id 属性を持つ最初の HTML 要素を探索
func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					return false
				}
			}
		}
		return true
	}, nil)
}
