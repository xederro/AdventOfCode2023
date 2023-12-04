package days

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"regexp"
	"strings"
)

type ticket struct {
	Winning []string
	Numbers []string
}

func Day4z1() {
	scanner := utils.ReadFile("day4")
	var tickets []ticket
	sum := 0.0

	compile, err := regexp.Compile("[ ]+")
	if err != nil {
		return
	}

	for scanner.Scan() {
		s := scanner.Text()
		splited := strings.Split(strings.Split(s, ":")[1], "|")
		tickets = append(tickets, ticket{
			Winning: compile.Split(splited[0], -1)[1:],
			Numbers: compile.Split(splited[1], -1)[1:],
		})
	}

	for _, t := range tickets {
		p := t.count()
		if p >= 0 {
			sum += math.Pow(2.0, float64(p))
		}
	}

	fmt.Println(sum)
}

func (t ticket) count() int {
	c := -1
	for _, w := range t.Winning {
		for _, n := range t.Numbers {
			if w == n {
				c++
			}
		}
	}
	return c
}
