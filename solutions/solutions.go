package solutions

import (
	"advent-of-go/utils"
	"slices"
	y2020 "advent-of-go/solutions/2020"
)

func Solutions() []utils.Solution {
	return slices.Concat[[]utils.Solution](y2020.Solutions())
}
