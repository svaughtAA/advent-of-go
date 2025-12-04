package day4

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        4,
		Part:       1,
		Calculator: pt1,
	}
}

func ParseInput(input string) [][]string {
	cols := [][]string{}
	for line := range strings.SplitSeq(input, "\n") {
		cols = append(cols, strings.Split(line, ""))
	}
	return cols
}

func AddIfPaperRoll(grid [][]string, x, y int, acc *int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]) {
		return
	}
	if grid[x][y] == "@" {
		*acc += 1
	}
}

func CountNeighbors(grid [][]string, x, y int) int {
	neighbors := 0

	// cross orientation
	AddIfPaperRoll(grid, x, y-1, &neighbors)
	AddIfPaperRoll(grid, x, y+1, &neighbors)
	AddIfPaperRoll(grid, x-1, y, &neighbors)
	AddIfPaperRoll(grid, x+1, y, &neighbors)

	// corners
	AddIfPaperRoll(grid, x-1, y-1, &neighbors)
	AddIfPaperRoll(grid, x+1, y-1, &neighbors)
	AddIfPaperRoll(grid, x-1, y+1, &neighbors)
	AddIfPaperRoll(grid, x+1, y+1, &neighbors)
	return neighbors
}

func pt1(input string) (string, error) {
	grid := ParseInput(input)
	count := 0
	for i := range len(grid) {
		for j := range len(grid[i]) {
			n := CountNeighbors(grid, i, j)
			if grid[i][j] == "@" && n < 4 {
				count++
			}
		}
	}
	return strconv.Itoa(count), nil
}
