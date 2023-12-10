package day10

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

type pipe struct {
	Type       string
	Visited    bool
	Coords     [2]int
	Neighbours []*pipe
	Prev       *pipe
	Distance   int
}

func Part1() {
	scanner := utils.ReadFile("day10")
	var scan [][]*pipe
	var starting *pipe

	for i := 0; scanner.Scan(); i++ {
		var row []*pipe
		s := strings.Split(scanner.Text(), "")

		for k, c := range s {
			p := pipe{Coords: [2]int{i, k}, Type: c, Visited: false}
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

	starting.BFS()
	maxD := starting
	for _, r := range scan {
		for _, p := range r {
			p.Visited = false
			if p.Distance > maxD.Distance {
				maxD = p
			}
		}
	}

	starting.DFS()
	fmt.Println(maxD.Distance)
}

func (p *pipe) top(scan *[][]*pipe) {
	if p.Coords[0] != 0 {
		if (*scan)[p.Coords[0]-1][p.Coords[1]].Type == "|" ||
			(*scan)[p.Coords[0]-1][p.Coords[1]].Type == "7" ||
			(*scan)[p.Coords[0]-1][p.Coords[1]].Type == "F" {
			p.Neighbours = append(p.Neighbours, (*scan)[p.Coords[0]-1][p.Coords[1]])
		}
	}
}

func (p *pipe) bottom(scan *[][]*pipe) {
	if p.Coords[0] != len(*scan)-1 {
		if (*scan)[p.Coords[0]+1][p.Coords[1]].Type == "|" ||
			(*scan)[p.Coords[0]+1][p.Coords[1]].Type == "L" ||
			(*scan)[p.Coords[0]+1][p.Coords[1]].Type == "J" {
			p.Neighbours = append(p.Neighbours, (*scan)[p.Coords[0]+1][p.Coords[1]])
		}
	}
}

func (p *pipe) left(scan *[][]*pipe) {
	if p.Coords[1] != 0 {
		if (*scan)[p.Coords[0]][p.Coords[1]-1].Type == "-" ||
			(*scan)[p.Coords[0]][p.Coords[1]-1].Type == "L" ||
			(*scan)[p.Coords[0]][p.Coords[1]-1].Type == "F" {
			p.Neighbours = append(p.Neighbours, (*scan)[p.Coords[0]][p.Coords[1]-1])
		}
	}
}
func (p *pipe) right(scan *[][]*pipe) {
	if p.Coords[1] != len((*scan)[p.Coords[0]])-1 {
		if (*scan)[p.Coords[0]][p.Coords[1]+1].Type == "-" ||
			(*scan)[p.Coords[0]][p.Coords[1]+1].Type == "7" ||
			(*scan)[p.Coords[0]][p.Coords[1]+1].Type == "J" {
			p.Neighbours = append(p.Neighbours, (*scan)[p.Coords[0]][p.Coords[1]+1])
		}
	}
}

type stack []*pipe

func (s *stack) Pop() *pipe {
	if s.IsEmpty() {
		return nil
	} else {
		i := len(*s) - 1
		val := (*s)[i]
		*s = (*s)[:i]

		return val
	}
}

func (s *stack) Push(value *pipe) {
	*s = append(*s, value)
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (p *pipe) DFS() {
	count := 1
	var s stack
	s.Push(p)
	p.Visited = true
	for !s.IsEmpty() {
		t := s.Pop()
		for _, u := range t.Neighbours {
			if !u.Visited {
				u.Prev = t
				u.Distance = count
				u.Visited = true
				s.Push(u)
			}
		}
		count++
	}
}

type queue []*pipe

func (q *queue) Dequeue() *pipe {
	val := (*q)[0]
	if !q.IsEmpty() {
		*q = (*q)[1:]
	}

	return val
}

func (q *queue) Enqueue(value *pipe) {
	*q = append(*q, value)
}

func (q *queue) IsEmpty() bool {
	return len(*q) == 0
}

func (p *pipe) BFS() {
	count := 1
	var q queue
	q.Enqueue(p)
	for !q.IsEmpty() {
		t := q.Dequeue()
		for _, u := range t.Neighbours {
			if !u.Visited {
				u.Prev = t
				u.Distance = count
				u.Visited = true
				q.Enqueue(u)
			}
		}
		count++
	}
}
