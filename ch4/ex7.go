package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(bytes []byte) {
	var lr, rr rune
	var lw, rw int

	for l, r := 0, len(bytes); l < r; l, r = l+rw, r-lw {
		lr, lw = utf8.DecodeRune(bytes[l:])
		rr, rw = utf8.DecodeLastRune(bytes[:r])
		copy(bytes[l+rw:], bytes[l+lw:r-rw])
		copy(bytes[l:], []byte(string(rr)))
		copy(bytes[r-lw:], []byte(string(lr)))
	}
}

func main() {
	s := []byte("hello, 世界!")
	fmt.Printf("%q\n", s)
	reverse(s)
	fmt.Printf("%q\n", s)
}
