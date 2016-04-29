package main

import "fmt"

func eliAdjDup(strings []string) []string {
	if len(strings) < 2 {
		return strings
	}

	out := strings[:0]
	pre := strings[0]
	out = append(out, pre)
	for _, s := range strings[1:] {
		if pre != s {
			out = append(out, s)
			pre = s
		}
	}
	return out
}

func main() {
	s := []string{"a", "b", "c", "c", "e", "e", "e", "f", "a"}
	fmt.Println(s)
	s = eliAdjDup(s)
	fmt.Println(s)
}
