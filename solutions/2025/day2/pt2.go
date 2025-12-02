package day2

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        2,
		Part:       2,
		Calculator: pt2,
	}
}

func IsRepeatedMultiple(n int) bool {
	str := strconv.Itoa(n)
	length := len(str)
	for i := 1; i < length/2+1; i++ {
		end := str[:i]
		target := ""
		for range length / i {
			target += end
		}
		if str == target {
			return true
		}
	}
	return false
}

func pt2(input string) (string, error) {
	ranges := ParseInput(input)
	sum := 0
	for _, r := range ranges {
		for i := r.min; i <= r.max; i++ {
			if IsRepeatedMultiple(i) {
				sum += i
			}
		}
	}
	return strconv.Itoa(sum), nil
}
