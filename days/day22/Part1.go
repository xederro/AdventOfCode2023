package day22

import (
	"adventOfCode/utils"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type brick struct {
	Start coords
	End   coords
	Under []*brick
}

type coords struct {
	X int
	Y int
	Z int
}

func Part1() {
	scanner := utils.ReadFile("day22")
	bricks := []*brick{}
	count := 0

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "~")
		start := strings.Split(s[0], ",")
		end := strings.Split(s[1], ",")

		sx, err := strconv.Atoi(start[0])
		if err != nil {
			return
		}

		sy, err := strconv.Atoi(start[1])
		if err != nil {
			return
		}

		sz, err := strconv.Atoi(start[2])
		if err != nil {
			return
		}

		ex, err := strconv.Atoi(end[0])
		if err != nil {
			return
		}

		ey, err := strconv.Atoi(end[1])
		if err != nil {
			return
		}

		ez, err := strconv.Atoi(end[2])
		if err != nil {
			return
		}

		bricks = append(bricks, &brick{
			Start: coords{
				X: sx,
				Y: sy,
				Z: sz,
			},
			End: coords{
				X: ex,
				Y: ey,
				Z: ez,
			},
		})
	}

	slices.SortFunc(bricks, func(a, b *brick) int {
		return cmp.Compare(min(a.Start.Z, a.End.Z), min(b.Start.Z, b.End.Z))
	})

	for _, b := range bricks {
		b.fall(bricks)
	}

	slices.SortFunc(bricks, func(a, b *brick) int {
		return cmp.Compare(min(a.Start.Z, a.End.Z), min(b.Start.Z, b.End.Z))
	})
	for _, b := range bricks {
		fmt.Println(b)
	}
	for _, c := range bricks {
		if !c.removable(bricks) {
			count++
		}
	}

	fmt.Println(count)
}

func (b *brick) removable(bricks []*brick) bool {
	for _, br := range bricks {
		if len(br.Under) == 1 && br.Under[0] == b {
			return true
		}
	}
	return false
}

func (b *brick) fall(bricks []*brick) {
	stuck := false
	for z := min(b.Start.Z, b.End.Z); !stuck && z != 1; z-- {
		stuck = !b.willFall(bricks, z)
		if !stuck {
			b.Start.Z--
			b.End.Z--
		}
	}
}

func (b *brick) willFall(bricks []*brick, z int) bool {
	for _, c := range bricks {
		if z-1 == max(c.Start.Z, c.End.Z) {
			if intersect(b.Start, b.End, c.Start, c.End) {
				b.Under = append(b.Under, c)
			}
		}
	}
	return len(b.Under) == 0
}

func intersect(A, B, C, D coords) bool {
	return B.X >= C.X &&
		A.X <= D.X &&
		B.Y >= C.Y &&
		A.Y <= D.Y
}
