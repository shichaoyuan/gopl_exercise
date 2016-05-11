package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(expand("$test haha $world", reverse))
}

func expand(s string, f func(string) string) string {
	re := regexp.MustCompile(`\$[a-zA-Z0-9]+`)
	return re.ReplaceAllStringFunc(s, f)
}

func reverse(s string) string {
	runes := []rune(s[1:])
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
