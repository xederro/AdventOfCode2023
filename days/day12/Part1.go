package day12

import (
	"adventOfCode/utils"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	scanner := utils.ReadFile("day12")
	possible := 0

	digits, err := regexp.Compile(`\d+`)
	if err != nil {
		return
	}

	for scanner.Scan() {
		s := scanner.Text()
		split := strings.Split(s, " ")
		dig := toInt(digits.FindAllString(split[1], -1))
		digSum := len(dig) - 1

		count := 0

		for _, v := range dig {
			digSum += v
		}

		for i := len(dig); i < 6; i++ {
			dig = append(dig, 0)
		}

		if digSum == len(split[0]) {
			count++
		} else {
			posStart := [6]int{len(split[0]) - 1, len(split[0]) - 1, len(split[0]) - 1, len(split[0]) - 1, len(split[0]) - 1, len(split[0]) - 1}
			posEnd := [6]int{len(split[0]) - 1, len(split[0]) - 1, len(split[0]) - 1, len(split[0]) - 1, len(split[0]) - 1, len(split[0]) - 1}
			posV := 0

			for k, v := range dig {
				posStart[k] = posV
				posEnd[k] = len(split[0]) - digSum + posV
				posV += v + 1
			}

			var possibleChecks []string

			for i := posStart[0]; i <= posEnd[0]; i++ {
				for j := posStart[1]; j <= posEnd[1]; j++ {
					if i+dig[0] < j {
						for k := posStart[2]; k <= posEnd[2]; k++ {
							if j+dig[1] < k {
								for l := posStart[3]; l <= posEnd[3]; l++ {
									if k+dig[2] < l {
										for m := posStart[4]; m <= posEnd[4]; m++ {
											if l+dig[3] < m {
												for n := posStart[5]; n <= posEnd[5]; n++ {
													if m+dig[4] < n {
														place := ""
														for o := 0; o < len(split[0]); o++ {
															if o >= i && o < i+dig[0] ||
																o >= j && o < j+dig[1] ||
																o >= k && o < k+dig[2] ||
																o >= l && o < l+dig[3] ||
																o >= m && o < m+dig[4] ||
																o >= n && o < n+dig[5] {
																place += "#"
															} else {
																place += "."
															}
														}
														if !slices.Contains(possibleChecks, place) {
															possibleChecks = append(possibleChecks, place)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}

			for _, check := range possibleChecks {
				repl := strings.Replace(check, ".", "[?.]", -1)
				repl = strings.Replace(repl, "#", "[?#]", -1)
				compile, err := regexp.Compile(repl)
				if err != nil {
					fmt.Println(err)
					return
				}
				if compile.MatchString(split[0]) {
					count++
				}
			}
		}
		possible += count
	}

	fmt.Println(possible)
}

func toInt(s []string) []int {
	var i []int
	for _, str := range s {
		atoi, err := strconv.Atoi(str)
		if err != nil {
			return nil
		}
		i = append(i, atoi)
	}
	return i
}
