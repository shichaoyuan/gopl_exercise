package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "pop elements: %v\n", err)
		os.Exit(1)
	}

	counter := make(map[string]int)
	visit(counter, doc)

	for k, v := range counter {
		fmt.Printf("%s\t%d\n", k, v)
	}
}

func visit(counter map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		counter[n.Data]++
	}
	if n.FirstChild != nil {
		visit(counter, n.FirstChild)
	}
	if n.NextSibling != nil {
		visit(counter, n.NextSibling)
	}
}
