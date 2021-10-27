package main

import (
	"fmt"
	"strings"
)

func main() {
	// 4. Substring
	value := "cat;dog"
	// Take substring from index 4 to length of string.
	substring := value[4:]
	fmt.Println(substring)
	// 5. Replace
	s := "this[things]I would like to remove"
	t := strings.Replace(s, "[things]", " ", -1)
	fmt.Printf("%s\n", t)

	// 6. Insert
	p := "green"
	index := 2
	q := p[:index] + "HI" + p[1:]
	fmt.Println(p, q)
}
