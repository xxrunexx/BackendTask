package main

import (
	"fmt"
)

func MaxSequence(arr []int) int {
	// your code here
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return arr[0]
	}
	maxSum, result := arr[0], 0
	for i := 0; i < len(arr)-1; i++ {
		result += arr[i]
		if result > maxSum {
			maxSum = result
		}
		if result < 0 {
			result = 0
		}
	}
	return maxSum
}

func main() {
	// fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6
	// fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))   // 7
	// fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))   // 7
	// fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))   // 8
	// fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))    // 12
	fmt.Println(MaxSequence([]int{-1, 3, 4, -2, 5, -7})) // 10
}
