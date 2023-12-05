package day5

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
	"strings"
)

type seed struct {
	Id          int
	Soil        int
	Fertilizer  int
	Water       int
	Light       int
	Temperature int
	Humidity    int
	Location    int
}

func Part1() {
	scanner := utils.ReadFile("day5")
	var seeds []seed
	minimum := -1
	var seedToSoil [][3]int
	var soilToFertilizer [][3]int
	var fertilizerToWater [][3]int
	var waterToLight [][3]int
	var lightToTemperature [][3]int
	var temperatureToHumidity [][3]int
	var humidityToLocation [][3]int
	var current *[][3]int

	if scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		for i := 1; i < len(s); i++ {
			atoi, err := strconv.Atoi(s[i])
			if err != nil {
				return
			}
			seeds = append(seeds, seed{Id: atoi})
		}
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

			*current = append(*current, [3]int{src, src + leng - 1, dst})
		}
	}

	for k := range seeds {
		for _, v := range seedToSoil {
			if v[0] <= seeds[k].Id && v[1] >= seeds[k].Id {
				seeds[k].Soil = v[2] - v[0] + seeds[k].Id
			}
		}
		if seeds[k].Soil == 0 {
			seeds[k].Soil = seeds[k].Id
		}

		for _, v := range soilToFertilizer {
			if v[0] <= seeds[k].Soil && v[1] >= seeds[k].Soil {
				seeds[k].Fertilizer = v[2] - v[0] + seeds[k].Soil
			}
		}
		if seeds[k].Fertilizer == 0 {
			seeds[k].Fertilizer = seeds[k].Soil
		}

		for _, v := range fertilizerToWater {
			if v[0] <= seeds[k].Fertilizer && v[1] >= seeds[k].Fertilizer {
				seeds[k].Water = v[2] - v[0] + seeds[k].Fertilizer
			}
		}
		if seeds[k].Water == 0 {
			seeds[k].Water = seeds[k].Fertilizer
		}

		for _, v := range waterToLight {
			if v[0] <= seeds[k].Water && v[1] >= seeds[k].Water {
				seeds[k].Light = v[2] - v[0] + seeds[k].Water
			}
		}
		if seeds[k].Light == 0 {
			seeds[k].Light = seeds[k].Water
		}

		for _, v := range lightToTemperature {
			if v[0] <= seeds[k].Light && v[1] >= seeds[k].Light {
				seeds[k].Temperature = v[2] - v[0] + seeds[k].Light
			}
		}
		if seeds[k].Temperature == 0 {
			seeds[k].Temperature = seeds[k].Light
		}

		for _, v := range temperatureToHumidity {
			if v[0] <= seeds[k].Temperature && v[1] >= seeds[k].Temperature {
				seeds[k].Humidity = v[2] - v[0] + seeds[k].Temperature
			}
		}
		if seeds[k].Humidity == 0 {
			seeds[k].Humidity = seeds[k].Temperature
		}

		for _, v := range humidityToLocation {
			if v[0] <= seeds[k].Humidity && v[1] >= seeds[k].Humidity {
				seeds[k].Location = v[2] - v[0] + seeds[k].Humidity
			}
		}
		if seeds[k].Location == 0 {
			seeds[k].Location = seeds[k].Humidity
		}

		if minimum == -1 {
			minimum = seeds[k].Location
		} else if seeds[k].Location < minimum {
			minimum = seeds[k].Location
		}
	}

	fmt.Println(minimum)
}
