package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpace(bytes []byte) []byte {
	out := bytes[:0]
	f := false
	for i := 0; i < len(bytes); {
		r, size := utf8.DecodeRune(bytes[i:])
		if unicode.IsSpace(r) {
			if !f {
				out = append(out, ' ')
			}
			f = true
		} else {
			for _, b := range bytes[i : i+size] {
				out = append(out, b)
			}
			f = false
		}
		i += size
	}

	return out
}

func main() {
	s := []byte("  \n hello   \t    世界 !")
	fmt.Printf("%q\n", s)
	s = squashSpace(s)
	fmt.Printf("%q\n", s)
}
