package day3

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Part1() {
	scanner := utils.ReadFile("day3")
	sum := 0
	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	for i := 0; i < len(input); i++ {
		var beginIndex *[2]int
		var endIndex *[2]int
		for j := 0; j < len(input[i]); j++ {
			switch input[i][j] {
			default:
				if beginIndex != nil && endIndex != nil {
					sum += find(*beginIndex, *endIndex, input)
					beginIndex = nil
					endIndex = nil
				}

			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if beginIndex == nil {
					beginIndex = &[2]int{i, j}
					endIndex = &[2]int{i, j}
				} else {
					endIndex = &[2]int{i, j}
				}
			}
		}
		if beginIndex != nil && endIndex != nil {
			sum += find(*beginIndex, *endIndex, input)
		}
	}
	fmt.Println(sum)
}

func find(begin [2]int, end [2]int, field []string) int {
	area := [4]int{begin[0], begin[1], end[0], end[1]}

	c, err := regexp.Compile(`[\!\@\#\$\%\^\&\*\+\=\-\/]`)
	if err != nil {
		return 0
	}

	if area[0]-1 >= 0 {
		area[0]--
	}
	if area[1]-1 >= 0 {
		area[1]--
	}
	if area[2]+1 < len(field) {
		area[2]++
	}
	if area[3]+1 < len(field[0]) {
		area[3]++
	}

	for i := area[0]; i <= area[2]; i++ {
		if c.Match([]byte(field[i][area[1] : area[3]+1])) {
			atoi, err := strconv.Atoi(field[begin[0]][begin[1] : end[1]+1])
			if err != nil {
				return 0
			}
			return atoi
		}
	}
	return 0
}
