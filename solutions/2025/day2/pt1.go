package day2

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        2,
		Part:       1,
		Calculator: pt1,
	}
}

type Range struct {
	min int
	max int
}

func ParseInput(input string) []Range {
	lines := strings.Split(input, ",")
	ranges := make([]Range, len(lines))
	for i, r := range lines {
		parts := strings.Split(r, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		ranges[i] = Range{min, max}
	}
	return ranges
}

func IsRepeated(n int) bool {
	str := strconv.Itoa(n)
	if len(str)&1 == 1 {
		return false
	}
	half := len(str) / 2
	return str[0:half] == str[half:]
}

func pt1(input string) (string, error) {
	ranges := ParseInput(input)
	sum := 0
	for _, r := range ranges {
		for i := r.min; i <= r.max; i++ {
			if IsRepeated(i) {
				sum += i
			}
		}
	}
	fmt.Printf("sum %v\n", sum)
	return strconv.Itoa(sum), nil
}
