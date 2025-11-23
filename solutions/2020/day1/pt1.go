package day1

import (
	"advent-of-go/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        1,
		Part:       1,
		Calculator: pt1,
	}
}

func ConvertLinesToInts(lines []string) []int {
	numbers := []int{}
	for _, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(fmt.Sprintf("can't parse line '%s'", line))
		}
		numbers = append(numbers, num)
	}
	return numbers
}

func pt1(input string) (string, error) {
	lines := strings.Split(input, "\n")
	numbers := ConvertLinesToInts(lines)
	for _, a := range numbers {
		for _, b := range numbers {
			sum := a + b
			if sum == 2020 {
				return strconv.Itoa(a * b), nil
			}
		}
	}
	return "", errors.New("Unimplemented!")
}
