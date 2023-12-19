package day19

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type part struct {
	X   int
	M   int
	A   int
	S   int
	Out bool
}

type instruction struct {
	Rules []rule
}

type rule struct {
	Part  string
	Sign  string
	Count int
	Dest  string
}

func Part1() {
	scanner := utils.ReadFile("day19")
	instructions := map[string]instruction{}
	parts := []part{}
	compile, err := regexp.Compile("[{}]")
	if err != nil {
		return
	}

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}

		i := instruction{}

		split := compile.Split(s, -1)
		ins := strings.Split(split[1], ",")
		for k, v := range ins {
			if k == len(ins)-1 {
				i.Rules = append(i.Rules, rule{Dest: v})
			} else {
				spl := strings.Split(v, ":")
				c, err := strconv.Atoi(spl[0][2:])
				if err != nil {
					fmt.Println(err, 2)
					return
				}
				i.Rules = append(i.Rules, rule{Dest: spl[1], Count: c, Sign: spl[0][1:2], Part: spl[0][0:1]})
			}
		}

		instructions[split[0]] = i
	}

	for scanner.Scan() {
		s := scanner.Text()
		split := compile.Split(s, -1)
		p := strings.Split(split[1], ",")

		prt := part{}

		for _, v := range p {
			atoi, err := strconv.Atoi(v[2:])
			if err != nil {
				fmt.Println(err, 1)
				return
			}
			switch v[0] {
			case 'x':
				prt.X = atoi
				break
			case 'm':
				prt.M = atoi
				break
			case 'a':
				prt.A = atoi
				break
			case 's':
				prt.S = atoi
				break
			}
		}

		parts = append(parts, prt)
	}

	count := 0

	for _, p := range parts {
		p.process(instructions)
		count += p.count()
	}

	fmt.Println(count)
}

func (p part) count() int {
	if !p.Out {
		return p.M + p.A + p.X + p.S
	}
	return 0
}

func (p *part) process(ins map[string]instruction) {
	s := "in"
	for s != "R" && s != "A" && s != "" {
		if i, ok := ins[s]; ok {
			s = i.through(*p)
		}
	}
	if s == "R" {
		p.Out = true
	}
}

func (i instruction) through(p part) string {
	for _, r := range i.Rules {
		if r.eval(p) {
			return r.Dest
		}
	}
	return ""
}

func (r rule) eval(p part) bool {
	switch r.Part {
	case "":
		return true
	case "x":
		switch r.Sign {
		case "<":
			return p.X < r.Count
		case ">":
			return p.X > r.Count
		case "=":
			return p.X == r.Count
		}
		break
	case "m":
		switch r.Sign {
		case "<":
			return p.M < r.Count
		case ">":
			return p.M > r.Count
		case "=":
			return p.M == r.Count
		}
		break
	case "a":
		switch r.Sign {
		case "<":
			return p.A < r.Count
		case ">":
			return p.A > r.Count
		case "=":
			return p.A == r.Count
		}
		break
	case "s":
		switch r.Sign {
		case "<":
			return p.S < r.Count
		case ">":
			return p.S > r.Count
		case "=":
			return p.S == r.Count
		}
		break
	}

	return false
}
