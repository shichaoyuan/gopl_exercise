package main

import "fmt"

func rotate(s []int, n int) {
	tmp := make([]int, n)
	copy(tmp, s[:n])
	copy(s, s[n:])
	copy(s[len(s)-n:], tmp)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	rotate(s, 1)
	fmt.Println(s)
}
