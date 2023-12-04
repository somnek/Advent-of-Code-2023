package main

import (
	"fmt"

	"Advent-of-Code-2023/utils"
)

func PartTwo() {
	lines := utils.ReadLines("input.txt")
	sum := solve2(lines)
	fmt.Println("part 2:", sum)
}

func solve2(lines []string) int {
	var sum int
	m := make(map[int]int)

	// map with each card as a key with val 1
	for x := 0; x < len(lines); x++ {
		m[x] = 1
	}

	for i, line := range lines {
		winPart, mePart := split(line)
		wc := intersect(winPart, mePart)
		for j := 0; j < len(wc); j++ {
			m[i+j+1] += m[i]
		}
	}

	fmt.Println(m)
	for _, v := range m {
		sum += v
	}
	return sum
}
