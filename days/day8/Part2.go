package day8

import (
	"adventOfCode/utils"
	"fmt"
	"slices"
	"sync"
)

var (
	length2  int = 0
	count2   []int64
	tree2    map[string][2]string = map[string][2]string{}
	starting []string             = []string{}
	roads    []map[string][]int   = []map[string][]int{{}, {}, {}, {}, {}, {}}
	steps2   string
	wg       sync.WaitGroup
)

func Part2() {
	scanner := utils.ReadFile("day8")

	if scanner.Scan() {
		steps2 = scanner.Text()
		length2 = len(steps2)
		scanner.Scan()
	}

	for scanner.Scan() {
		s := scanner.Text()
		tree2[s[0:3]] = [2]string{s[7:10], s[12:15]}
		if s[2] == 'A' {
			starting = append(starting, s[0:3])
		}
	}

	for k, v := range starting {
		v := v
		k := k
		wg.Add(1)
		go func() {
			turnAll(k, v, 0)
		}()
	}
	wg.Wait()

	fmt.Println(roads)
	for _, road := range roads {
		for _, ints := range road {
			count2 = append(count2, int64(ints[0]))
		}
	}
	fmt.Println(SCM(count2))
}

func turnAll(pos int, step string, cur int) {
	if step[2] == 'Z' {
		if v, ok := roads[pos][fmt.Sprint(cur%length2, step)]; ok {
			slices.Contains(v, cur)
			wg.Done()
			return
		} else {
			roads[pos][fmt.Sprint(cur%length2, step)] = append(roads[pos][fmt.Sprint(cur, step)], cur)
		}
	}

	s := ""
	if steps2[cur%length2] == 'R' {
		s = tree2[step][1]
	} else {
		s = tree2[step][0]
	}
	go turnAll(pos, s, cur+1)
	return
}

func SCM(i []int64) int64 {
	var w int64 = 1

	for _, v := range i {
		w = (v * w) / GCF(v, w)
	}

	return w
}

func GCF(i1, i2 int64) int64 {
	if i2 == 0 {
		return i1
	}
	return GCF(i2, i1%i2)
}
