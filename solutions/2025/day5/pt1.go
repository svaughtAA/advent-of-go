package day5

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

const testInput = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        5,
		Part:       1,
		Calculator: pt1,
	}
}

type IMS struct {
	ranges [][2]int
	ids    []int
}

func ParseInput(input string) *IMS {
	ims := &IMS{ranges: [][2]int{}, ids: []int{}}
	parts := strings.Split(input, "\n\n")
	for line := range strings.SplitSeq(parts[0], "\n") {
		rangeParts := strings.Split(line, "-")
		min, _ := strconv.Atoi(rangeParts[0])
		max, _ := strconv.Atoi(rangeParts[1])
		r := [2]int{min, max}
		ims.ranges = append(ims.ranges, r)
	}
	for line := range strings.SplitSeq(parts[1], "\n") {
		id, _ := strconv.Atoi(line)
		ims.ids = append(ims.ids, id)
	}
	return ims
}

func (ims *IMS) CountFreshIds() int {
	count := 0
	for _, id := range ims.ids {
		fresh := false
		idx := 0
		for !fresh && idx < len(ims.ranges) {
			r := ims.ranges[idx]
			if r[0] <= id && id <= r[1] {
				fresh = true
			}
			idx++
		}
		if fresh {
			count++
		}
	}
	return count
}

func pt1(input string) (string, error) {
	ims := ParseInput(input)
	res := ims.CountFreshIds()
	return strconv.Itoa(res), nil
}
