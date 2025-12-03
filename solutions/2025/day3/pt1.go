package day3

import (
	"advent-of-go/utils"
	"math"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        3,
		Part:       1,
		Calculator: pt1,
	}
}

func GetLargestJoltage(b []int, n int) int {
	// if n==1 then look for max across entire array
	if n == 1 {
		mx := 0
		for _, x := range b {
			if x > mx {
				mx = x
			}
		}
		return mx
	}

	// if n>1 then find the largest element in the array
	// leaving space at the end for n-1 digits
	idx := 0
	for i, x := range b[:len(b)-n+1] {
		if x > b[idx] {
			idx = i
		}
	}
	place := int(math.Pow10(n - 1))
	return b[idx]*place + GetLargestJoltage(b[idx+1:], n-1)
}

func ParseInput(input string) [][]int {
	banks := [][]int{}
	for line := range strings.SplitSeq(input, "\n") {
		bnk := []int{}
		for char := range strings.SplitSeq(line, "") {
			b, _ := strconv.Atoi(char)
			bnk = append(bnk, b)
		}
		banks = append(banks, bnk)
	}
	return banks
}

func pt1(input string) (string, error) {
	banks := ParseInput(input)
	sum := 0
	for _, bank := range banks {
		sum += GetLargestJoltage(bank, 2)
	}
	return strconv.Itoa(sum), nil
}
