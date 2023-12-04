package day1

import (
	"adventOfCode/utils"
	"fmt"
	"slices"
	"strconv"
)

func Part2() {
	var digits = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	scanner := utils.ReadFile("day1")
	sum := 0

	for scanner.Scan() {
		s := scanner.Text()
		first := 10
		last := 10
		firstIndex := 10
		lastIndex := 10

		all := [10][]int{}

		for k, v := range digits {
			all[k] = utils.FindIndices(s, strconv.Itoa(k))
			all[k] = append(all[k], utils.FindIndices(s, v)...)
		}

		for k := range all {
			if len(all[k]) != 0 {
				newFirst := slices.Min(all[k])
				newLast := slices.Max(all[k])
				if first == 10 || firstIndex > newFirst {
					firstIndex = newFirst
					first = k
				}

				if last == 10 || lastIndex < newLast {
					lastIndex = newLast
					last = k
				}
			}
		}

		sum += first*10 + last
	}
	fmt.Println(sum)
}
