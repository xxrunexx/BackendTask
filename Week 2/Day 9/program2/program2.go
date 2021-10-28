package main

import "fmt"

func fibo(n int) int {
	// Bottom-up Tabulation
	// var memo map[int]int = map[int]int{}
	if n <= 1 {
		return n
	}
	nMinDua, nMinSatu, fiboN := 0, 1, 0
	for i := 2; i <= n; i++ {
		fiboN = nMinDua + nMinSatu
		nMinDua = nMinSatu
		nMinSatu = fiboN
	}
	return fiboN
}

func main() {
	fmt.Println(fibo(0))  // 0
	fmt.Println(fibo(1))  // 1
	fmt.Println(fibo(2))  // 1
	fmt.Println(fibo(3))  // 2
	fmt.Println(fibo(5))  // 5
	fmt.Println(fibo(6))  // 8
	fmt.Println(fibo(7))  // 13
	fmt.Println(fibo(9))  // 34
	fmt.Println(fibo(10)) // 55
}
