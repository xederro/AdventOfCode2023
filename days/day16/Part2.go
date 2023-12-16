package day16

import (
	"adventOfCode/utils"
	"fmt"
	"strings"
)

func Part2() {
	scanner := utils.ReadFile("day16")
	var wall [][]tile
	var res int

	for scanner.Scan() {
		var row []tile
		s := scanner.Text()
		split := strings.Split(s, "")
		for _, str := range split {
			var c int
			switch str {
			default:
				c = DOT
				break
			case "|":
				c = PIPE
				break
			case "-":
				c = MINUS
				break
			case "\\":
				c = BACKSLASH
				break
			case "/":
				c = SLASH
				break
			}
			row = append(row, tile{Type: c})
		}
		wall = append(wall, row)
	}

	for i := 0; i < len(wall); i++ {
		tempWall := copyWall(wall)
		tmp := 0
		lightBeam(&tempWall, -1, i, RIGHT)
		for _, tiles := range tempWall {
			for _, t := range tiles {
				if t.Energized {
					tmp++
				}
			}
		}
		res = max(res, tmp)

		tempWall = copyWall(wall)
		tmp = 0
		lightBeam(&tempWall, len(wall[0]), i, LEFT)
		for _, tiles := range tempWall {
			for _, t := range tiles {
				if t.Energized {
					tmp++
				}
			}
		}
		res = max(res, tmp)
	}

	for i := 0; i < len(wall[0]); i++ {
		tempWall := copyWall(wall)
		tmp := 0
		lightBeam(&tempWall, i, -1, BOTTOM)
		for _, tiles := range tempWall {
			for _, t := range tiles {
				if t.Energized {
					tmp++
				}
			}
		}
		res = max(res, tmp)

		tempWall = copyWall(wall)
		tmp = 0
		lightBeam(&tempWall, i, len(wall), TOP)
		for _, tiles := range tempWall {
			for _, t := range tiles {
				if t.Energized {
					tmp++
				}
			}
		}
		res = max(res, tmp)
	}

	fmt.Println(res)
}

func copyWall(wall [][]tile) [][]tile {
	var newWall [][]tile
	for _, tiles := range wall {
		var newRow []tile
		for _, t := range tiles {
			newRow = append(newRow, tile{Type: t.Type})
		}
		newWall = append(newWall, newRow)
	}
	return newWall
}
