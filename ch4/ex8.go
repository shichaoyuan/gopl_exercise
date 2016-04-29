package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	letterCounts := make(map[rune]int)
	digitCounts := make(map[rune]int)

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}
		if unicode.IsLetter(r) {
			letterCounts[r]++
		}
		if unicode.IsDigit(r) {
			digitCounts[r]++
		}
	}
	fmt.Printf("letter\tcount\n")
	for l, n := range letterCounts {
		fmt.Printf("%q\t%d\n", l, n)
	}
	fmt.Printf("digit\tcount\n")
	for d, n := range digitCounts {
		fmt.Printf("%q\t%d\n", d, n)
	}
}
