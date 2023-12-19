package day19

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type instruction2 struct {
	Rules []rule2
}

type rule2 struct {
	Part    string
	PartNum int
	Sign    string
	Count   int
	Dest    string
}

const (
	x = iota
	m
	a
	s
)

func Part2() {
	scanner := utils.ReadFile("day19")
	instructions := map[string]instruction2{}
	compile, err := regexp.Compile("[{}]")
	if err != nil {
		return
	}

	for scanner.Scan() {
		o := scanner.Text()

		if o == "" {
			break
		}

		i := instruction2{}

		split := compile.Split(o, -1)
		ins := strings.Split(split[1], ",")
		for k, v := range ins {
			if k == len(ins)-1 {
				i.Rules = append(i.Rules, rule2{Dest: v})
			} else {
				spl := strings.Split(v, ":")
				c, err := strconv.Atoi(spl[0][2:])
				if err != nil {
					fmt.Println(err, 2)
					return
				}

				num := 0
				switch spl[0][0] {
				case 'x':
					num = x
					break
				case 'm':
					num = m
					break
				case 'a':
					num = a
					break
				case 's':
					num = s
					break
				}
				i.Rules = append(i.Rules, rule2{Dest: spl[1], Count: c, Sign: spl[0][1:2], Part: spl[0][0:1], PartNum: num})
			}
		}

		instructions[split[0]] = i
	}

	fmt.Println(count("in", instructions, [4][2]int{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}))
}

func count(ins string, instructions map[string]instruction2, ranges [4][2]int) int {
	c := 0
	switch ins {
	case "A":
		return (ranges[0][1] - ranges[0][0] + 1) * (ranges[1][1] - ranges[1][0] + 1) * (ranges[2][1] - ranges[2][0] + 1) * (ranges[3][1] - ranges[3][0] + 1)
	case "R":
		return 0
	default:
		for _, r := range instructions[ins].Rules {
			tmp := copyRange(ranges)
			switch r.Sign {
			case ">":
				if ranges[r.PartNum][1] > r.Count {
					if ranges[r.PartNum][0] < r.Count {
						tmp[r.PartNum][0] = r.Count + 1
						ranges[r.PartNum][1] = r.Count
					}
					c += count(r.Dest, instructions, tmp)
				}
				break
			case "<":
				if ranges[r.PartNum][0] < r.Count {
					if ranges[r.PartNum][1] > r.Count {
						tmp[r.PartNum][1] = r.Count - 1
						ranges[r.PartNum][0] = r.Count
					}
					c += count(r.Dest, instructions, tmp)
				}
				break
			default:
				c += count(r.Dest, instructions, ranges)
				break
			}
		}
		return c
	}
}

func copyRange(ranges [4][2]int) [4][2]int {
	var r [4][2]int
	for k, ints := range ranges {
		r[k] = [2]int{ints[0], ints[1]}
	}
	return r
}
