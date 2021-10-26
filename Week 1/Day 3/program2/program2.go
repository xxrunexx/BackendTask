package main

import (
	"fmt"
)

func pow(x, n int) int {
	hasil := 1
	if n == 0 {
		return hasil
	}
	for n >= 1 {
		if n%2 != 0 {
			hasil *= x
		}
		x *= x
		n /= 2
	}
	return hasil
}

func main() {
	var x, n int
	fmt.Println(pow(2, 3))  // 8
	fmt.Println(pow(7, 2))  // 49
	fmt.Println(pow(10, 5)) // 100000
	fmt.Println(pow(17, 6)) // 24137569
	fmt.Println(pow(5, 3))  // 125
	fmt.Print("Masukkan Input x : ")
	fmt.Scanf("%d", &x)
	fmt.Print("Masukkan Input n : ")
	fmt.Scanf("%d", &n)
	fmt.Println(pow(x, n))
}
