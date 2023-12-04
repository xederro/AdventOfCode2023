package day2

import (
	"adventOfCode/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Part1() {
	scanner := utils.ReadFile("day2")
	sum := 0

	for scanner.Scan() {
		s := scanner.Text()
		stripped := strings.Split(s, " ")
		if n := len(stripped); n > 2 {
			possible := true
			id, err := strconv.Atoi(stripped[1][:len(stripped[1])-1])
			if err != nil {
				log.Fatalln(err)
			}

			for i := 2; i < n; i += 2 {
				c, err := strconv.Atoi(stripped[i])
				if err != nil {
					log.Fatalln(err)
				}
				switch stripped[i+1] {
				case "red,", "red;", "red":
					if c > 12 {
						possible = false
					}
					break
				case "green,", "green;", "green":
					if c > 13 {
						possible = false
					}
					break
				case "blue,", "blue;", "blue":
					if c > 14 {
						possible = false
					}
					break
				}
			}

			if possible {
				sum += id
			}
		}
	}
	fmt.Println(sum)
}
