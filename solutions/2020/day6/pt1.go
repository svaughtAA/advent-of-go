package day6

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        6,
		Part:       1,
		Calculator: pt1,
	}
}

type Group struct {
	person []string
}

func (ga Group) GetUniqueAnswers() map[rune]bool {
	unique := map[rune]bool{}
	for _, pa := range ga.person {
		for _, a := range pa {
			unique[a] = true
		}
	}
	return unique
}

func ParseInput(input string) []Group {
	groups := strings.Split(input, "\n\n")
	groupAnswersList := []Group{}
	for _, grp := range groups {
		ppl := strings.Split(grp, "\n")
		groupAnswers := Group{}
		for _, p := range ppl {
			groupAnswers.person = append(groupAnswers.person, p)
		}
		groupAnswersList = append(groupAnswersList, groupAnswers)
		groupAnswers = Group{}
	}
	return groupAnswersList
}

func pt1(input string) (string, error) {
	gal := ParseInput(input)
	sum := 0
	for _, ga := range gal {
		sum += len(ga.GetUniqueAnswers())
	}
	fmt.Printf("sum %d\n", sum)
	return strconv.Itoa(sum), nil
}
