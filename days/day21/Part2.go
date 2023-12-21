package day21

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

type plot2 struct {
	Visitable map[[2]int]bool
}

func Part2() {
	scanner := utils.ReadFile("day21")
	a := int64(0)
	b := int64(0)
	c := int64(0)
	mapa := map[[2]int]*plot2{}
	possible := map[[4]int]*plot2{}

	maxx := 0
	maxy := 0
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "")
		for i, p := range s {
			if i > maxx {
				maxx = i
			}
			if p == "#" {
				continue
			}

			plt := plot2{
				Visitable: make(map[[2]int]bool),
			}

			if p == "S" {
				plt.Visitable[[2]int{0, 0}] = true
				possible[[4]int{maxy, i, 0, 0}] = &plt
			}

			mapa[[2]int{maxy, i}] = &plt
		}
		maxy++
	}

	maxx++

	for i := 1; i <= 65+2*131; i++ {
		temp := map[[4]int]*plot2{}
		for ints, p := range possible {
			cur := p.visit(mapa, ints, maxx, maxy)
			for c, p := range cur {
				temp[c] = p
			}
		}

		possible = temp
		if i == 65 {
			for _, p := range mapa {
				for _, p1 := range p.Visitable {
					if p1 {
						a++
					}
				}
			}
		}
		if i == 65+131 {
			for _, p := range mapa {
				for _, p1 := range p.Visitable {
					if p1 {
						b++
					}
				}
			}
		}
	}

	for _, p := range mapa {
		for _, p1 := range p.Visitable {
			if p1 {
				c++
			}
		}
	}

	//https://www.wolframalpha.com/input?i=quadratic+fit+calculator&assumption=%7B%22F%22%2C+%22QuadraticFitCalculator%22%2C+%22data3x%22%7D+-%3E%22%7B0%2C+1%2C+2%7D%22&assumption=%7B%22F%22%2C+%22QuadraticFitCalculator%22%2C+%22data3y%22%7D+-%3E%22%7B3847%2C+34165%2C+94697%7D%22
	x := int64(26501365 / 131)
	fmt.Println(((c-a-(2*(b-a)))/2)*x*x + (b-a-((c-a-(2*(b-a)))/2))*x + a)
}

func (p *plot2) visit(m map[[2]int]*plot2, c [4]int, maxx, maxy int) map[[4]int]*plot2 {
	cur := map[[4]int]*plot2{}
	p.Visitable[[2]int{c[2], c[3]}] = false

	bottom := [4]int{(maxy + c[0] + 1) % maxy, c[1], c[2], c[3]}
	if c[0]+1 == maxy {
		bottom[2]++
	}
	if b, bok := m[[2]int{bottom[0], bottom[1]}]; bok {
		if v, ok := b.Visitable[[2]int{bottom[2], bottom[3]}]; !ok || !v {
			b.Visitable[[2]int{bottom[2], bottom[3]}] = true
			cur[bottom] = b
		}
	}

	top := [4]int{(maxy + c[0] - 1) % maxy, c[1], c[2], c[3]}
	if c[0]-1 < 0 {
		top[2]--
	}
	if t, tok := m[[2]int{top[0], top[1]}]; tok {
		if v, ok := t.Visitable[[2]int{top[2], top[3]}]; !ok || !v {
			t.Visitable[[2]int{top[2], top[3]}] = true
			cur[top] = t
		}
	}

	left := [4]int{c[0], (maxx + c[1] - 1) % maxx, c[2], c[3]}
	if c[1]-1 < 0 {
		left[3]--
	}
	if l, lok := m[[2]int{left[0], left[1]}]; lok {
		if v, ok := l.Visitable[[2]int{left[2], left[3]}]; !ok || !v {
			l.Visitable[[2]int{left[2], left[3]}] = true
			cur[left] = l
		}
	}

	right := [4]int{c[0], (maxx + c[1] + 1) % maxx, c[2], c[3]}
	if c[1]+1 == maxx {
		right[3]++
	}
	if r, rok := m[[2]int{right[0], right[1]}]; rok {
		if v, ok := r.Visitable[[2]int{right[2], right[3]}]; !ok || !v {
			r.Visitable[[2]int{right[2], right[3]}] = true
			cur[right] = r
		}
	}

	return cur
}
