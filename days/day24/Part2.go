package day24

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
)

type hail2 struct {
	P coord2
	V vector2
}

type coord2 struct {
	X int64
	Y int64
	Z int64
}

type vector2 struct {
	DX int64
	DY int64
	DZ int64
}

func Part2() {
	scanner := utils.ReadFile("day24")
	var hailstones []*hail2

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

		hailstones = append(hailstones, &hail2{
			P: coord2{
				X: int64(px),
				Y: int64(py),
				Z: int64(pz),
			},
			V: vector2{
				DX: int64(dx),
				DY: int64(dy),
				DZ: int64(dz),
			},
		})
	}

	solve(*hailstones[0], *hailstones[1], *hailstones[2])
}

// https://www.wolframalpha.com/widgets/view.jsp?id=87f689e302f4424451632785e8b63e16
func solve(h1, h2, h3 hail2) {
	//0 = x + vxt*t1 - x1 - xv1*t1
	fmt.Printf("0=x+a*t-%d-%d*t\n", h1.P.X, h1.V.DX)
	//0 = y + vyt*t1 - y1 - yv1*t1
	fmt.Printf("0=y+b*t-%d-%d*t\n", h1.P.Y, h1.V.DY)
	//0 = z + vzt*t1 - z1 - zv1*t1
	fmt.Printf("0=z+c*t-%d-%d*t\n", h1.P.Z, h1.V.DZ)
	//0 = x + vxt*t2 - x2 - xv2*t2
	fmt.Printf("0=x+a*r-%d-%d*r\n", h2.P.X, h2.V.DX)
	//0 = y + vyt*t2 - y2 - yv2*t2
	fmt.Printf("0=y+b*r-%d-%d*r\n", h2.P.Y, h2.V.DY)
	//0 = z + vzt*t2 - z2 - zv2*t2
	fmt.Printf("0=z+c*r-%d-%d*r\n", h2.P.Z, h2.V.DZ)
	//0 = x + vxt*t3 - x3 - xv3*t3
	fmt.Printf("0=x+a*s-%d-%d*s\n", h3.P.X, h3.V.DX)
	//0 = y + vyt*t3 - y3 - yv3*t3
	fmt.Printf("0=y+b*s-%d-%d*s\n", h3.P.Y, h3.V.DY)
	//0 = z + vzt*t3 - z3 - zv3*t3
	fmt.Printf("0=z+c*s-%d-%d*s\n", h3.P.Z, h3.V.DZ)

	//https://www.wolframalpha.com/widgets/view.jsp?id=87f689e302f4424451632785e8b63e16 paste it here :V
}
