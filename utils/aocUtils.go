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
	for i := 0; i < len(haystack); i++ {
		j := strings.Index(haystack[i:], needle)
		if j < 0 {
			break
		}
		i += j
		x = append(x, i)
	}
	return x
}

func LCM(i ...int64) int64 {
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
