package day2

import (
	"adventOfCode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part2() {
	scanner := utils.ReadFile("day2")
	sum := 0

	for scanner.Scan() {
		s := scanner.Text()
		stripped := strings.Split(s, " ")
		if n := len(stripped); n > 2 {
			r, g, b := 0, 0, 0

			for i := 2; i < n; i += 2 {
				c, err := strconv.Atoi(stripped[i])
				if err != nil {
					log.Fatalln(err)
				}
				switch stripped[i+1] {
				case "red,", "red;", "red":
					if c > r {
						r = c
					}
					break
				case "green,", "green;", "green":
					if c > g {
						g = c
					}
					break
				case "blue,", "blue;", "blue":
					if c > b {
						b = c
					}
					break
				}
			}

			sum += r * g * b
		}
	}
	fmt.Println(sum)
}
