package main

import (
	"fmt"
	"sync"
)

var initPopCountOnce sync.Once

var pc [256]byte

func initPopCount() {
	fmt.Println("initPopCount")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	initPopCountOnce.Do(initPopCount)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	n := PopCount(128)
	fmt.Println(n)
	n = PopCount(10)
	fmt.Println(n)
}
