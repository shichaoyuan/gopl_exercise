// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func Test_Len(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(64)
	x.Add(13)
	x.Add(32)

	fmt.Println("Test_Len")
	fmt.Println(x.Len())

}

func Test_Remove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(64)
	x.Add(364)

	fmt.Println(x.String())

	x.Remove(64)
	x.Remove(32)

	fmt.Println(x.String())

}

func Test_Clear(t *testing.T) {
	var x IntSet
	x.Add(1)

	fmt.Println(x.String())

	x.Clear()

	fmt.Println(x.String())
}

func Test_Copy(t *testing.T) {
	var x IntSet
	x.Add(1)

	fmt.Println(x.String())

	r := x.Copy()
	r.Add(2)

	fmt.Println(x.String())
	fmt.Println(r.String())
}

func Test_AddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 4, 5, 7)

	fmt.Println(x.String())
}

func Test_IntersectWith(t *testing.T) {
	var x IntSet
	x.AddAll(1, 4, 5, 7, 200)

	var y IntSet
	y.AddAll(1, 2, 5, 7, 10, 33, 100, 30)

	x.IntersectWith(&y)

	fmt.Println(x.String())
}

func Test_Elems(t *testing.T) {
	var x IntSet
	x.AddAll(1, 4, 5, 7, 200)

	for _, e := range x.Elems() {
		fmt.Printf("%d ", e)
	}
	fmt.Println()
}
