package day15

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

func Part1() {
	scanner := utils.ReadFile("day15")
	var res int

	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, ",")
		for _, str := range split {
			res += hash(str)
		}
	}

	fmt.Println(res)
}

func hash(str string) int {
	res := 0
	for _, v := range str {
		res = ((res + int(v)) * 17) % 256
	}
	return res
}
