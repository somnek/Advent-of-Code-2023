package main

import (
	"fmt"
	"math"
	"strings"

	"Advent-of-Code-2023/utils"
)

func PartOne() {
	lines := utils.ReadLines("input.txt")
	sum := solve(lines)
	fmt.Println("part 1:", sum)
}

func solve(lines []string) int {
	var sum float64
	for _, line := range lines {
		winPart, mePart := split(line)
		wc := intersect(winPart, mePart)
		c := len(wc)
		if len(wc) > 0 {
			sum += math.Pow(2, float64(c-1))
		}
	}

	return int(sum)
}

// want | has
func split(line string) ([]int, []int) {
	var left, right []int

	mainSplit := strings.Split(line, "|")
	leftSplitRaw := mainSplit[0]

	// convert
	left = process(strings.Split(leftSplitRaw, ":")[1])
	right = process(mainSplit[1])

	return left, right
}

// convert digits in the list str -> int
func process(s string) []int {
	var l []int
	split := strings.Split(s, " ")

	for _, a := range split {
		if a != "" {
			d := utils.StrToInt(a)
			l = append(l, d)
		}
	}
	return l
}

func intersect(a, b []int) []int {
	var l []int

	for _, aa := range a {
		for _, bb := range b {
			if aa == bb {
				l = append(l, aa)
			}
		}
	}
	return l
}
