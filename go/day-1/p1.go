package main

import (
	"Advent-of-Code-2023/utils"
	"fmt"
	"log"

	"strconv"
	"strings"
)

func PartOne() {
	var num, sum int

	lines := utils.ReadLines(INPUT)
	for _, line := range lines {
		num = findNum(line)
		sum += num
	}
	fmt.Println(sum)
}

func findNum(s string) int {
	size := len(s)
	numArray := make([]string, 2)
	var err error
	var num int

	for i := 0; i < size; i++ {
		lStr := string(s[i])
		rStr := string(s[size-1-i])

		// left
		if numArray[0] == "" && isInt(lStr) {
			numArray[0] = lStr
		}

		// right
		if numArray[1] == "" && isInt(rStr) {
			numArray[1] = rStr
		}
	}

	// join left & right
	joined := strings.Join(numArray, "")
	if num, err = strconv.Atoi(joined); err != nil {
		log.Fatalf("%v | %v for %v", err, joined, s)
	}
	return num
}

func isInt(c string) bool {
	if _, err := strconv.Atoi(c); err == nil {
		return true
	}
	return false
}
