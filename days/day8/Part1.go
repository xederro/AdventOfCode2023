package day8

import (
	"adventOfCode/utils"
	"fmt"
)

var (
	current int                  = 0
	length  int                  = 0
	count   int                  = 0
	tree    map[string][2]string = map[string][2]string{}
	steps   string
)

func Part1() {
	scanner := utils.ReadFile("day8")

	//compile, err := regexp.Compile(" +")
	//if err != nil {
	//	return
	//}

	if scanner.Scan() {
		steps = scanner.Text()
		length = len(steps)
		scanner.Scan()
	}

	for scanner.Scan() {
		s := scanner.Text()
		tree[s[0:3]] = [2]string{s[7:10], s[12:15]}

	}

	s := "AAA"
	for s != "ZZZ" {
		s = turn(s)
	}

	fmt.Println(count)
}

func turn(step string) string {
	count++
	var s string
	if steps[current] == 'R' {
		s = tree[step][1]
	} else {
		s = tree[step][0]
	}
	current = (current + 1) % length
	return s
}
