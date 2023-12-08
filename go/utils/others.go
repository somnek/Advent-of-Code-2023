package utils

import (
	"strconv"
)

func GroupChar(s string) map[rune]int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	return m
}

func CountChar(s string, c rune) int {
	count := 0
	for _, v := range s {
		if v == c {
			count++
		}
	}
	return count
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
