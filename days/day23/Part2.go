package day23

import (
	"adventOfCode/utils"
	"fmt"
)

func Part2() {
	var mapa [][]*path
	var start *path
	var stop *path
	scanner := utils.ReadFile("day23")
	for i := 0; scanner.Scan(); i++ {
		var row []*path
		s := fmt.Sprintf(scanner.Text(), "")
		for j, v := range s {
			t := -1
			switch v {
			case '.':
				t = PATH
				break
			case 'v':
				t = BSLOPE
				break
			case '>':
				t = RSLOPE
				break
			case '^':
				t = TSLOPE
				break
			case '<':
				t = LSLOPE
				break
			}

			var p *path

			if t != -1 {
				p = &path{
					Type:   t,
					Coords: [2]int{i, j},
				}
			}

			if i == 0 && p != nil {
				start = p
			}

			row = append(row, p)
		}
		mapa = append(mapa, row)
	}

	for _, row := range mapa {
		for _, v := range row {
			if v != nil {
				if v.Coords[0] == len(mapa)-1 {
					stop = v
				}
				v.connect2(mapa)
			}
		}
	}
	fmt.Println(start.DFS(start, stop, 0) + 1)
}

func (p *path) connect2(mapa [][]*path) {
	if p.Coords[0] == 0 {
		p.Neighbours = append(p.Neighbours, mapa[1][p.Coords[1]])
		return
	}
	if p.Coords[0] == len(mapa)-1 {
		p.Neighbours = append(p.Neighbours, mapa[p.Coords[0]-1][p.Coords[1]])
		return
	}

	if mapa[p.Coords[0]-1][p.Coords[1]] != nil {
		p.Neighbours = append(p.Neighbours, mapa[p.Coords[0]-1][p.Coords[1]])
	}
	if mapa[p.Coords[0]+1][p.Coords[1]] != nil {
		p.Neighbours = append(p.Neighbours, mapa[p.Coords[0]+1][p.Coords[1]])
	}
	if mapa[p.Coords[0]][p.Coords[1]-1] != nil {
		p.Neighbours = append(p.Neighbours, mapa[p.Coords[0]][p.Coords[1]-1])
	}
	if mapa[p.Coords[0]][p.Coords[1]+1] != nil {
		p.Neighbours = append(p.Neighbours, mapa[p.Coords[0]][p.Coords[1]+1])
	}
}
