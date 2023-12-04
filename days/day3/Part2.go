package day3

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
)

func Part2() {
	scanner := utils.ReadFile("day3")
	sum := 0
	input := []string{}
	gears := map[[2]int][2]int{}

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
					temp, atoi := findStar(*beginIndex, *endIndex, input)
					for _, t := range temp {
						if _, ok := gears[t]; ok {
							gears[t] = [2]int{gears[t][0] + 1, gears[t][1] * atoi}
						} else {
							gears[t] = [2]int{1, atoi}
						}
					}
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
			temp, atoi := findStar(*beginIndex, *endIndex, input)
			for _, t := range temp {
				if _, ok := gears[t]; ok {
					gears[t] = [2]int{gears[t][0] + 1, gears[t][1] * atoi}
				} else {
					gears[t] = [2]int{1, atoi}
				}
			}
		}
	}

	for _, v := range gears {
		if v[0] == 2 {
			sum += v[1]
		}
	}

	fmt.Println(sum)
}

func findStar(begin [2]int, end [2]int, field []string) ([][2]int, int) {
	area := [4]int{begin[0], begin[1], end[0], end[1]}

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

	j := [][2]int{}

	for i := area[0]; i <= area[2]; i++ {
		temp := utils.FindIndices(field[i][area[1]:area[3]+1], "*")
		for _, t := range temp {
			j = append(j, [2]int{i, area[1] + t})
		}
	}

	if len(j) != 0 {
		atoi, err := strconv.Atoi(field[begin[0]][begin[1] : end[1]+1])
		if err != nil {
			return [][2]int{}, 0
		}
		return j, atoi
	}

	return [][2]int{}, 0
}
