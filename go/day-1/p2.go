package main

import (
	"Advent-of-Code-2023/utils"
	"fmt"
	"log"
	"strconv"

	"strings"
)

var (
	M = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
	}
)

func PartTwo() {
	var num, sum int

	lines := utils.ReadLines(INPUT)
	for i, line := range lines {
		num = findNum2(line)
		sum += num
		fmt.Println(i, line, num)
	}
	fmt.Println(sum)
}

func findNum2(s string) int {
	var err error
	var num int
	var lFound, rFound bool

	size := len(s)
	numArray := make([]string, 2)
	converted := -1

	for i := 0; i < size; i++ {
		lIdx, rIdx := i, size-1-i
		lStr, rStr := string(s[lIdx]), string(s[rIdx])

		// left
		if !lFound {
			converted = letterToDigit(s[lIdx:size])
			if isInt(lStr) {
				numArray[0] = lStr
				lFound = true
			} else if converted > -1 {
				numArray[0] = fmt.Sprintf("%d", converted)
				lFound = true
			}
		}

		// right
		if !rFound {
			converted = letterToDigit(s[rIdx:size])
			if isInt(rStr) {
				numArray[1] = rStr
				rFound = true
			} else if converted > -1 {
				numArray[1] = fmt.Sprintf("%d", converted)
				rFound = true
			}
		}
	}

	fmt.Println("array: ", numArray)
	joined := strings.Join(numArray, "")
	if num, err = strconv.Atoi(joined); err != nil {
		log.Fatalf("%v | %v for %v", err, joined, s)
	}
	return num
}

func letterToDigit(chunk string) int {
	for a, v := range M {
		if len(chunk) >= len(a) {
			if chunk[:len(a)] == a {
				return v
			}
		}
	}
	return -1
}
