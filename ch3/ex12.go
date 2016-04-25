package main

import (
	"fmt"
	"sort"
)

func anagram(a, b string) bool {
	return normal(a) == normal(b)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func normal(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	fmt.Println(anagram("eata", "aaet"))
	fmt.Println(anagram("eat", "tae"))
	fmt.Println(anagram("eat", "taa"))
}
