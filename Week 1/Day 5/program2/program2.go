package main

import (
	"fmt"
)

func Caesar(offset int, input string) string {
	const lowerCase = "abcdefghijklmnopqrstuvwxyz"
	const upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	offset %= 26
	cipheredWord := []byte(input)

	for i, byteVal := range cipheredWord {
		if byteVal >= 'a' && byteVal <= 'z' {
			cipheredWord[i] = lowerCase[(int((26+(byteVal-'a')))+offset)%26]
		} else if byteVal >= 'A' && byteVal <= 'Z' {
			cipheredWord[i] = upperCase[(int((26+(byteVal-'a')))+offset)%26]
		} else {
			cipheredWord = nil
		}
	}
	return string(cipheredWord)
}

func main() {
	fmt.Println(Caesar(3, "abc"))
	// def
	fmt.Println(Caesar(2, "alta"))
	// cnvc
	fmt.Println(Caesar(10, "alterraacademy"))
	// kvdobbkkmknowi
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	// bcdefghijklmnopqrstuvwxyza
	fmt.Println(Caesar(1000, "abcdefghijklmnopqrstuvwxyz"))
	// mnopqrstuvwxyzabcdefghijkl
}
