package day14

import (
	"adventOfCode/utils"
	"fmt"
	"slices"
	"strings"
)

func Part2() {
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

	var cycle []int
	var c []int
	var start int

	for i := 0; i < 1000000000; i++ {
		rollRound(&dish)
		//print(dish)
		cycle = append(cycle, calc2(&dish))
		c, start = findCycle(cycle)
		if len(c) != 0 {
			break
		}
	}

	fmt.Println(c[(1000000000-start-1)%len(c)])
}

func rollRound(dish *[][]int) {
	//north
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
	//west
	for i := 1; i < len((*dish)[0]); i++ {
		for j := i; j > 0; j-- {
			for k := 0; k < len(*dish); k++ {
				if (*dish)[k][j-1] == 0 && (*dish)[k][j] == 1 {
					(*dish)[k][j-1] = 1
					(*dish)[k][j] = 0
				}
			}
		}
	}
	//south
	for i := len(*dish); i >= 0; i-- {
		for j := i; j < len(*dish)-1; j++ {
			for k := 0; k < len((*dish)[j]); k++ {
				if (*dish)[j+1][k] == 0 && (*dish)[j][k] == 1 {
					(*dish)[j+1][k] = 1
					(*dish)[j][k] = 0
				}
			}
		}
	}
	//east
	for i := len((*dish)[0]); i >= 0; i-- {
		for j := i; j < len((*dish)[0])-1; j++ {
			for k := 0; k < len(*dish); k++ {
				if (*dish)[k][j+1] == 0 && (*dish)[k][j] == 1 {
					(*dish)[k][j+1] = 1
					(*dish)[k][j] = 0
				}
			}
		}
	}
}

func calc2(dish *[][]int) int {
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

func findCycle(cycle []int) ([]int, int) {
	for i := len(cycle) / 2; i > 2; i-- {
		if slices.Equal(cycle[len(cycle)-i:], cycle[len(cycle)-2*i:len(cycle)-2*i+i]) {
			return cycle[len(cycle)-i:], len(cycle) - 2*i
		}
	}
	return nil, 0
}
