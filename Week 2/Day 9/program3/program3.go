package main

import (
	"fmt"
	"math"
)

func Frog(jumps []int) int {
	// your code here
	start, i, count := jumps[0], 0, 0
	for i < len(jumps) {
		singleJump := math.Abs(float64(jumps[i+1] - jumps[len(jumps)-1]))
		doubleJump := math.Abs(float64(jumps[i+2] - jumps[len(jumps)-1]))
		if singleJump < doubleJump {
			count += int(math.Abs(float64(start - jumps[i+1])))
			start = jumps[i+1]
			i += 1
		} else {
			count += int(math.Abs(float64(start - jumps[i+2])))
			start = jumps[i+2]
			i += 2
		}
		if i == len(jumps)-2 {
			count += int(math.Abs(float64(start - jumps[i+1])))
			return count
		} else if i == len(jumps)-3 {
			count += int(math.Abs(float64(start - jumps[i+2])))
			return count
		}
	}
	return count
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))         // 30
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50})) // 40
}

// 10			30			40			20
// jumps[0]		[1]			[2]			[3]
