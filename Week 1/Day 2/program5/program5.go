package main

import (
	"fmt"
)

func countPalindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			fmt.Print(input, " Bukan Palindrome ")
			return false
		}
	}
	fmt.Print(input, " Palindrome ")
	return true
}

func main() {
	var input string

	fmt.Println(countPalindrome("civic"))       // true
	fmt.Println(countPalindrome("katak"))       // true
	fmt.Println(countPalindrome("kasur rusak")) // true
	fmt.Println(countPalindrome("mistar"))      // false
	fmt.Println(countPalindrome("lion"))        // false

	fmt.Print("Input Word : ")
	fmt.Scanf("%s", &input)
	fmt.Println(countPalindrome(input))
}
