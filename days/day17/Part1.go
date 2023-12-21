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
	Heat     int
	Routes   [4]*block
	Streak   int
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
			b := block{Heat: atoi, Distance: math.MaxInt32}
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
	//stop := route[len(route)-1][len(route[0])-1]

	fmt.Println(start)
}

// 4d djikstra x,y,dir,timesmoved holding int least path
