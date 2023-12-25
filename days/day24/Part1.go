package day24

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
)

type hail struct {
	P coord
	V vector
}

type coord struct {
	X float64
	Y float64
	Z float64
}

type vector struct {
	DX float64
	DY float64
	DZ float64
}

func Part1() {
	count := 0
	scanner := utils.ReadFile("day24")
	var hailstones []*hail

	compile, err := regexp.Compile(`[,@ ]+`)
	if err != nil {
		return
	}

	for i := 0; scanner.Scan(); i++ {
		s := compile.Split(scanner.Text(), -1)
		px, err := strconv.Atoi(s[0])
		if err != nil {
			return
		}

		py, err := strconv.Atoi(s[1])
		if err != nil {
			return
		}

		pz, err := strconv.Atoi(s[2])
		if err != nil {
			return
		}

		dx, err := strconv.Atoi(s[3])
		if err != nil {
			return
		}

		dy, err := strconv.Atoi(s[4])
		if err != nil {
			return
		}

		dz, err := strconv.Atoi(s[5])
		if err != nil {
			return
		}

		hailstones = append(hailstones, &hail{
			P: coord{
				X: float64(px),
				Y: float64(py),
				Z: float64(pz),
			},
			V: vector{
				DX: float64(dx),
				DY: float64(dy),
				DZ: float64(dz),
			},
		})
	}

	for k, h1 := range hailstones {
		for i := k + 1; i < len(hailstones); i++ {
			if collide(*h1, *hailstones[i]) {
				count++
			}
		}
	}

	fmt.Println(count)
}

// A(x - x1) + B(y - y1) + C(z - z1) = 0
func collide(H1, H2 hail) bool {
	minc, maxc := float64(200000000000000), float64(400000000000000)

	px1, py1, dx1, dy1, px2, py2, dx2, dy2 := H1.P.X, H1.P.Y, H1.V.DX, H1.V.DY, H2.P.X, H2.P.Y, H2.V.DX, H2.V.DY
	if dx1*dy2 != dy1*dx2 {
		x := py1*dx1*dx2/(dx1*dy2-dy1*dx2) - px1*dy1*dx2/(dx1*dy2-dy1*dx2) - dx1*py2*dx2/(dx1*dy2-dy1*dx2) + dx1*px2*dy2/(dx1*dy2-dy1*dx2)

		a1 := x/dx1 - px1/dx1
		a2 := x/dx2 - px2/dx2

		y := py1 + a1*dy1

		//fmt.Println(H1, H2, x, y, a1, a2)

		if a1 < 0 || a2 < 0 {
			//past
			return false
		}

		if (x < minc || x > maxc) || (y < minc || y > maxc) {
			//outside
			return false
		}
	} else {
		//parallel
		return false
	}
	return true
}
