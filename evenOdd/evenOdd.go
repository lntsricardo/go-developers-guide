package main

import "fmt"

// Creates a new type of evenOdd, which is a slice of int
type evenOdd []int

func newEvenOddSlice() evenOdd {
	return evenOdd{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

func (e evenOdd) printEvenOdd() {
	for _, eo := range e {
		if eo%2 == 0 {
			fmt.Println(eo, "is even")
		} else {
			fmt.Println(eo, "is odd")
		}
	}
}
