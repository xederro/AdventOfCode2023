package day13

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

func Part2() {
	scanner := utils.ReadFile("day13")
	var table [][]bool
	count := 0

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			count += finMirrorSmudge(&table)
			table = [][]bool{}
			continue
		}
		var row []bool
		split := strings.Split(s, "")
		for _, v := range split {
			isRock := false
			if v == "#" {
				isRock = true
			}
			row = append(row, isRock)
		}
		table = append(table, row)
	}
	count += finMirrorSmudge(&table)

	fmt.Println(count)
}

func finMirrorSmudge(tab *[][]bool) int {
	lenX := len((*tab)[0])
	lenY := len(*tab)
	res := 0

Row:
	for i := 0; i < lenX-1; i++ {
		foundSmudge := false
		for j := 0; j < lenY; j++ {
			for k := 0; k < min(i+1, lenX-i-1); k++ {
				a := (*tab)[j][i-k]
				b := (*tab)[j][i+k+1]

				if a != b && foundSmudge {
					continue Row
				}
				if a != b {
					foundSmudge = true
				}
			}
		}
		if foundSmudge {
			res += i + 1
			break
		}
	}

Col:
	for i := 0; i < lenY-1; i++ {
		foundSmudge := false
		for j := 0; j < lenX; j++ {
			for k := 0; k < min(i+1, lenY-i-1); k++ {
				a := (*tab)[i-k][j]
				b := (*tab)[i+k+1][j]

				if a != b && foundSmudge {
					continue Col
				}
				if a != b {
					foundSmudge = true
				}
			}
		}
		if foundSmudge {
			res += (i + 1) * 100
			break
		}
	}

	return res
}
