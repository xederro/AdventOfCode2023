package day15

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
)

type lens struct {
	Name  string
	Focal int
}

func Part2() {
	scanner := utils.ReadFile("day15")
	var res int
	var boxes [256][]lens

	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, ",")
		for _, str := range split {
			box(&boxes, str)
		}
	}
	for k, v := range boxes {
		for pos, val := range v {
			res += (k + 1) * (pos + 1) * val.Focal
		}
	}

	fmt.Println(res)
}

func hash2(str string) int {
	res := 0
	for _, v := range str {
		res = ((res + int(v)) * 17) % 256
	}
	return res
}

func box(boxes *[256][]lens, str string) {
	if strings.Contains(str, "=") {
		split := strings.Split(str, "=")
		atoi, err := strconv.Atoi(split[1])
		if err != nil {
			return
		}
		hash := hash2(split[0])
		if i := indexOf((*boxes)[hash], split[0]); i != -1 {
			(*boxes)[hash][i].Focal = atoi
		} else {
			(*boxes)[hash] = append((*boxes)[hash], lens{
				Name:  split[0],
				Focal: atoi,
			})
		}
	} else {
		hash := hash2(str[:len(str)-1])
		if i := indexOf((*boxes)[hash], str[:len(str)-1]); i != -1 {
			(*boxes)[hash] = append((*boxes)[hash][:i], (*boxes)[hash][i+1:]...)
		}
	}
}

func indexOf(box []lens, str string) int {
	for k, v := range box {
		if v.Name == str {
			return k
		}
	}
	return -1
}
