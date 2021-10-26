package main

import (
	"fmt"
	"math"
)

func primaSegiEmpat(high, wide, start int) {

	area := high * wide
	prime := []int{}
	for i := start + 1; true; i++ {
		if primeNumber(i) {
			prime = append(prime, i) //masukin slice bil prima sampe 1-1000
		}
		if len(prime) == area {
			break
		}
	}
	sum := 0
	i := 0
	for i < area {
		if i%high == 0 {
			fmt.Println()
		}
		sum += prime[i]
		fmt.Printf("%d\t", prime[i])
		i++
	}
	fmt.Printf("\n%d\n", sum)
}

func primeNumber(number int) bool {
	if number < 2 {
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

	primaSegiEmpat(2, 3, 13)
	/*
	   17 19
	   23 29
	   31 37
	   156
	*/

	primaSegiEmpat(5, 2, 1)
	/*
	   2  3  5  7 11
	   13 17 19 23 29
	   129
	*/

}
