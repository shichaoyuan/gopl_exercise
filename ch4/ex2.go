package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var a = flag.String("a", "sha256", "algorithm of SHA")

func main() {
	flag.Parse()
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "read stdin: %v\n", err)
	}
	switch *a {
	case "sha256":
		fmt.Fprintf(os.Stdout, "%x", sha256.Sum256(b))
	case "sha384":
		fmt.Fprintf(os.Stdout, "%x", sha512.Sum384(b))
	case "sha512":
		fmt.Fprintf(os.Stdout, "%x", sha512.Sum512(b))
	default:
		fmt.Fprintf(os.Stderr, "Unsupported Algorithm: %s\n", *a)
	}
}
