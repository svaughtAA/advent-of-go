package day6

import (
	"advent-of-go/utils"
	"slices"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        6,
		Part:       2,
		Calculator: pt2,
	}
}

// returns slice of operators and width of associated column
func ParseOperators(row string) ([]string, []int) {
	ops := []string{}
	wds := []int{}
	runes := []rune(row)
	width := 0
	for i := len(runes) - 1; i >= 0; i-- {
		width++
		if runes[i] != ' ' {
			wds = append(wds, width)
			ops = append(ops, string(runes[i]))
			width = 0
		}
	}
	slices.Reverse(ops)
	slices.Reverse(wds)
	return ops, wds
}

func ParseNumberGrid(rows []string) [][]string {
	grid := [][]string{}
	for _, r := range rows {
		grid = append(grid, strings.Split(r, ""))
	}
	return grid
}

func GetSubGrid(grid [][]string, wds []int, n int) [][]string {
	offset := 0
	width := wds[n]
	for i, wd := range wds {
		offset += wd
		if i == n {
			break
		}
	}
	subGrid := make([][]string, len(grid))
	for i := range subGrid {
		subGrid[i] = grid[i][offset-width : offset]
	}
	return subGrid
}

func TransposeGrid(grid [][]string) [][]string {
	output := make([][]string, len(grid[0]))
	for i := range grid[0] {
		output[i] = make([]string, len(grid))
	}
	for i := range grid {
		for j := 0; j < len(grid[0]); j++ {
			output[j][i] = grid[i][j]
		}
	}
	return output
}

func ParseInputDifferently(input string) ([][]string, []string, []int) {
	allLines := strings.Split(input, "\n")
	rows, opsLine := allLines[:len(allLines)-1], allLines[len(allLines)-1:][0]+" "
	grid := ParseNumberGrid(rows)
	ops, wds := ParseOperators(opsLine)
	return grid, ops, wds
}

func MapToInts(values [][]string) []int {
	nums := []int{}
	for _, parts := range values {
		str := strings.TrimSpace(strings.Join(parts, ""))
		int, _ := strconv.Atoi(str)
		if int != 0 { // dumb edge case god I hate grids
			nums = append(nums, int)
		}
	}
	return nums
}

func pt2(input string) (string, error) {
	grid, ops, wds := ParseInputDifferently(input)
	sum := 0
	for i, op := range ops {
		g := GetSubGrid(grid, wds, i)
		t := TransposeGrid(g)
		nums := MapToInts(t)
		if op == "*" {
			sum += utils.Reduce(nums, 1, func(x, acc int) int {
				return acc * x
			})
		}
		if op == "+" {
			sum += utils.Reduce(nums, 0, func(x, acc int) int {
				return acc + x
			})
		}
	}
	return strconv.Itoa(sum), nil
}
