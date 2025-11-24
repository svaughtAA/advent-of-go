package day3

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        3,
		Part:       2,
		Calculator: pt2,
	}
}

func TestSlope(right int, down int, trees [][]rune) int {
	xoffset := 0
	hit := 0
	for i := 0; i < len(trees); i += down {
		row := trees[i]
		if row[xoffset%len(row)] == '#' {
			hit += 1
		}
		xoffset += right
	}
	return hit
}

// Right 1, down 1.
// Right 3, down 1. (This is the slope you already checked.)
// Right 5, down 1.
// Right 7, down 1.
// Right 1, down 2.
func pt2(input string) (string, error) {
	trees := ParseIntoGrid(input)
	allStrides := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	acc := 1
	for _, strides := range allStrides {
		var right, down = strides[0], strides[1]
		acc = acc * TestSlope(right, down, trees)
	}
	return strconv.Itoa(acc), nil
}
