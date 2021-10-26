package main

import (
	"fmt"
)

func pangkat(base, pangkat int) int {
	// code Here
	hasil := 1
	for i := 1; i <= pangkat; i++ {
		hasil = hasil * base
	}
	return hasil
}

func main() {
	var x, n int
	fmt.Println(pangkat(2, 3))  // 8
	fmt.Println(pangkat(5, 3))  // 125
	fmt.Println(pangkat(10, 2)) // 100
	fmt.Println(pangkat(2, 5))  // 32
	fmt.Println(pangkat(7, 3))  // 343

	// Input nilai x
	fmt.Print("Input Angka : ")
	fmt.Scanf("%d", &x)
	// Input nilai n
	fmt.Print("Input Pangkat : ")
	fmt.Scanf("%d", &n)
	fmt.Println(pangkat(x, n))
}
