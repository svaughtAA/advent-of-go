package day5

import (
	"advent-of-go/utils"
	"sort"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        5,
		Part:       2,
		Calculator: pt2,
	}
}

// sort ranges first, then merge one-by-one until
// all ranges are included in merged ranges
func (ims *IMS) MergeRanges() [][2]int {
	// sort by min of each, makes subsequent merging simpler
	sort.Slice(ims.ranges, func(i, j int) bool {
		return ims.ranges[i][0] < ims.ranges[j][0]
	})

	// merge ranges
	merged := [][2]int{ims.ranges[0]}
	for i := 1; i < len(ims.ranges); i++ {
		r := ims.ranges[i]
		mr := &merged[len(merged)-1]

		// because sorting above guarantees mr[0] > r[0]
		// we only need to check if mr[1] > r[0] to determine
		// if the two ranges are overlapping or if one includes
		// the other
		if r[0] <= mr[1]+1 {
			mr[1] = max(r[1], mr[1])
		} else {
			merged = append(merged, r)
		}
	}
	return merged
}

func (ims *IMS) CountPotentiallyFreshIds() int {
	count := 0
	merged := ims.MergeRanges()
	for _, r := range merged {
		count += r[1] - r[0] + 1
	}
	return count
}

func pt2(input string) (string, error) {
	ims := ParseInput(input)
	cnt := ims.CountPotentiallyFreshIds()
	return strconv.Itoa(cnt), nil
}
