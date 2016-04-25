package main

import (
	"bytes"
	"fmt"
	"strings"
)

func commaAfterDot(s string) string {
	l := len(s)
	if l <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i := 0; i < l; i += 3 {
		if i+3 < l {
			buf.WriteString(s[i : i+3])
			buf.WriteByte(',')
		} else {
			buf.WriteString(s[i:l])
		}
	}
	return buf.String()
}

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

func enhanceComma(s string) string {
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		s = s[1:]
	}
	dot := strings.Index(s, ".")
	if dot != -1 {
		buf.WriteString(comma(s[0:dot]))
		buf.WriteByte('.')
		buf.WriteString(commaAfterDot(s[dot+1 : len(s)]))
	} else {
		buf.WriteString(comma(s))
	}
	return buf.String()
}

func main() {
	fmt.Println(enhanceComma("-123"))
	fmt.Println(enhanceComma("-1234"))
	fmt.Println(enhanceComma("-1234.56789"))
	fmt.Println(enhanceComma("12345.67890"))
}
