package day4

import (
	"advent-of-go/utils"
	"math"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        4,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	grid := ParseInput(input)
	count := 0
	justRemoved := math.MaxInt
	toRemove := [][2]int{}
	for justRemoved > 0 {
		for i := range len(grid) {
			for j := range len(grid[i]) {
				n := CountNeighbors(grid, i, j)
				if grid[i][j] == "@" && n < 4 {
					count++
					toRemove = append(toRemove, [2]int{i, j})
				}
			}
		}
		for _, coords := range toRemove {
			i, j := coords[0], coords[1]
			grid[i][j] = "."
		}
		justRemoved = len(toRemove)
		toRemove = [][2]int{}
	}
	return strconv.Itoa(count), nil
}
