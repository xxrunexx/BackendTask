package main

import (
	"fmt"
)

func arrayMerge(arrayA, arrayB []string) []string {
	merge := append(arrayA, arrayB...)
	check := make(map[string]int)
	res := make([]string, 0, 0)
	for _, val := range merge {
		check[val] = 1
	}
	// fmt.Printf("%+v\n", check)
	for letter := range check {
		res = append(res, letter)
	}

	return res
}

func main() {

	fmt.Println(arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	// ["king", "devil jin", "akuma", "eddie", "steve", "geese"]

	fmt.Println(arrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	// ["sergei", "jin", "steve", "bryan"]

	fmt.Println(arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	// ["alisa", "yoshimitsu", "devil jin", "law"]

	fmt.Println(arrayMerge([]string{}, []string{"devil jin", "sergei"}))
	// ["devil jin", "sergei"]

	fmt.Println(arrayMerge([]string{"hwoarang"}, []string{}))
	// ["hwoarang"]

	fmt.Println(arrayMerge([]string{}, []string{}))
	// []
}
