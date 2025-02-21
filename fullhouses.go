package main

import (
	"fmt"
)

func main() {

	var cards = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10, 11}
	var sets [][]int
	var hands []string

	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			if 3*(cards[i])+2*(cards[j]) >= 35 {
				sets = append(sets, []int{cards[i], cards[j]})
			}
			if 3*(cards[j])+2*(cards[i]) >= 35 {
				sets = append(sets, []int{cards[j], cards[i]})
			}
		}
	}

	for _, set := range sets {
		hands = append(hands, fmt.Sprintf("%ds full of %ds \n", set[0], set[1]))
	}

	fmt.Println(hands, len(sets))
}
