package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       1,
		Calculator: pt1,
	}
}

type Dial struct {
	position int
	count    int
}

// rotates the dial, counts every time the dial is left
// pointing at zero (not counting pass-throughs)
func (d *Dial) RotateCountingRestZeros(delta int) {
	for delta < 0 {
		delta += 100
	}
	d.position = (d.position + delta) % 100
	if d.position == 0 {
		d.count++
	}
}

func ParseRotations(input string) []int {
	lines := strings.Split(input, "\n")
	result := make([]int, len(lines))
	for i, line := range lines {
		n, _ := strconv.Atoi(line[1:])
		result[i] = n
		if line[0] == 'L' {
			result[i] *= -1
		}
	}
	return result
}

func pt1(input string) (string, error) {
	rotations := ParseRotations(input)
	d := Dial{50, 0}
	for _, r := range rotations {
		d.RotateCountingRestZeros(r)
	}
	fmt.Printf("count: %d\n", d.count)
	return strconv.Itoa(d.count), nil
}
