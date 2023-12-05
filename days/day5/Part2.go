// it runs for ~2 minutes. Second smallest value was my answer I'm yet to fint why it returns wront smallest number.
// I have some ideas but time is problem
package day5

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type seed2 struct {
	Id          int64
	Soil        int64
	Fertilizer  int64
	Water       int64
	Light       int64
	Temperature int64
	Humidity    int64
	Location    int64
}

var wg sync.WaitGroup
var seedToSoil [][3]int64
var soilToFertilizer [][3]int64
var fertilizerToWater [][3]int64
var waterToLight [][3]int64
var lightToTemperature [][3]int64
var temperatureToHumidity [][3]int64
var humidityToLocation [][3]int64
var current *[][3]int64
var minimum int64 = -1
var strs []string

func Part2() {
	scanner := utils.ReadFile("day5")

	if scanner.Scan() {
		strs = strings.Split(scanner.Text(), " ")
	}

	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			continue
		}

		switch s {
		case "seed-to-soil map:":
			current = &seedToSoil
			break
		case "soil-to-fertilizer map:":
			current = &soilToFertilizer
			break
		case "fertilizer-to-water map:":
			current = &fertilizerToWater
			break
		case "water-to-light map:":
			current = &waterToLight
			break
		case "light-to-temperature map:":
			current = &lightToTemperature
			break
		case "temperature-to-humidity map:":
			current = &temperatureToHumidity
			break
		case "humidity-to-location map:":
			current = &humidityToLocation
			break
		default:
			break
		}

		for scanner.Scan() {
			t := scanner.Text()
			if t == "" {
				break
			}
			c := strings.Split(t, " ")
			dst, err := strconv.Atoi(c[0])
			if err != nil {
				return
			}
			src, err := strconv.Atoi(c[1])
			if err != nil {
				return
			}
			leng, err := strconv.Atoi(c[2])
			if err != nil {
				return
			}

			*current = append(*current, [3]int64{int64(src), int64(src + leng - 1), int64(dst)})
		}
	}

	for i := 1; i < len(strs); i += 2 {
		start, err := strconv.Atoi(strs[i])
		if err != nil {
			return
		}
		leng, err := strconv.Atoi(strs[i+1])
		if err != nil {
			return
		}
		k := int64(10)
		l := int64(leng)/k + 1
		for t := k - 1; t >= 0; t-- {
			wg.Add(1)
			j := t
			go func() {
				m := fill(int64(start)+(j*l), l)
				if minimum == -1 {
					minimum = m
				} else if m < minimum {
					minimum = m
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()
	fmt.Println(minimum) //second smallest???
}

func fill(start int64, leng int64) int64 {
	var m int64 = -1
	se := seed2{}
	for j := leng - 1; j >= 0; j-- {
		se.Id = start + j
		s := se.assign()

		if m == -1 {
			m = s
		} else if s < m {
			m = s
		}
	}
	fmt.Println(m)
	return m
}

func (s *seed2) assign() int64 {
	for _, v := range seedToSoil {
		if v[0] <= s.Id && v[1] >= s.Id {
			s.Soil = v[2] - v[0] + s.Id
		}
	}
	if s.Soil == 0 {
		s.Soil = s.Id
	}

	for _, v := range soilToFertilizer {
		if v[0] <= s.Soil && v[1] >= s.Soil {
			s.Fertilizer = v[2] - v[0] + s.Soil
		}
	}
	if s.Fertilizer == 0 {
		s.Fertilizer = s.Soil
	}

	for _, v := range fertilizerToWater {
		if v[0] <= s.Fertilizer && v[1] >= s.Fertilizer {
			s.Water = v[2] - v[0] + s.Fertilizer
		}
	}
	if s.Water == 0 {
		s.Water = s.Fertilizer
	}

	for _, v := range waterToLight {
		if v[0] <= s.Water && v[1] >= s.Water {
			s.Light = v[2] - v[0] + s.Water
		}
	}
	if s.Light == 0 {
		s.Light = s.Water
	}

	for _, v := range lightToTemperature {
		if v[0] <= s.Light && v[1] >= s.Light {
			s.Temperature = v[2] - v[0] + s.Light
		}
	}
	if s.Temperature == 0 {
		s.Temperature = s.Light
	}

	for _, v := range temperatureToHumidity {
		if v[0] <= s.Temperature && v[1] >= s.Temperature {
			s.Humidity = v[2] - v[0] + s.Temperature
		}
	}
	if s.Humidity == 0 {
		s.Humidity = s.Temperature
	}

	for _, v := range humidityToLocation {
		if v[0] <= s.Humidity && v[1] >= s.Humidity {
			s.Location = v[2] - v[0] + s.Humidity
		}
	}
	if s.Location == 0 {
		s.Location = s.Humidity
	}

	return s.Location
}
