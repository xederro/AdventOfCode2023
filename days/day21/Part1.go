package day21

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

type plot struct {
	Visitable bool
	Even      bool
}

func Part1() {
	scanner := utils.ReadFile("day21")
	count := 0
	mapa := map[[2]int]*plot{}
	possible := map[[2]int]*plot{}

	j := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		for i, p := range s {
			if p == "#" {
				continue
			}

			plt := plot{}

			if p == "S" {
				plt.Even = true
				plt.Visitable = true
				possible[[2]int{j, i}] = &plt
			}

			mapa[[2]int{j, i}] = &plt
		}
		j++
	}

	n := 64
	for i := 1; i <= n; i++ {
		temp := map[[2]int]*plot{}
		e := i%2 == 0
		for ints, p := range possible {
			cur := p.visit(mapa, ints, e)
			for c, p := range cur {
				temp[c] = p
			}
		}

		possible = temp
	}

	for _, p := range mapa {
		if p.Visitable && p.Even {
			count++
		}
	}

	fmt.Println(count)
}

func (p *plot) visit(m map[[2]int]*plot, c [2]int, even bool) map[[2]int]*plot {
	cur := map[[2]int]*plot{}

	if v, ok := m[[2]int{c[0] + 1, c[1]}]; ok && !v.Visitable {
		v.Visitable = true
		v.Even = even
		cur[[2]int{c[0] + 1, c[1]}] = v
	}

	if v, ok := m[[2]int{c[0] - 1, c[1]}]; ok && !v.Visitable {
		v.Visitable = true
		v.Even = even
		cur[[2]int{c[0] - 1, c[1]}] = v
	}

	if v, ok := m[[2]int{c[0], c[1] + 1}]; ok && !v.Visitable {
		v.Visitable = true
		v.Even = even
		cur[[2]int{c[0], c[1] + 1}] = v
	}

	if v, ok := m[[2]int{c[0], c[1] - 1}]; ok && !v.Visitable {
		v.Visitable = true
		v.Even = even
		cur[[2]int{c[0], c[1] - 1}] = v
	}

	return cur
}
