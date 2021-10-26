package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	// your code here
	prime := []int{}
	for i := 1; true; i++ {
		if primeNumber(i) {
			prime = append(prime, i)
		}
		if len(prime) == number {
			break
		}
	}
	return prime[number-1]
}
func primeNumber(number int) bool {
	if number <= 1 {
		return false
	} else if number == 2 {
		return true
	}
	batas := math.Sqrt(float64(number))
	for i := 2; i <= int(batas)+1; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(primeX(1))  // 2
	fmt.Println(primeX(5))  // 11
	fmt.Println(primeX(8))  // 19
	fmt.Println(primeX(9))  // 23
	fmt.Println(primeX(10)) // 29
}
