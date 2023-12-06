package day6

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
)

type race struct {
	Time     int
	Distance int
}

func Part1() {
	scanner := utils.ReadFile("day6")
	ways := 1
	var races []race

	compile, err := regexp.Compile(" +")
	if err != nil {
		return
	}

	if scanner.Scan() {
		s := compile.Split(scanner.Text(), -1)
		for i := 1; i < len(s); i++ {
			atoi, err := strconv.Atoi(s[i])
			if err != nil {
				return
			}
			races = append(races, race{Time: atoi})
		}
	}
	if scanner.Scan() {
		s := compile.Split(scanner.Text(), -1)
		for i := 1; i < len(s); i++ {
			atoi, err := strconv.Atoi(s[i])
			if err != nil {
				return
			}
			races[i-1].Distance = atoi
		}
	}

	for _, r := range races {
		ways *= r.CountWaysToWin()
	}

	fmt.Println(ways)
}

func (r race) CountWaysToWin() int {
	ways := 0

	for i := 0; i < r.Time; i++ {
		if i*(r.Time-i) > r.Distance {
			ways++
		}
	}

	return ways
}
