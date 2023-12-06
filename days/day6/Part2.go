package day6

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"strconv"
	"sync"
)

type race2 struct {
	Time     int64
	Distance int64
}

var wg sync.WaitGroup

func Part2() {
	scanner := utils.ReadFile("day6")
	race := race2{}
	var ways int64 = 0
	divs := []int64{2, 3, 5, 7, 11, 13, 17, 19}

	compile, err := regexp.Compile(" +")
	if err != nil {
		return
	}

	if scanner.Scan() {
		s := compile.ReplaceAllString(scanner.Text(), "")
		atoi, err := strconv.Atoi(s[5:])
		if err != nil {
			return
		}
		race.Time = int64(atoi)
	}
	if scanner.Scan() {
		s := compile.ReplaceAllString(scanner.Text(), "")
		atoi, err := strconv.Atoi(s[9:])
		if err != nil {
			return
		}
		race.Distance = int64(atoi)
	}

	k := int64(1)

	for _, div := range divs {
		if race.Time%div == 0 {
			k += div
		}
	}

	t := race.Time / k

	for i := k - 1; i >= 0; i-- {
		wg.Add(1)
		i := i
		go func() {
			ways += race.CountWaysToWin(i*t, t)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ways)
}

func (r race2) CountWaysToWin(btime int64, span int64) int64 {
	w := int64(0)

	for i := btime; i < btime+span; i++ {
		if i*(r.Time-i) > r.Distance {
			w++
		}
	}
	return w
}
