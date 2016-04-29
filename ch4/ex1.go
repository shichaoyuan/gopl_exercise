package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]uint8

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + uint8(i&1)
	}
}

func NumOfDiffBits(a, b [32]uint8) int {
	var c [32]uint8
	n := 0
	for i := range a {
		c[i] = a[i] ^ b[i]
		n += int(pc[c[i]])
	}
	return n
}

func main() {
	c1 := sha256.Sum256([]byte("b"))
	c2 := sha256.Sum256([]byte("a"))
	fmt.Printf("%x\n%x\n", c1, c2)

	n := NumOfDiffBits(c1, c2)
	fmt.Printf("Number of bits: %d\n", n)
}
