package day7

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type hand struct {
	Points int
	Bid    int
	Str    string
}

var cards = [13]string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}

func Part1() {
	scanner := utils.ReadFile("day7")
	count := 0
	var hands []hand

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
		points := 1 << (4 * 5)
		pointsInCard := 0

		for i := 0; i < len(cards); i++ {
			temp := utils.FindIndices(s[0], cards[i])
			if len(temp) > 0 {
				points *= len(temp) + 1
				for _, pos := range temp {
					pointsInCard += (13 - i) << (4 * (5 - pos - 1))
				}
			}
		}

		hands = append(hands, hand{Bid: atoi, Points: points + pointsInCard, Str: s[0]})
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Points > hands[j].Points
	})

	for i, h := range hands {
		count += (i + 1) * h.Bid
	}

	fmt.Println(count)
}
