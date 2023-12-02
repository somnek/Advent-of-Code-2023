package main

import (
	"Advent-of-Code-2023/utils"
	"fmt"
	"math"
)

func PartTwo() {
	var sum int
	lines := utils.ReadLines("input.txt")
	for _, line := range lines {
		parsed := parseLine(line)
		sum += solve2(parsed)
	}
	fmt.Println(sum)
}

func solve2(game Game) int {
	// find max of each color
	var maxRed, maxGreen, maxBlue int

	for _, set := range game.Sets {
		maxRed = int(math.Max(float64(set.Red), float64(maxRed)))
		maxGreen = int(math.Max(float64(set.Green), float64(maxGreen)))
		maxBlue = int(math.Max(float64(set.Blue), float64(maxBlue)))

	}

	return maxRed * maxGreen * maxBlue

}
