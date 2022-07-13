package utils

import (
	"sort"
)

type Cards []uint8

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func IsStraight(cards Cards) bool {
	if len(cards) < 5 || len(cards) > 7 {
		return false
	}

	sort.Sort(cards)

	// Get first card
	previousCard := cards[0]

	// if exist AS, then is considered as 1 for first card
	if cards[len(cards)-1] == 14 {
		previousCard = 1
	}

	validCardsCounter := 1
	for _, card := range cards {
		if previousCard+1 == card || previousCard == card {
			validCardsCounter++
		} else {
			validCardsCounter = 1
		}

		if validCardsCounter == 5 {
			return true
		}

		previousCard = card
	}

	return false
}
