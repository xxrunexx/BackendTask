package main

import "fmt"

// import (
// 	"fmt"
// )

func loop(number int) int {
	for i := 0; i <= number; i++ {
		if number > 100 {
			panic("TOO BIG!")
		}
	}
	return number
}

func main() {
	fmt.Println(loop(1000))
}
