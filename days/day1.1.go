package days

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
)

func Day1z1() {
	scanner := utils.ReadFile("day1")
	sum := 0

	for scanner.Scan() {
		s := scanner.Text()
		var first int = 10
		var last int = 10

		for _, r := range s {
			atoi, err := strconv.Atoi(string(r))
			if err != nil {
				continue
			}
			if first == 10 {
				first = atoi
			}
			last = atoi
		}
		sum += first*10 + last
	}
	fmt.Println(sum)
}
