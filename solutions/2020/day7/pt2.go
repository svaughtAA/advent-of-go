package day7

import (
	"advent-of-go/utils"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        7,
		Part:       2,
		Calculator: pt2,
	}
}

func CountBags(bag string, rules map[string]([]string)) int {
	children := rules[bag]
	if len(children) == 0 {
		return 1
	}
	result := 1
	for _, child := range children {
		result = result + CountBags(child, rules)
	}
	if bag == "shiny gold" {
		result -= 1
	}
	return result
}

func pt2(input string) (string, error) {
	tree := ParseInput(input)
	result := CountBags("shiny gold", tree)
	return strconv.Itoa(result), nil
}
