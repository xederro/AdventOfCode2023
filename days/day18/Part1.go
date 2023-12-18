package day18

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//type dig struct {
//	Coords coords
//	Colour string
//}

//type coords struct {
//	X int
//	Y int
//}

func Part1() {
	var digArea map[[2]int]string = map[[2]int]string{}
	scanner := utils.ReadFile("day18")
	last := [2]int{0, 0}

	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, " ")
		atoi, err := strconv.Atoi(split[1])
		if err != nil {
			return
		}

		colour := strings.Trim(split[2], "()")

		for i := 1; i <= atoi; i++ {
			var d [2]int
			switch split[0] {
			case "L":
				d = [2]int{last[0], last[1] - 1}
				break
			case "R":
				d = [2]int{last[0], last[1] + 1}
				break
			case "U":
				d = [2]int{last[0] - 1, last[1]}
				break
			case "D":
				d = [2]int{last[0] + 1, last[1]}
				break
			}

			last = d
			digArea[d] = colour
		}
	}

	arr := arraify(digArea)
	//for _, r := range arr {
	//	for _, s := range r {
	//		if s == "" {
	//			fmt.Print(" ")
	//		} else if s == "." {
	//			fmt.Print(".")
	//		} else {
	//			fmt.Print("#")
	//		}
	//	}
	//	fmt.Println()
	//}
	fmt.Println(count(arr))
}

func arraify(digArea map[[2]int]string) [][]string {
	sX, sY, bX, bY := math.MaxInt, math.MaxInt, -math.MaxInt, -math.MaxInt
	for ints := range digArea {
		if ints[0] < sY {
			sY = ints[0]
		}
		if ints[0] > bY {
			bY = ints[0]
		}
		if ints[1] < sX {
			sX = ints[1]
		}
		if ints[1] > bX {
			bX = ints[1]
		}
	}
	difX, difY := bX-sX+1, bY-sY+1
	var digSite [][]string

	for i := -1; i < difY+1; i++ {
		var digRow []string
		for j := -1; j < difX+1; j++ {
			if v, ok := digArea[[2]int{i + sY, j + sX}]; ok {
				digRow = append(digRow, v)
			} else {
				digRow = append(digRow, "")
			}
		}
		digSite = append(digSite, digRow)
	}
	outside(&digSite, 0, 0)

	return digSite
}

func outside(area *[][]string, y, x int) {
	(*area)[y][x] = "."
	if y > 0 && (*area)[y-1][x] == "" {
		outside(area, y-1, x)
	}
	if y < len(*area)-1 && (*area)[y+1][x] == "" {
		outside(area, y+1, x)
	}
	if x > 0 && (*area)[y][x-1] == "" {
		outside(area, y, x-1)
	}
	if x < len((*area)[y])-1 && (*area)[y][x+1] == "" {
		outside(area, y, x+1)
	}
}

func count(arr [][]string) int {
	res := 0
	for _, r := range arr {
		for _, s := range r {
			if s != "." {
				res++
			}
		}
	}
	return res
}
