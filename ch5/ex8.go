package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	url := os.Args[1]
	id := os.Args[2]

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	node := ElementByID(doc, id)
	fmt.Printf("%#v", node)

}

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, nil)
}

func forEachNode(n *html.Node, id string, pre, post func(*html.Node, string) bool) *html.Node {
	if pre != nil {
		if pre(n, id) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node := forEachNode(c, id, pre, post)
		if node != nil {
			return node
		}
	}

	if post != nil {
		if post(n, id) {
			return n
		}
	}

	return nil
}

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}
	return false
}
