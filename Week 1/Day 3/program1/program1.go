package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) bool {
	// if number <= 1 {
	// 	fmt.Print(number, " Bukan Bilangan Prima")
	// } else {
	// 	fmt.Print(number, " Bilangan Prima")
	// }
	if number <= 1 {
		fmt.Print(number, " Bukan Bilangan Prima")
		return false
	}
	batas := math.Sqrt(float64(number))
	for i := 2; i <= int(batas); i++ {
		if number%i == 0 {
			fmt.Print(number, " Bukan Bilangan Prima ")
			return false
		}
	}

	fmt.Println(number, " Bilangan Prima")
	return true
}

func main() {
	var number int
	fmt.Println(primeNumber(1000000007))  // true
	fmt.Println(primeNumber(1500450271))  // true
	fmt.Println(primeNumber(1000000000))  // false
	fmt.Println(primeNumber(10000000019)) // true
	fmt.Println(primeNumber(10000000033)) // true
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanf("%d", &number)
	fmt.Print(primeNumber(number))
}
