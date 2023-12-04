package main

import (
	"strings"

	"Advent-of-Code-2023/utils"
)

// The approach:
// 1. Pad the input with "." on all sides
// 2. Iterate through the input and check if the current char is a number
// 3. If it is, check if it has adjacent symbols
// 4. If it does, add it to the sum

func PartOne() {
	var sum int

	lines := utils.ReadLines("input.txt")
	size := len(lines)
	padLines(&lines)

	for i := 1; i < size+1; i++ {
		nums := solve(i, lines)
		for _, num := range nums {
			sum += num
		}
	}
	println("sum: ", sum)
}

func solve(idx int, lines []string) []int {
	var num string
	var nums []int
	var containsSymbol bool

	cols := len(lines[idx])
	curr, prev, next := lines[idx], lines[idx-1], lines[idx+1]

	for i := 1; i < cols; i++ {
		c := rune(curr[i])
		if IsDigit(c) {
			num += string(c)
			if num != "" && containsSymbol {
				containsSymbol = true
			} else {
				containsSymbol = hasAdjc(i, curr, prev, next)
			}
		} else if num != "" {
			if containsSymbol {
				nums = append(nums, utils.StrToInt(num))
			}
			num = ""
			containsSymbol = false
		}
	}
	return nums
}

func checkRow(row string, i int) bool {
	for j := -1; j < 2; j++ {
		if isSymbol(rune(row[i+j])) {
			return true
		}
	}
	return false
}

func hasAdjc(i int, curr, prev, next string) bool {
	// usage
	if checkRow(prev, i) || checkRow(curr, i) || checkRow(next, i) {
		return true
	}
	return false
}

func padLines(sPtr *[]string) {
	s := *sPtr
	col := len(s[0]) + 2
	sLen := len(s)

	// pad left & right
	for i := 0; i < sLen; i++ {
		s[i] = "." + s[i] + "."
	}

	// pad top & bottoms
	empty := []string{strings.Repeat(".", col)}
	s = append(empty, s...)
	s = append(s, empty...)
	*sPtr = s
}

func isSymbol(c rune) bool {
	// is number or .
	if IsDigit(c) || (c == '.') {
		return false
	}
	return true
}

func IsDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}
