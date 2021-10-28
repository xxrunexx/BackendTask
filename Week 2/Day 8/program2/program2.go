package main

import "fmt"

func MoneyCoins(money int) []int {
	pecahan := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	hasil := []int{}
	for i := len(pecahan) - 1; i >= 0; i-- {
		for money >= pecahan[i] {
			money -= pecahan[i]
			hasil = append(hasil, pecahan[i])
		}
		if money == 0 {
			return hasil
		}
	}
	return hasil
}

func main() {
	fmt.Println(MoneyCoins(123))   // [100 20 1 1 1]
	fmt.Println(MoneyCoins(432))   // [200 200 20 10 1 1]
	fmt.Println(MoneyCoins(543))   // [500, 20, 20, 1, 1, 1]
	fmt.Println(MoneyCoins(7752))  // [5000, 2000, 500, 200, 50, 1, 1]
	fmt.Println(MoneyCoins(15321)) // [10000 5000 200 100 20 1]
}
