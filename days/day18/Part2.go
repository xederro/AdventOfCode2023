package day18

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type line struct {
	Start [2]int
	Stop  [2]int
}

// https://en.wikipedia.org/wiki/Shoelace_formula
func Part2() {
	var lines []*line
	scanner := utils.ReadFile("day18")
	last := &line{Stop: [2]int{0, 0}}

	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, " ")

		ohMyGod := strings.Trim(split[2], "(#)")

		atoi, err := strconv.ParseInt(ohMyGod[:len(ohMyGod)-1], 16, 32)
		if err != nil {
			return
		}

		a := int(atoi)

		d := line{Start: last.Stop}
		switch ohMyGod[len(ohMyGod)-1:] {
		case "2":
			d.Stop = [2]int{d.Start[0], d.Start[1] - a}
			break
		case "0":
			d.Stop = [2]int{d.Start[0], d.Start[1] + a}
			break
		case "3":
			d.Stop = [2]int{d.Start[0] - a, d.Start[1]}
			break
		case "1":
			d.Stop = [2]int{d.Start[0] + a, d.Start[1]}
			break
		}

		last = &d
		lines = append(lines, &d)
	}

	fmt.Println(countInside(lines))
}

func countInside(lines []*line) int {
	s := 0
	for i := 0; i < len(lines); i++ {
		s += (lines[i].Start[0] + lines[i].Stop[0]) * (lines[i].Start[1] - lines[i].Stop[1])
	}

	for _, l := range lines {
		if l.Start[1] == l.Stop[1] {
			s += int(math.Abs(float64(l.Start[0] - l.Stop[0])))
		} else {
			s += int(math.Abs(float64(l.Start[1] - l.Stop[1])))
		}
	}
	return s/2 + 1
}
