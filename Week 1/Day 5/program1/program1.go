package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	var c string
	if strings.Contains(b, a) {
		c := a
		return c
	} else if strings.Contains(a, b) {
		c := b
		return c
	}
	return c
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))     // AKA
	fmt.Println(Compare("KANGOORO", "KANG"))  // KANG
	fmt.Println(Compare("KI", "KIJANG"))      // KI
	fmt.Println(Compare("KUPU-KUPU", "KUPU")) // KUPU
	fmt.Println(Compare("ILALANG", "ILA"))    // ILA
}
