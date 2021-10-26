package main

import (
	"fmt"
	"os"
)

func cetakTablePerkalian(number int) {
	for i := 1; i <= number; i++ {
		for j := 1; j <= number; j++ {
			n := i * j
			if n < 10 {
				fmt.Print(n, "  ")
			} else {
				fmt.Print(n, " ")
			}
		}
		fmt.Println("")
	}
}

func main() {
	var n int
	cetakTablePerkalian(9)
	fmt.Printf("Masukkan jumlah n : ")
	fmt.Scanf("%d", &n)
	// Restrict input only 1-30
	if n <= 30 {
		cetakTablePerkalian(n)
	} else {
		fmt.Println("Input hanya dari 1-30!")
		os.Exit(0)
	}
}
