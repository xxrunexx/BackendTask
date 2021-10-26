package main

import (
	"fmt"
)

const PHI = 3.14

func main() {
	var t, r float64
	fmt.Print("Masukkan Nilai Tinggi : ")
	fmt.Scan(&t)
	fmt.Print("Masukkan Nilai Jari-jari : ")
	fmt.Scan(&r)
	hasil := (2 * PHI * r) * (r + t)
	fmt.Println("Luas Permukaan Tabung : ", hasil)
}

// Expected Input
// t : 20.0, r : 4
// Expected Output
// 602.88
