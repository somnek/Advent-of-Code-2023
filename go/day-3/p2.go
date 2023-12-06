package main

import (
	"Advent-of-Code-2023/utils"
	"fmt"
	"reflect"
)

var (
	zz = []int{-1, 0, 1} // use this to scan in 3x3 square
)

type Pos struct {
	x int
	y int
}

// Approach:
// - trying slower & more disgusting approach (nested loops)
func PartTwo() {
	lines := utils.ReadLines("input.txt")
	rowL := len(lines)
	colL := len(lines[0])
	var n, ans int

	track := map[string][]int{}

	for ri := 0; ri < rowL; ri++ {
		gears := []string{}
		n = 0
		hasPart := false
		for ci := 0; ci < colL+1; ci++ {
			if ci < colL && IsDigit(rune(lines[ri][ci])) {
				n = n*10 + utils.StrToInt(string(lines[ri][ci]))
				for _, z1 := range zz {
					for _, z2 := range zz {
						if 0 <= ri+z1 && ri+z1 < rowL && 0 <= ci+z2 && ci+z2 < colL {
							c := lines[ri+z1][ci+z2]
							if isSymbol(rune(c)) {
								hasPart = true
							}
							if c == '*' {
								ss := fmt.Sprintf("%d%d", ri+z1, ci+z2)
								if !contains(gears, ss) {
									gears = append(gears, ss)
								}
							}
						}
					}
				}
			} else if n > 0 {
				for _, gear := range gears {
					track[gear] = append(track[gear], n)
				}
				n = 0
				hasPart = false
				gears = []string{}
			}
			_ = hasPart
		}
	}

	for _, v := range track {
		if len(v) == 2 {
			ans += v[0] * v[1]
		}
	}
	fmt.Println(track)
	fmt.Println(ans)
}

func contains(pos []string, e string) bool {
	for _, p := range pos {
		if reflect.DeepEqual(p, e) {
			return true
		}
	}
	return false
}
