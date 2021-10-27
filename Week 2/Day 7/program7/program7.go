package main

import (
	"fmt"
)

func playingDomino(cards [][]int, deck []int) interface{} {
	closedCard := true
	for i, _ := range cards {
		if cards[i][0] == deck[1] || cards[i][1] == deck[1] {
			fmt.Print("[", cards[i][0], ", ", cards[i][1], "] ")
			closedCard = false
			break
		}
	}
	if closedCard {
		fmt.Print("Tutup Kartu")
	}
	return ""
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})) // [3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})) // [6 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))              // "tutup kartu"
}
