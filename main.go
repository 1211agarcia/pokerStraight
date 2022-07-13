package main

import (
	"fmt"
	"github.com/1211agarcia/pokerStraight/utils"
)

func main() {

	cards := utils.Cards{2, 3, 4, 5, 14}

	if utils.IsStraight(cards) {
		fmt.Printf("%v is a straight", cards)
	} else {
		fmt.Printf("%v is not a straight", cards)
	}
}
