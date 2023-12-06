package day6

import (
	"adventOfCode/utils"
	"fmt"
	"math"
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

// look: https://www.desmos.com/calculator/94kippzj2t
func (r race) CountWaysToWin() int {
	D := math.Sqrt(math.Pow(float64(r.Time), 2.0) - 4.0*float64(r.Distance))
	z1 := math.Ceil(((float64(r.Time) + D) / 2.0) - 1)
	z2 := math.Floor(((float64(r.Time) - D) / 2.0) + 1)
	return int(z1 - z2 + 1)
}
