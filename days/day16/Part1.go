package day16

import (
	"adventOfCode/utils"
	"fmt"
	"slices"
	"strings"
)

const (
	TOP = iota
	LEFT
	BOTTOM
	RIGHT
	BACKSLASH
	SLASH
	MINUS
	PIPE
	DOT
)

type tile struct {
	Energized bool
	Type      int
	Checked   []int
}

func Part1() {
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

	lightBeam(&wall, -1, 0, RIGHT)

	for _, tiles := range wall {
		for _, t := range tiles {
			if t.Energized {
				res++
			}
		}
	}

	fmt.Println(res)
}

func lightBeam(wall *[][]tile, x int, y int, direction int) {
	if y != -1 && y != len(*wall) && x != -1 && x != len((*wall)[0]) {
		(*wall)[y][x].Energized = true
	}
	switch direction {
	case TOP:
		if 0 == y {
			break
		}
		if slices.Contains((*wall)[y-1][x].Checked, TOP) {
			break
		}
		(*wall)[y-1][x].Checked = append((*wall)[y-1][x].Checked, TOP)
		switch (*wall)[y-1][x].Type {
		case DOT:
			lightBeam(wall, x, y-1, TOP)
			break
		case PIPE:
			lightBeam(wall, x, y-1, TOP)
			break
		case MINUS:
			lightBeam(wall, x, y-1, LEFT)
			lightBeam(wall, x, y-1, RIGHT)
			break
		case BACKSLASH:
			lightBeam(wall, x, y-1, LEFT)
			break
		case SLASH:
			lightBeam(wall, x, y-1, RIGHT)
			break
		}
		break
	case LEFT:
		if 0 == x {
			break
		}
		if slices.Contains((*wall)[y][x-1].Checked, LEFT) {
			break
		}
		(*wall)[y][x-1].Checked = append((*wall)[y][x-1].Checked, LEFT)
		switch (*wall)[y][x-1].Type {
		case DOT:
			lightBeam(wall, x-1, y, LEFT)
			break
		case PIPE:
			lightBeam(wall, x-1, y, TOP)
			lightBeam(wall, x-1, y, BOTTOM)
			break
		case MINUS:
			lightBeam(wall, x-1, y, LEFT)
			break
		case BACKSLASH:
			lightBeam(wall, x-1, y, TOP)
			break
		case SLASH:
			lightBeam(wall, x-1, y, BOTTOM)
			break
		}
		break
	case BOTTOM:
		if len(*wall)-1 == y {
			break
		}
		if slices.Contains((*wall)[y+1][x].Checked, BOTTOM) {
			break
		}
		(*wall)[y+1][x].Checked = append((*wall)[y+1][x].Checked, BOTTOM)
		switch (*wall)[y+1][x].Type {
		case DOT:
			lightBeam(wall, x, y+1, BOTTOM)
			break
		case PIPE:
			lightBeam(wall, x, y+1, BOTTOM)
			break
		case MINUS:
			lightBeam(wall, x, y+1, LEFT)
			lightBeam(wall, x, y+1, RIGHT)
			break
		case BACKSLASH:
			lightBeam(wall, x, y+1, RIGHT)
			break
		case SLASH:
			lightBeam(wall, x, y+1, LEFT)
			break
		}
		break
	case RIGHT:
		if len((*wall)[0])-1 == x {
			break
		}
		if slices.Contains((*wall)[y][x+1].Checked, RIGHT) {
			break
		}
		(*wall)[y][x+1].Checked = append((*wall)[y][x+1].Checked, RIGHT)
		switch (*wall)[y][x+1].Type {
		case DOT:
			lightBeam(wall, x+1, y, RIGHT)
			break
		case PIPE:
			lightBeam(wall, x+1, y, TOP)
			lightBeam(wall, x+1, y, BOTTOM)
			break
		case MINUS:
			lightBeam(wall, x+1, y, RIGHT)
			break
		case BACKSLASH:
			lightBeam(wall, x+1, y, BOTTOM)
			break
		case SLASH:
			lightBeam(wall, x+1, y, TOP)
			break
		}
		break
	}
}
