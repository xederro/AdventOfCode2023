package day9

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Part1() {
	scanner := utils.ReadFile("day9")
	count := 0

	compile, err := regexp.Compile(" +")
	if err != nil {
		return
	}

	for scanner.Scan() {
		var row []int
		s := compile.Split(scanner.Text(), -1)

		for _, v := range s {
			atoi, err := strconv.Atoi(v)
			if err != nil {
				return
			}
			row = append(row, atoi)
		}

		count += extraNext(row)
	}

	fmt.Println(count)
}

func extraNext(values []int) int {
	var rows [][]int
	rows = append(rows, values)
	i := 0
	for true {
		end := true
		for _, v := range rows[i] {
			if v != 0 {
				end = false
			}
		}
		if end {
			break
		}

		var row []int
		for j := 0; j < len(rows[i])-1; j++ {
			row = append(row, rows[i][j+1]-rows[i][j])
		}
		rows = append(rows, row)
		i++
	}
	for j := i - 1; j >= 0; j-- {
		v := rows[j][len(rows[j])-1] + rows[j+1][len(rows[j+1])-1]
		rows[j] = append(rows[j], v)
	}
	return rows[0][len(rows[0])-1]
}
