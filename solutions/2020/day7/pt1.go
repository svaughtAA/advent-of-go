package day7

import (
	"advent-of-go/utils"
	"regexp"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        7,
		Part:       1,
		Calculator: pt1,
	}
}

func ParseInput(input string) map[string]([]string) {
	parentRegex := regexp.MustCompile(`^(?P<parent>.*?) bags contain (?P<rest>.*?)$`)
	childRegex := regexp.MustCompile(`.*?(?P<amt>[0-9]) (?P<child>.*?) bag?(?P<rest>.*?)$`)
	results := make(map[string]([]string))
	for rule := range strings.SplitSeq(input, "\n") {
		matches := parentRegex.FindStringSubmatch(rule)
		parent := matches[parentRegex.SubexpIndex("parent")]
		rest := matches[parentRegex.SubexpIndex("rest")]
		results[parent] = []string{}
		for childRegex.MatchString(rest) {
			matches := childRegex.FindStringSubmatch(rest)
			amt, _ := strconv.Atoi(matches[childRegex.SubexpIndex("amt")])
			child := matches[childRegex.SubexpIndex("child")]
			rest = matches[childRegex.SubexpIndex("rest")]
			for range amt {
				results[parent] = append(results[parent], child)
			}
		}
	}
	return results
}

var containsGold = map[string]bool{
	"shiny gold": true,
}

func ContainsShinyGold(bag string, rules map[string]([]string)) bool {
	hasGold, ok := containsGold[bag]
	if ok {
		return hasGold
	}
	children := utils.Unique(rules[bag])
	for _, child := range children {
		hasGold = ContainsShinyGold(child, rules)
		containsGold[child] = hasGold
		if hasGold {
			break
		}
	}
	containsGold[bag] = hasGold
	return hasGold
}

func pt1(input string) (string, error) {
	tree := ParseInput(input)
	count := 0
	for bag := range tree {
		if bag != "shiny gold" && ContainsShinyGold(bag, tree) {
			count++
		}
	}
	return strconv.Itoa(count), nil
}
