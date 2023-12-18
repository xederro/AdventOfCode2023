package day17

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	TOP = iota
	LEFT
	BOTTOM
	RIGHT
)

type block struct {
	Visited  [4]bool
	Heat     int
	Routes   [4]*block
	Streak   [4]int
	Distance int
}

func Part1() {
	var nodes []*block
	scanner := utils.ReadFile("day17")
	var route [][]*block

	for scanner.Scan() {
		var row []*block
		s := scanner.Text()
		split := strings.Split(s, "")
		for _, str := range split {
			atoi, err := strconv.Atoi(str)
			if err != nil {
				return
			}
			b := block{Heat: atoi, Distance: math.MaxInt32, Streak: [4]int{0, 0, 0, 0}}
			row = append(row, &b)
			nodes = append(nodes, &b)
		}
		route = append(route, row)
	}

	for i := 0; i < len(route); i++ {
		for j := 0; j < len(route[i]); j++ {
			var routes [4]*block
			if i != 0 {
				routes[TOP] = route[i-1][j]
			}
			if j != 0 {
				routes[LEFT] = route[i][j-1]
			}
			if i != len(route)-1 {
				routes[BOTTOM] = route[i+1][j]
			}
			if j != len(route[i])-1 {
				routes[RIGHT] = route[i][j+1]
			}
			route[i][j].Routes = routes
		}
	}

	start := route[0][0]
	stop := route[len(route)-1][len(route[0])-1]

	fmt.Println(start.dijkstra(stop, TOP, 0))
	for _, blocks := range route {
		for _, b := range blocks {
			fmt.Print(b.Distance, "\t")
		}
		fmt.Println()
	}
}

func (b *block) dijkstra(end *block, dir int, streak int) int {
	if b == end {
		return 0
	}

	if b == nil || b.Visited[dir] || streak > 2 {
		return math.MaxInt32
	}

	b.Visited[dir] = true

	minDistance := math.MaxInt32
	for i, neighbor := range b.Routes {
		dist := math.MaxInt32
		if b.Routes[i] != nil && !b.Routes[i].Visited[i] {
			if dir == i {
				if streak+1 < 2 {
					dist = neighbor.dijkstra(end, i, streak+1) + neighbor.Heat
				}
			} else {
				dist = neighbor.dijkstra(end, i, 0) + neighbor.Heat
			}
		}

		if dist < minDistance {
			minDistance = dist
		}
	}

	b.Visited[dir] = true

	if minDistance < b.Distance {
		b.Distance = minDistance
	}
	return b.Distance
}
