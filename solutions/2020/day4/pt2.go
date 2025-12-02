package day4

import (
	"advent-of-go/utils"
	"regexp"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        4,
		Part:       2,
		Calculator: pt2,
	}
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func (p Passport) validateByr() bool {
	byr, err := strconv.Atoi(p["byr"])
	return err == nil && byr >= 1920 && byr <= 2002 && len(p["byr"]) == 4
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func (p Passport) validateIyr() bool {
	iyr, err := strconv.Atoi(p["iyr"])
	return err == nil && iyr >= 2010 && iyr <= 2020 && len(p["iyr"]) == 4
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func (p Passport) validateEyr() bool {
	eyr, err := strconv.Atoi(p["eyr"])
	return err == nil && eyr >= 2020 && eyr <= 2030 && len(p["eyr"]) == 4
}

// hgt (Height) - a number followed by either cm or in:
// If cm, the number must be at least 150 and at most 193.
// If in, the number must be at least 59 and at most 76.
func (p Passport) validateHgt() bool {
	hgt := p["hgt"]
	if len(hgt) < 3 {
		return false
	}
	hgtFormat := hgt[len(hgt)-2:]
	hgtValue := hgt[:len(hgt)-2]
	if hgtFormat == "cm" {
		hgtCm, err := strconv.Atoi(hgtValue)
		return err == nil && hgtCm >= 150 && hgtCm <= 193
	}
	if hgtFormat == "in" {
		hgtIn, err := strconv.Atoi(hgtValue)
		return err == nil && hgtIn >= 59 && hgtIn <= 76
	}
	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func (p Passport) validateHcl() bool {
	matched, err := regexp.MatchString("#([0-9a-f]){6}", p["hcl"])
	return err == nil && matched
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func (p Passport) validateEcl() bool {
	ecl := p["ecl"]
	return (ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth")
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func (p Passport) validatePid() bool {
	pid := p["pid"]
	_, err := strconv.Atoi(pid)
	return err == nil && len(pid) == 9
}

// cid (Country ID) - ignored, missing or not.
func (p Passport) validateCid() bool {
	return true
}

func ValidatePassport(p Passport) bool {
	if !p.HasAllReqFields() {
		return false
	}
	valid := true
	valid = p.validateCid() && valid
	valid = p.validateByr() && valid
	valid = p.validateIyr() && valid
	valid = p.validateEyr() && valid
	valid = p.validateHgt() && valid
	valid = p.validateHcl() && valid
	valid = p.validateEcl() && valid
	valid = p.validatePid() && valid
	return valid
}

func pt2(input string) (string, error) {
	passports := ParseToPassports(input)
	valid := 0
	for _, p := range passports {
		v := ValidatePassport(p)
		if v {
			valid++
		}
	}
	return strconv.Itoa(valid), nil
}
