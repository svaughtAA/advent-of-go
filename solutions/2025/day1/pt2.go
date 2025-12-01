package day1

import (
	"advent-of-go/utils"
	"fmt"
	"strconv"
)

func Pt2() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        1,
		Part:       2,
		Calculator: pt2,
	}
}

// go doesn't have built in abs(int), only support floats
func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

// rotates the dial, counts every time the dial
// passes through or lands on zero
func (d *Dial) RotateCountingAllZeros(delta int) {
	newPosition := d.position + delta
	d.count += abs(newPosition / 100)
	if d.position != 0 && newPosition <= 0 {
		d.count++
	}
	for newPosition < 0 {
		newPosition += 100
	}
	d.position = newPosition % 100
}

func pt2(input string) (string, error) {
	rotations := ParseRotations(input)
	d := Dial{50, 0}
	for _, rdelta := range rotations {
		d.RotateCountingAllZeros(rdelta)
	}
	fmt.Printf("count %d\n", d.count)
	return strconv.Itoa(d.count), nil
}
