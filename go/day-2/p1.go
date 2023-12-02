package main

import (
	"fmt"
	"math"
	"strings"

	"Advent-of-Code-2023/utils"
)

type Set struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	ID   int
	Sets []Set
}

const (
	elfRed   = 12
	elfGreen = 13
	elfBlue  = 14
)

// the approach is to find max of each color
func PartOne() {
	var sum int
	lines := utils.ReadLines("input.txt")
	for _, line := range lines {
		parsed := parseLine(line)
		sum += solve(parsed)
	}
	fmt.Println(sum)
}

func solve(game Game) int {
	// find max of each color
	var maxRed, maxGreen, maxBlue int

	for _, set := range game.Sets {
		maxRed = int(math.Max(float64(set.Red), float64(maxRed)))
		maxGreen = int(math.Max(float64(set.Green), float64(maxGreen)))
		maxBlue = int(math.Max(float64(set.Blue), float64(maxBlue)))

	}

	// return 0 if elf has lesser
	if maxRed > elfRed || maxGreen > elfGreen || maxBlue > elfBlue {
		return 0
	}

	// return id
	return game.ID

}

func parseLine(line string) Game {
	splitMain := strings.Split(line, ": ")
	partA, partB := splitMain[0], splitMain[1]

	// parse ID
	id := utils.StrToInt(strings.Split(partA, " ")[1])

	// parse Sets
	sets := parseSets(partB)

	return Game{
		ID:   id,
		Sets: sets,
	}
}

func parseSets(s string) []Set {
	var sets []Set

	// 3 sets per game
	raw := strings.Split(s, "; ") // e.g.: ["1 blue, 3 red, 2 green", ...]

	for _, setRaw := range raw { // e.g.: "1 blue, 3 red, 2 green"
		set := newSet(setRaw)
		sets = append(sets, set)
	}
	return sets
}

func newSet(s string) Set {
	var red, green, blue int
	pairs := strings.Split(s, ", ") // ["1 blue", "3 red",. ..]

	for _, pair := range pairs {
		split := strings.Split(pair, " ")
		count := utils.StrToInt(split[0])
		color := split[1]

		switch color {
		case "red":
			red = count
		case "green":
			green = count
		case "blue":
			blue = count
		}
	}

	return Set{
		Red:   red,
		Green: green,
		Blue:  blue,
	}

}
