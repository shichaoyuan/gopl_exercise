package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var depth int

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil
	}

	forEachNode(doc, startElement, endElement)

	return nil
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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		str := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			str = str + fmt.Sprintf(" %s=%q", a.Key, a.Val)
		}
		if n.FirstChild != nil {
			str = str + ">"
		} else {
			str = str + "/>"
		}
		fmt.Println(str)
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
	} else if n.Type == html.TextNode {
		if len(strings.TrimSpace(n.Data)) > 0 {
			fmt.Printf("%*s%s\n", depth*2, "", n.Data)
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
