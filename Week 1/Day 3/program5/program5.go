package main

import (
	"fmt"
)

func pairSum(arr []int, target int) []int {
	// mapping := make(map[int]int)
	var hasilArr []int
Loop:
	for i, val1 := range arr {
		for j, val2 := range arr {
			check := val1 + val2 // Join val1+val2 to check
			if target == check { // Check similiraties target & check
				hasilArr = append(hasilArr, i)
				hasilArr = append(hasilArr, j)
				break Loop
			}
		}
	}
	// Error
	// for i := 0; i < len(arr); i++ {
	// 	if mapping[target-arr[i]] == 0 {
	// 		mapping[arr[i]] = i
	// 	} else {
	// 		hasilArr = append(hasilArr, arr[mapping[target-arr[i]]])
	// 		hasilArr = append(hasilArr, arr[i])
	// 	}
	// }
	return hasilArr
}

func main() {
	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6)) // [1, 3]
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))  // [0, 2]
	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12))   // [2, 3]
	fmt.Println(pairSum([]int{1, 4, 6, 8}, 10))   // [1, 2]
	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6))    // [0, 1]
}
