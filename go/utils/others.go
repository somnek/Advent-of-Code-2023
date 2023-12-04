package utils

import (
	"strconv"
)

func StrToInt(str string) int {

	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}
