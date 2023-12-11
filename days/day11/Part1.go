package day11

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strings"
)

type coords struct {
	X int
	Y int
}

type galaxy struct {
	Coords coords
}

func Part1() {
	scanner := utils.ReadFile("day11")
	var mapa [][]string
	var rowValue []int
	var galaxies []*galaxy
	distance := 0
	scale := 2

	for i := 0; scanner.Scan(); i++ {
		var row []string
		s := strings.Split(scanner.Text(), "")
		count := 0
		for _, c := range s {
			row = append(row, c)
			if c == "." {
				count++
			}
		}

		mapa = append(mapa, row)
		rowValue = append(rowValue, i)
		if count >= len(s) {
			i += scale - 1
		}
	}

	move := 0
	for x := 0; x < len(mapa[0]); x++ {
		count := 0
		for y := 0; y < len(mapa); y++ {
			if mapa[y][x] == "." {
				count++
			} else {
				galaxies = append(galaxies, &galaxy{Coords: coords{X: x + move, Y: rowValue[y]}})
			}
		}

		if count >= len(mapa) {
			move += scale - 1
		}
	}

	for j := 0; j < len(galaxies)-1; j++ {
		for i := j + 1; i < len(galaxies); i++ {
			distance += galaxies[j].distance(galaxies[i])
		}
	}

	fmt.Println(distance)
}

func (g *galaxy) distance(t *galaxy) int {
	return int(math.Abs(float64(g.Coords.X-t.Coords.X)) + math.Abs(float64(g.Coords.Y-t.Coords.Y)))
}
