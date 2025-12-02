package day6

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        6,
		Part:       2,
		Calculator: pt2,
	}
}

func (ga Group) GetQuestionsWithOnlyYes() map[rune]bool {
	seen := make(map[rune]int)
	for _, p := range ga.person {
		for _, q := range p {
			seen[q] += 1
		}
	}
	yes := make(map[rune]bool)
	for q, amt := range seen {
		if amt == len(ga.person) {
			yes[q] = true
		}
	}
	return yes
}

func pt2(input string) (string, error) {
	gal := ParseInput(input)
	sum := 0
	for _, ga := range gal {
		sum += len(ga.GetQuestionsWithOnlyYes())
	}
	fmt.Printf("sum %d\n", sum)
	return strconv.Itoa(sum), nil
}
