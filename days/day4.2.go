package days

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strings"
)

type ticket2 struct {
	Winning []string
	Numbers []string
	Count   int
}

func Day4z2() {
	scanner := utils.ReadFile("day4")
	var tickets []ticket2
	sum := 0

	compile, err := regexp.Compile("[ ]+")
	if err != nil {
		return
	}

	for scanner.Scan() {
		s := scanner.Text()
		splited := strings.Split(strings.Split(s, ":")[1], "|")
		tickets = append(tickets, ticket2{
			Winning: compile.Split(splited[0], -1)[1:],
			Numbers: compile.Split(splited[1], -1)[1:],
			Count:   1,
		})
	}

	for i := 0; i < len(tickets); i++ {
		for j := tickets[i].count(); j >= 0; j-- {
			tickets[i+j+1].Count += tickets[i].Count
		}
		sum += tickets[i].Count
	}

	fmt.Println(sum)
}

func (t ticket2) count() int {
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
