package day10

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strings"
)

type pipe2 struct {
	Type       string
	Visited    bool
	Coords     [2]int
	Neighbours []*pipe2
	Loop       bool
	Enclosed   bool
}

func Part2() {
	field := 0
	scanner := utils.ReadFile("day10")
	var scan [][]*pipe2
	var starting *pipe2

	for i := 0; scanner.Scan(); i++ {
		var row []*pipe2
		s := strings.Split(scanner.Text(), "")

		for k, c := range s {
			p := pipe2{Coords: [2]int{i, k}, Type: c, Visited: false, Enclosed: false}
			row = append(row, &p)
			if c == "S" {
				starting = &p
			}
		}

		scan = append(scan, row)
	}

	for _, r := range scan {
		for _, p := range r {
			switch p.Type {
			case "|":
				p.top(&scan)
				p.bottom(&scan)
				break
			case "-":
				p.left(&scan)
				p.right(&scan)
				break
			case "L":
				p.top(&scan)
				p.right(&scan)
				break
			case "J":
				p.top(&scan)
				p.left(&scan)
				break
			case "7":
				p.bottom(&scan)
				p.left(&scan)
				break
			case "F":
				p.bottom(&scan)
				p.right(&scan)
				break
			case "S":
				p.top(&scan)
				p.left(&scan)
				p.bottom(&scan)
				p.right(&scan)
				break
			default:
				break
			}
		}
	}

	rLF, _ := regexp.Compile("(LF)+")
	rJ7, _ := regexp.Compile("(J7)+")
	rFL, _ := regexp.Compile("(FL)+")
	r7J, _ := regexp.Compile("(7J)+")
	r7F, _ := regexp.Compile("(7F)+")
	rJL, _ := regexp.Compile("(JL)+")
	rF7, _ := regexp.Compile("(F7)+")
	rLJ, _ := regexp.Compile("(LJ)+")
	two, _ := regexp.Compile("[7LFJ]{2}")

	starting.DFS()
	for _, r := range scan {
		for _, p := range r {
			right := ""
			left := ""
			top := ""
			bottom := ""
			if !p.Loop {
				for i := p.Coords[1] + 1; i < len(r); i++ {
					if r[i].Loop &&
						r[i].Type != "-" {
						right += r[i].Type
					}
				}
				right = rF7.ReplaceAllString(right, "")
				right = rLJ.ReplaceAllString(right, "")
				right = two.ReplaceAllString(right, "1")

				for i := p.Coords[1] - 1; i >= 0; i-- {
					if r[i].Loop &&
						r[i].Type != "-" {
						left += r[i].Type
					}
				}
				left = r7F.ReplaceAllString(left, "")
				left = rJL.ReplaceAllString(left, "")
				left = two.ReplaceAllString(left, "1")
				for i := p.Coords[0] + 1; i < len(scan); i++ {
					if scan[i][p.Coords[1]].Loop &&
						scan[i][p.Coords[1]].Type != "|" {
						bottom += scan[i][p.Coords[1]].Type
					}
				}
				bottom = rFL.ReplaceAllString(bottom, "")
				bottom = r7J.ReplaceAllString(bottom, "")
				bottom = two.ReplaceAllString(bottom, "1")
				for i := p.Coords[0] - 1; i >= 0; i-- {
					if scan[i][p.Coords[1]].Loop &&
						scan[i][p.Coords[1]].Type != "|" {
						top += scan[i][p.Coords[1]].Type
					}
				}
				top = rLF.ReplaceAllString(top, "")
				top = rJ7.ReplaceAllString(top, "")
				top = two.ReplaceAllString(top, "1")
			}

			if !p.Loop &&
				len(top) > 0 && len(top)%2 == 1 &&
				len(bottom) > 0 && len(bottom)%2 == 1 &&
				len(left) > 0 && len(left)%2 == 1 &&
				len(right) > 0 && len(right)%2 == 1 {
				fmt.Println(p, top, len(top), bottom, len(bottom), left, len(left), right, len(right))
				field++
			}
		}
	}

	fmt.Println(field)
}

func (p *pipe2) top(scan *[][]*pipe2) {
	if p.Coords[0] != 0 {
		if (*scan)[p.Coords[0]-1][p.Coords[1]].Type == "|" ||
			(*scan)[p.Coords[0]-1][p.Coords[1]].Type == "7" ||
			(*scan)[p.Coords[0]-1][p.Coords[1]].Type == "F" {
			p.Neighbours = append(p.Neighbours, (*scan)[p.Coords[0]-1][p.Coords[1]])
		}
	}
}

func (p *pipe2) bottom(scan *[][]*pipe2) {
	if p.Coords[0] != len(*scan)-1 {
		if (*scan)[p.Coords[0]+1][p.Coords[1]].Type == "|" ||
			(*scan)[p.Coords[0]+1][p.Coords[1]].Type == "L" ||
			(*scan)[p.Coords[0]+1][p.Coords[1]].Type == "J" {
			p.Neighbours = append(p.Neighbours, (*scan)[p.Coords[0]+1][p.Coords[1]])
		}
	}
}

func (p *pipe2) left(scan *[][]*pipe2) {
	if p.Coords[1] != 0 {
		if (*scan)[p.Coords[0]][p.Coords[1]-1].Type == "-" ||
			(*scan)[p.Coords[0]][p.Coords[1]-1].Type == "L" ||
			(*scan)[p.Coords[0]][p.Coords[1]-1].Type == "F" {
			p.Neighbours = append(p.Neighbours, (*scan)[p.Coords[0]][p.Coords[1]-1])
		}
	}
}
func (p *pipe2) right(scan *[][]*pipe2) {
	if p.Coords[1] != len((*scan)[p.Coords[0]])-1 {
		if (*scan)[p.Coords[0]][p.Coords[1]+1].Type == "-" ||
			(*scan)[p.Coords[0]][p.Coords[1]+1].Type == "7" ||
			(*scan)[p.Coords[0]][p.Coords[1]+1].Type == "J" {
			p.Neighbours = append(p.Neighbours, (*scan)[p.Coords[0]][p.Coords[1]+1])
		}
	}
}

type stack2 []*pipe2

func (s *stack2) Pop() *pipe2 {
	if s.IsEmpty() {
		return nil
	} else {
		i := len(*s) - 1
		val := (*s)[i]
		*s = (*s)[:i]

		return val
	}
}

func (s *stack2) Push(value *pipe2) {
	*s = append(*s, value)
}

func (s *stack2) IsEmpty() bool {
	return len(*s) == 0
}

func (p *pipe2) DFS() {
	var s stack2
	s.Push(p)
	p.Visited = true
	p.Loop = true
	for !s.IsEmpty() {
		t := s.Pop()
		for _, u := range t.Neighbours {
			if !u.Visited {
				u.Loop = true
				u.Visited = true
				s.Push(u)
			}
		}
	}
}
