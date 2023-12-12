package day12

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func Part2() {
	var possible int = 0
	scanner := utils.ReadFile("day12")

	digits, err := regexp.Compile(`\d+`)
	if err != nil {
		return
	}

	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, " ")

		split[0] = strings.Repeat(split[0]+"?", 5)
		split[1] = strings.Repeat(split[1]+" ", 5)

		go func() {
			wg.Add(1)
			possible += calc(
				split[0][:len(split[0])-1]+".",
				toInt2(digits.FindAllString(split[1][:len(split[1])-1], -1)),
			)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(possible)
}

func toInt2(s []string) []int {
	var i []int
	for _, str := range s {
		atoi, err := strconv.Atoi(str)
		if err != nil {
			return nil
		}
		i = append(i, atoi)
	}
	return i
}

func calc(line string, dig []int) int {
	var table [][][]int

	for i := 0; i < len(line)+1; i++ {
		var tab [][]int
		for j := 0; j < len(dig)+2; j++ {
			var row []int
			for o := 0; o < len(line)+2; o++ {
				row = append(row, 0)
			}
			tab = append(tab, row)
		}
		table = append(table, tab)
	}

	table[0][0][0] = 1

	for i := 0; i < len(line); i++ {
		for j := 0; j < len(dig)+1; j++ {
			for k := 0; k < len(line)+1; k++ {
				cur := table[i][j][k]
				if cur == 0 {
					continue
				}
				if line[i] == '.' || line[i] == '?' {
					if k == 0 || k == dig[j-1] {
						table[i+1][j][0] += cur
					}
				}
				if line[i] == '#' || line[i] == '?' {
					if k == 0 {
						table[i+1][j+1][k+1] += cur
					} else {
						table[i+1][j][k+1] += cur
					}
				}
			}
		}
	}
	return table[len(line)][len(dig)][0]
}
