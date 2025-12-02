package day5

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        5,
		Part:       1,
		Calculator: pt1,
	}
}

type BoardingPass struct {
	RowPartitions []string
	ColPartitions []string
}

func ParseInput(input string) []BoardingPass {
	lines := strings.Split(input, "\n")
	bps := make([]BoardingPass, len(lines))
	for i, line := range lines {
		bps[i] = BoardingPass{
			RowPartitions: strings.Split(line[:7], ""),
			ColPartitions: strings.Split(line[7:], ""),
		}
	}
	return bps
}

func partition(min int, max int, takeUpper bool) (int, int) {
	if takeUpper {
		return min + (max-min+1)/2, max
	} else {
		return min, max - (max-min+1)/2
	}
}

func (bp BoardingPass) FindRow() int {
	min, max := 0, 127
	for _, p := range bp.RowPartitions {
		min, max = partition(min, max, p == "B")
	}
	utils.Assert(min == max, "min doesn't equal max, bad `FindRow` logic")
	return min // should be same as max
}

func (bp BoardingPass) FindCol() int {
	min, max := 0, 7
	for _, p := range bp.ColPartitions {
		min, max = partition(min, max, p == "R")
	}
	utils.Assert(min == max, "min doesn't equal max, bad `FindCol` logic")
	return min // should be the same as max
}

func pt1(input string) (string, error) {
	bps := ParseInput(input)
	max := 0
	for _, bp := range bps {
		row, col := bp.FindRow(), bp.FindCol()
		max = utils.Max(max, row*8+col)
	}
	return strconv.Itoa(max), nil
}
