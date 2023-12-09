package main

import (
	"Advent-of-Code-2023/utils"
	"fmt"
	"strings"
)

// TODO: try Binary Search
func PartTwo() {
	// parse
	lines := utils.ReadLines("input.txt")
	tToBeat := utils.StrToInt(strings.Join(strings.Fields(lines[0])[1:], ""))
	dToBeat := utils.StrToInt(strings.Join(strings.Fields(lines[1])[1:], ""))

	ans := calc(tToBeat, dToBeat)
	fmt.Println("ans: ", ans)
}
