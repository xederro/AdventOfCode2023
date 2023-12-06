package day6

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type race2 struct {
	Time     int64
	Distance int64
}

func Part2() {
	scanner := utils.ReadFile("day6")
	race := race2{}

	compile, err := regexp.Compile(" +")
	if err != nil {
		return
	}

	if scanner.Scan() {
		s := compile.ReplaceAllString(scanner.Text(), "")
		atoi, err := strconv.Atoi(s[5:])
		if err != nil {
			return
		}
		race.Time = int64(atoi)
	}
	if scanner.Scan() {
		s := compile.ReplaceAllString(scanner.Text(), "")
		atoi, err := strconv.Atoi(s[9:])
		if err != nil {
			return
		}
		race.Distance = int64(atoi)
	}

	fmt.Println(race.CountWaysToWin())
}

func (r race2) CountWaysToWin() int {
	D := math.Sqrt(math.Pow(float64(r.Time), 2.0) - 4.0*float64(r.Distance))
	z1 := math.Ceil(((float64(r.Time) + D) / 2.0) - 1)
	z2 := math.Floor(((float64(r.Time) - D) / 2.0) + 1)
	return int(z1 - z2 + 1)
}
