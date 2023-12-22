package day17

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type block struct {
	Heat     int
	Routes   []*block
	Distance int
	Visited  bool
}

func Part1() {
	scanner := utils.ReadFile("day17")
	//x, y, dir, streak
	var route [][][4][3]*block

	var pq utils.PriorityQueue[block]
	pq.SetComparator(func(a, b *block) bool {
		if a.Distance < b.Distance {
			return false
		}
		return true
	})

	for scanner.Scan() {
		s := scanner.Text()
		var row [][4][3]*block
		split := strings.Split(s, "")
		for _, str := range split {
			atoi, err := strconv.Atoi(str)
			if err != nil {
				return
			}
			var cell [4][3]*block
			for i := 0; i < 4; i++ {
				for j := 0; j < 3; j++ {
					b := block{Heat: atoi, Distance: math.MaxInt, Visited: false}
					cell[i][j] = &b
					pq.Enqueue(&b)
				}
			}
			row = append(row, cell)
		}
		route = append(route, row)
	}

	maxy, maxx := len(route)-1, len(route[0])-1

	for y, row := range route {
		for x, cell := range row {
			for dir, param := range cell {
				for streak, b := range param {
					var r []*block

					for newDir := 0; newDir < 4; newDir++ {
						curry := y
						currx := x

						if y > 0 && newDir == 0 {
							curry--
						}
						if y < maxy && newDir == 1 {
							curry++
						}
						if x > 0 && newDir == 2 {
							currx--
						}
						if x < maxx && newDir == 3 {
							currx++
						}
						//fmt.Println(curry, currx, y, x, dir)
						for i := 0; i < 4; i++ {
							if i == dir {
								if streak != 2 {
									r = append(r, route[curry][currx][i][streak+1])
								}
								continue
							}
							r = append(r, route[curry][currx][i][0])
						}
					}

					b.Routes = append(b.Routes, r...)
				}
			}
		}
	}

	route[0][0][0][0].Distance = 0
	route[0][0][1][0].Distance = 0
	route[0][0][2][0].Distance = 0
	route[0][0][3][0].Distance = 0

	stop := []*block{
		route[maxy][maxx][0][0],
		route[maxy][maxx][0][1],
		route[maxy][maxx][0][2],
		route[maxy][maxx][1][0],
		route[maxy][maxx][1][1],
		route[maxy][maxx][1][2],
		route[maxy][maxx][2][0],
		route[maxy][maxx][2][1],
		route[maxy][maxx][2][2],
		route[maxy][maxx][3][0],
		route[maxy][maxx][3][1],
		route[maxy][maxx][3][2],
	}

	djikstra(pq)

	minimal := math.MaxInt
	for _, tmp := range stop {
		if tmp.Distance < minimal {
			minimal = tmp.Distance
		}
	}

	fmt.Println(minimal)
}

// 4d djikstra x,y,dir,timesmoved holding int least path
func djikstra(pq utils.PriorityQueue[block]) {
	for !pq.IsEmpty() {
		u := pq.Dequeue()
		u.Visited = true

		for _, n := range u.Routes {
			if !n.Visited {
				tmp := u.Distance + n.Heat
				if tmp < n.Distance {
					n.Distance = tmp
				}
			}
		}
	}
}
