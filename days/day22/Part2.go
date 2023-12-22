package day22

import (
	"adventOfCode/utils"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type brick2 struct {
	Start coords2
	End   coords2
	Would int
	Under []*brick2
	Above []*brick2
}

type coords2 struct {
	X int
	Y int
	Z int
}

func Part2() {
	scanner := utils.ReadFile("day22")
	bricks := []*brick2{}
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

		bricks = append(bricks, &brick2{
			Start: coords2{
				X: sx,
				Y: sy,
				Z: sz,
			},
			End: coords2{
				X: ex,
				Y: ey,
				Z: ez,
			},
		})
	}

	slices.SortFunc(bricks, func(a, b *brick2) int {
		return cmp.Compare(min(a.Start.Z, a.End.Z), min(b.Start.Z, b.End.Z))
	})

	for _, b := range bricks {
		b.fall(bricks)
	}

	slices.SortFunc(bricks, func(a, b *brick2) int {
		return cmp.Compare(min(a.Start.Z, a.End.Z), min(b.Start.Z, b.End.Z))
	})
	for _, b := range bricks {
		b.destroy()
		count += b.Would
	}

	fmt.Println(count - len(bricks))
}

func (b *brick2) destroy() int {
	var curr utils.Queue[brick2]
	var wbd []*brick2

	curr.Enqueue(b)
	wbd = append(wbd, b)
	for !curr.IsEmpty() {
		u := curr.Dequeue()
		for _, v := range u.Above {
			un := 0
			for _, und := range v.Under {
				if !slices.Contains(wbd, und) {
					un++
				}
			}
			if un == 0 {
				if !slices.Contains(wbd, v) {
					wbd = append(wbd, v)
				}
				if !slices.Contains(curr, v) {
					curr.Enqueue(v)
				}
			}
		}
	}
	b.Would = len(wbd)
	return len(wbd)
}

func (b *brick2) calc() int {
	var curr utils.Queue[brick2]
	var wbd []*brick2

	count := 0

	curr.Enqueue(b)
	wbd = append(wbd, b)
	for !curr.IsEmpty() {
		u := curr.Dequeue()
		for _, v := range u.Above {
			un := 0
			for _, und := range v.Under {
				if !slices.Contains(wbd, und) {
					un++
				}
			}
			if un == 0 {
				if !slices.Contains(wbd, v) {
					wbd = append(wbd, v)
				}
				if !slices.Contains(curr, v) {
					curr.Enqueue(v)
				}
			}
		}
	}

	for _, c := range wbd {
		count += c.Would
	}

	return count - b.Would
}

func (b *brick2) fall(bricks []*brick2) {
	stuck := false
	for z := min(b.Start.Z, b.End.Z); !stuck && z != 1; z-- {
		stuck = !b.willFall(bricks, z)
		if !stuck {
			b.Start.Z--
			b.End.Z--
		}
	}
}

func (b *brick2) willFall(bricks []*brick2, z int) bool {
	for _, c := range bricks {
		if z-1 == max(c.Start.Z, c.End.Z) {
			if intersect2(b.Start, b.End, c.Start, c.End) {
				b.Under = append(b.Under, c)
				c.Above = append(c.Above, b)
			}
		}
	}
	return len(b.Under) == 0
}

func intersect2(A, B, C, D coords2) bool {
	return B.X >= C.X &&
		A.X <= D.X &&
		B.Y >= C.Y &&
		A.Y <= D.Y
}
