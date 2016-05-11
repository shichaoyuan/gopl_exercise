package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex3: %v\n", err)
		os.Exit(1)
	}

	for _, text := range visit(nil, doc) {
		fmt.Printf("%s\n", text)
	}
}

func visit(texts []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "script" {
		return texts
	}

	if n.Type == html.TextNode {
		texts = append(texts, n.Data)
	}
	if n.FirstChild != nil {
		texts = visit(texts, n.FirstChild)
	}
	if n.NextSibling != nil {
		texts = visit(texts, n.NextSibling)
	}
	return texts
}
