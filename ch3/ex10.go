package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}
	var buf bytes.Buffer
	r := l % 3
	if r > 0 {
		buf.WriteString(s[:r])
		buf.WriteByte(',')
	}
	for ; r < l; r += 3 {
		buf.WriteString(s[r : r+3])
		if r+3 < l {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func main() {
	fmt.Println(comma("12"))
	fmt.Println(comma("123"))
	fmt.Println(comma("1234"))
	fmt.Println(comma("123456789"))
	fmt.Println(comma("1234567890"))
}
