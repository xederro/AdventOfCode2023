package day17

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type block2 struct {
	Heat     int
	Routes   []*block2
	Prev     *block2
	Distance int
	Visited  bool
	Coords   [4]int
}

func Part2() {
	scanner := utils.ReadFile("day17")
	var route [][][4][10]*block2

	var pq utils.PriorityQueue[block2]
	pq.SetComparator(func(a, b *block2) bool {
		if a.Distance < b.Distance {
			return false
		}
		return true
	})

	for scanner.Scan() {
		s := scanner.Text()
		var row [][4][10]*block2
		split := strings.Split(s, "")
		for _, str := range split {
			atoi, err := strconv.Atoi(str)
			if err != nil {
				return
			}
			var cell [4][10]*block2
			for i := 0; i < 4; i++ {
				for j := 0; j < 10; j++ {
					b := block2{Heat: atoi, Distance: math.MaxInt, Visited: false}
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
					b.Coords = [4]int{y, x, dir, streak}
					if streak < 3 {
						curry := y
						currx := x

						if dir == 0 {
							if y > 0 {
								curry--
							} else {
								continue
							}
						}
						if dir == 1 {
							if y < maxy {
								curry++
							} else {
								continue
							}
						}
						if dir == 2 {
							if x > 0 {
								currx--
							} else {
								continue
							}
						}
						if dir == 3 {
							if x < maxx {
								currx++
							} else {
								continue
							}
						}

						b.Routes = append(b.Routes, route[curry][currx][dir][streak+1])
					} else {
						for newDir := 0; newDir < 4; newDir++ {
							curry := y
							currx := x

							if newDir == 0 {
								if y > 0 {
									curry--
								} else {
									continue
								}
							}
							if newDir == 1 {
								if y < maxy {
									curry++
								} else {
									continue
								}
							}
							if newDir == 2 {
								if x > 0 {
									currx--
								} else {
									continue
								}
							}
							if newDir == 3 {
								if x < maxx {
									currx++
								} else {
									continue
								}
							}

							if newDir == dir {
								if streak == 9 {
									continue
								}
								b.Routes = append(b.Routes, route[curry][currx][newDir][streak+1])
							} else if newDir == 0 && dir == 1 || newDir == 1 && dir == 0 || newDir == 2 && dir == 3 || newDir == 3 && dir == 2 {
								continue
							} else {
								b.Routes = append(b.Routes, route[curry][currx][newDir][0])
							}
						}
					}
				}
			}
		}
	}

	route[0][0][0][0].Distance = 0
	route[0][0][1][0].Distance = 0
	route[0][0][2][0].Distance = 0
	route[0][0][3][0].Distance = 0

	var stop []*block2

	for i := 0; i < 4; i++ {
		for j := 3; j < 10; j++ {
			stop = append(stop, route[maxy][maxx][i][j])
		}
	}

	djikstra2(pq)

	minimal := math.MaxInt
	for _, tmp := range stop {
		if tmp.Distance < minimal {
			minimal = tmp.Distance
		}
	}

	fmt.Println(minimal)
}

func djikstra2(pq utils.PriorityQueue[block2]) {
	for !pq.IsEmpty() {
		u := pq.Dequeue()
		u.Visited = true

		for _, n := range u.Routes {
			if !n.Visited {
				tmp := u.Distance + n.Heat
				if tmp < n.Distance {
					n.Distance = tmp
					n.Prev = u
				}
			}
		}
	}
}
