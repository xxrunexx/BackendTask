package main

import "fmt"

func BinarySearch(array []int, x int) {
	left := 0
	right := len(array) - 1
	for left < right {
		mid := (left + right) / 2
		if array[mid] == x {
			left = mid
			right = mid
		}
		if array[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == len(array) || array[left] != x {
		left = -1
	}
	fmt.Println(left)
}

func main() {
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3)                  // 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5)                 // 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53)  // 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // -1
}
