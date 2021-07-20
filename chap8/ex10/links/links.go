package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func cancelled(done <-chan struct{}) bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// Extract: 指定された URL へ HTTP GET リクエストをおこない,
// レスポンスを HTML として Parse して, その HTML ドキュメント内のリンクを返す
func Extract(url string, done <-chan struct{}) ([]string, error) {
	if cancelled(done) {
		return nil, nil
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Cancel = done

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %s", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

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
