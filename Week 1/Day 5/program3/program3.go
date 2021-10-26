package main

import (
	"fmt"
)

func Swap(a, b *int) {
	temp := *a // Move pointer a -> temp
	*a = *b    // Move b pointer -> pointer a
	*b = temp  // Move temp value -> b pointer

}

func main() {
	a := 10
	b := 20
	Swap(&a, &b)
	fmt.Printf("Nilai a : %d\nNilai b : %d\n", a, b)
}
