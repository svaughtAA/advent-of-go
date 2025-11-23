package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        1,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	lines := strings.Split(input, "\n")
	numbers := ConvertLinesToInts(lines)
	for _, a := range numbers {
		for _, b := range numbers {
			for _, c := range numbers {
				sum := a + b + c
				if sum == 2020 {
					return strconv.Itoa(a * b * c), nil
				}
			}
		}
	}
	return "", fmt.Errorf("Couldn't get solution")
}

