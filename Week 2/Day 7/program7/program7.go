package main

import (
	"fmt"
)

type pairCards struct {
	card1 int
	card2 int
}

type output interface{}

func playingDomino(cards [][]int, deck []int) interface{} {
	closedCard := true
	var card []pairCards
	for i, _ := range cards {
		if cards[i][0] == deck[1] || cards[i][1] == deck[1] {
			card = append(card, pairCards{cards[i][0], cards[i][1]})
			closedCard = false
			break
		}
	}
	var hasil []output
	for _, i := range card {
		hasil = append(hasil, i.card1, i.card2)
	}
	if closedCard {
		return "Tutup Kartu"
	}
	return hasil
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3})) // [3, 4]
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6})) // [6 5]
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))              // "tutup kartu"
}
