package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// soleTitle: doc 中の最初の空ではない title 要素のテキストと
// title 要素が1つだけ出なかったらエラーを返す
func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}

	defer func() {
		switch p := recover(); p {
		case nil:
			// no panic
		case bailout{}:
			// 予期された panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	// 2つ以上の空ではない title を見つけたら再帰から抜け出させる
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, err
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

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			log.Print(err)
		}

		doc, err := html.Parse(resp.Body)
		if err != nil {
			log.Print(err)
		}
		resp.Body.Close()

		title, err := soleTitle(doc)
		if err != nil {
			log.Print(err)
		}
		fmt.Println(title)
	}
}
