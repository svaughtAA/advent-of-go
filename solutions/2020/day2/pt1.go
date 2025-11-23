package day2

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

type Policy struct {
	min      int
	max      int
	char     string
	password string
}

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        2,
		Part:       1,
		Calculator: pt1,
	}
}

// examples:
// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
func ParsePolicies(input string) []*Policy {
	lines := strings.Split(input, "\n")
	policies := []*Policy{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		policyPart := parts[0]
		password := parts[1]

		policyParts := strings.Split(policyPart, " ")
		rng := strings.Split(policyParts[0], "-")
		chr := policyParts[1]

		min, err := strconv.Atoi(rng[0])
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(rng[1])
		if err != nil {
			panic(err)
		}

		policies = append(policies, &Policy{
			password: password,
			min:      min,
			max:      max,
			char:     chr,
		})
	}
	return policies
}

func pt1(input string) (string, error) {
	policies := ParsePolicies(input)
	valid := 0
	for _, p := range policies {
		occ := strings.Count(p.password, p.char)
		if occ >= p.min && occ <= p.max {
			valid += 1
		}
	}
	return strconv.Itoa(valid), nil
}
