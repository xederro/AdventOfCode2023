package day14

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

func Part1() {
	scanner := utils.ReadFile("day14")
	var dish [][]int

	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, "")
		var row []int
		for _, c := range split {
			switch c {
			case ".":
				row = append(row, 0)
				break
			case "O":
				row = append(row, 1)
				break
			case "#":
				row = append(row, 2)
				break
			}
		}
		dish = append(dish, row)
	}

	roll(&dish)

	fmt.Println(calc(&dish))
}

func roll(dish *[][]int) {
	for i := 1; i < len(*dish); i++ {
		for j := i; j > 0; j-- {
			for k := 0; k < len((*dish)[i]); k++ {
				if (*dish)[j-1][k] == 0 && (*dish)[j][k] == 1 {
					(*dish)[j-1][k] = 1
					(*dish)[j][k] = 0
				}
			}
		}
	}
}

func calc(dish *[][]int) int {
	rows := len(*dish)
	res := 0
	for i := 0; i < len(*dish); i++ {
		for k := 0; k < len((*dish)[i]); k++ {
			if (*dish)[i][k] == 1 {
				res += rows - i
			}
		}
	}
	return res
}
