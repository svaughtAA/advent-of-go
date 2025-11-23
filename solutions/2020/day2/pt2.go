package day2

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        2,
		Part:       2,
		Calculator: pt2,
	}
}

type NewPolicy struct {
	posA     int
	posB     int
	char     string
	password string
}

// examples:
// 1-3 a: abcde
// 1-3 b: cdefg
// 2-9 c: ccccccccc
func ParseNewPolicies(input string) []*NewPolicy {
	lines := strings.Split(input, "\n")
	policies := []*NewPolicy{}
	for _, line := range lines {
		parts := strings.Split(line, ":")
		policyPart := parts[0]
		password := parts[1]

		policyParts := strings.Split(policyPart, " ")
		rng := strings.Split(policyParts[0], "-")
		chr := policyParts[1]

		posA, err := strconv.Atoi(rng[0])
		if err != nil {
			panic(err)
		}
		posB, err := strconv.Atoi(rng[1])
		if err != nil {
			panic(err)
		}

		policies = append(policies, &NewPolicy{
			password: password,
			posA:     posA,
			posB:     posB,
			char:     chr,
		})
	}
	return policies
}

func pt2(input string) (string, error) {
	policies := ParseNewPolicies(input)
	valid := 0
	for _, p := range policies {
		charAtPosA := string(p.password[p.posA])
		charAtPosB := string(p.password[p.posB])
		if charAtPosA == p.char && charAtPosB != p.char {
			valid += 1
		}
		if charAtPosB == p.char && charAtPosA != p.char {
			valid += 1
		}
	}
	return strconv.Itoa(valid), nil
}
