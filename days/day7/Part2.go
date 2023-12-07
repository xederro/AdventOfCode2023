package day7

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
)

type hand2 struct {
	Points int
	Bid    int
	Str    string
}

var cards2 = [13]string{"J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"}

func Part2() {
	scanner := utils.ReadFile("day7")
	count := 0
	var hands []hand2

	compile, err := regexp.Compile(" +")
	if err != nil {
		return
	}

	for scanner.Scan() {
		s := compile.Split(scanner.Text(), -1)
		atoi, err := strconv.Atoi(s[1])
		if err != nil {
			return
		}
		points := 10000000000
		pointsInCard := 0

		var handCards []int
		var joker int

		for i := 0; i < len(cards2); i++ {
			temp := utils.FindIndices(s[0], cards2[i])
			if len(temp) > 0 {
				if i == 0 {
					joker = len(temp)
				} else {
					handCards = append(handCards, len(temp)+1)
				}
				for _, pos := range temp {
					pointsInCard += (13 - i) * int(math.Pow(100.0, float64(5.0-pos-1)))
				}
			}
		}

		sort.Slice(handCards, func(i, j int) bool {
			return handCards[i] > handCards[j]
		})

		if len(handCards) == 0 {
			handCards = append(handCards, joker+1)
		} else {
			handCards[0] += joker
		}

		for _, v := range handCards {
			points *= v
		}

		hands = append(hands, hand2{Bid: atoi, Points: points + pointsInCard, Str: s[0]})
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Points > hands[j].Points
	})

	for i, h := range hands {
		count += (i + 1) * h.Bid
	}

	fmt.Println(count)
}
