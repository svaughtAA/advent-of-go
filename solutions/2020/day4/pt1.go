package day4

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        4,
		Part:       1,
		Calculator: pt1,
	}
}

type Passport map[string]string

// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
// byr:1937 iyr:2017 cid:147 hgt:183cm
//
// iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
// hcl:#cfa07d byr:1929
//
// hcl:#ae17e1 iyr:2013
// eyr:2024
// ecl:brn pid:760753108 byr:1931
// hgt:179cm
//
// hcl:#cfa07d eyr:2025 pid:166559648
// iyr:2011 ecl:brn hgt:59in
func ParseToPassports(input string) []Passport {
	passports := []Passport{}
	for info := range strings.SplitSeq(input, "\n\n") {
		passport := map[string]string{}
		chunks := strings.Split(info, "\n")
		for f := range strings.SplitSeq(strings.Join(chunks, " "), " ") {
			field := strings.Split(f, ":")
			passport[field[0]] = field[1]
		}
		passports = append(passports, passport)
	}
	return passports
}

var ReqFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
}

func (p Passport) HasAllReqFields() bool {
	for _, f := range ReqFields {
		if _, ok := p[f]; !ok {
			return false
		}
	}
	return true
}

// byr (Birth Year)
// iyr (Issue Year)
// eyr (Expiration Year)
// hgt (Height)
// hcl (Hair Color)
// ecl (Eye Color)
// pid (Passport ID)
// cid (Country ID) -- this one is optional
func pt1(input string) (string, error) {
	passports := ParseToPassports(input)
	valid := 0
	for _, p := range passports {
		if p.HasAllReqFields() {
			valid++
		}
	}
	return strconv.Itoa(valid), nil
}
