package day3

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        3,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	banks := ParseInput(input)
	sum := 0
	for _, bank := range banks {
		sum += GetLargestJoltage(bank, 12)
	}
	return strconv.Itoa(sum), nil
}
