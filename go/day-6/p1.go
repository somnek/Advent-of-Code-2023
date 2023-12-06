package main

import (
	"Advent-of-Code-2023/utils"
	"fmt"
	"strings"
)

type Pair struct {
	T, D int // time distance
}

func PartOne() {
	ans := 1
	lines := utils.ReadLines("input.txt")
	pairs := ToPairs(lines)

	for _, p := range pairs {
		c := calc(p.T, p.D)
		ans *= c
	}
	fmt.Println("ans: ", ans)
}

func calc(tToBeat, dToBeat int) int {
	var c int
	for hold := 1; hold < tToBeat; hold++ {
		dRes := (tToBeat - hold) * hold
		if dRes > dToBeat {
			c++
		}
	}
	return c
}

func ToPairs(lines []string) []Pair {
	var pairs []Pair
	// parse
	times, dists := ParseLine(lines[0]), ParseLine(lines[1])
	size := len(times)
	for i := 0; i < size; i++ {
		pair := Pair{
			T: times[i],
			D: dists[i],
		}
		pairs = append(pairs, pair)
	}
	return pairs
}

func ParseLine(line string) []int {
	var l []int
	nums := strings.Fields(line)[1:]
	for _, n := range nums {
		l = append(l, utils.StrToInt(n))
	}
	return l
}
