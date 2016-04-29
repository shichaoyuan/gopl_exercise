package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCounts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		wordCounts[word]++
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
		os.Exit(1)
	}
	total := 0
	for _, v := range wordCounts {
		total += v
	}
	fmt.Printf("word\tfreq\n")
	for k, v := range wordCounts {
		fmt.Printf("%q\t%.2f\n", k, float64(v)/float64(total))
	}
}
