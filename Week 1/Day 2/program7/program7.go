package main

import (
	"fmt"
)

func playWithAsterisk(n int) {
	var temp int
	for i := 0; i < n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k <= i; k++ {
			if k == 0 || i == 0 {
				temp = 1
			} else {
				temp = temp * (i - k + 1) / k
			}
			fmt.Print(" *")
		}
		fmt.Println("")
	}
}

func main() {
	var n int
	playWithAsterisk((5))
	fmt.Print("Masukkan banyak n : ")
	fmt.Scanf("%d", &n)
	playWithAsterisk(n)
}
