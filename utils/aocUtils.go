package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFile(fileName string) *bufio.Scanner {
	file, err := os.Open(fmt.Sprintf("files\\%s.txt", fileName))

	if err != nil {
		fmt.Println(err)
	}

	return bufio.NewScanner(file)
}

func FindIndices(haystack, needle string) []int {
	var x []int
	for i := 0; i < len(haystack)-len(needle); i++ {
		j := strings.Index(haystack[i:], needle)
		if j < 0 {
			break
		}
		i += j
		x = append(x, i)
	}
	return x
}
