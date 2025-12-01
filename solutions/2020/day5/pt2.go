package day5

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        5,
		Part:       2,
		Calculator: pt2,
	}
}

func pt2(input string) (string, error) {
	bps := ParseInput(input)
	takenSeats := map[int]bool{}
	for _, bp := range bps {
		row, col := bp.FindRow(), bp.FindCol()
		takenSeats[row*8+col] = true
	}
	mySeat := 0
	for i := 0; i < len(takenSeats); i++ {
		if takenSeats[i+1] && takenSeats[i-1] && !takenSeats[i] {
			mySeat = i
		}
	}
	fmt.Printf("result %d", mySeat)
	return strconv.Itoa(mySeat), nil
}
