package day23

import (
	"adventOfCode/utils"
	"fmt"
)

func Part1() {
	scanner := utils.ReadFile("day23")
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(s)
	}
}
