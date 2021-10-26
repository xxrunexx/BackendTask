package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	// your code here
	prime := []int{} //Store prime to slice
	for i := 2; true; i++ {
		if primeNumber(i) {
			prime = append(prime, i)
		}
		if len(prime) == number {
			//2,3,5,7,11
			break
		}
	}
	return prime[number-1]
}
func primeNumber(number int) bool {
	if number < 2 {
		return false
	} else if number == 2 {
		return true
	}
	akar := math.Sqrt(float64(number))
	for i := 2; i <= int(akar); i++ {
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
