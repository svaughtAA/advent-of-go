package day6

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2025,
		Day:        6,
		Part:       1,
		Calculator: pt1,
	}
}

const testInput = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   + `

func Transpose(matrix [][]int) [][]int {
	output := make([][]int, len(matrix[0]))
	for i := range matrix[0] {
		output[i] = make([]int, len(matrix))
	}
	for i := range matrix {
		for j := 0; j < len(matrix[0]); j++ {
			output[j][i] = matrix[i][j]
		}
	}
	return output
}

func ParseInput(input string) ([][]int, []string) {
	allLines := strings.Split(input, "\n")
	lines, opsLine := allLines[:len(allLines)-1], allLines[len(allLines)-1:][0]+" "
	rows := [][]int{}
	for _, l := range lines {
		nums := []int{}
		for str := range strings.FieldsSeq(l) {
			n, _ := strconv.Atoi(strings.TrimSpace(str))
			nums = append(nums, n)
		}
		rows = append(rows, nums)
	}
	ops := strings.Fields(opsLine)
	return rows, ops
}

func pt1(input string) (string, error) {
	rows, ops := ParseInput(input)
	nums := Transpose(rows)
	sum := 0
	for i, op := range ops {
		if op == "*" {
			sum += utils.Reduce(nums[i], 1, func(x, acc int) int {
				return acc * x
			})
		}
		if op == "+" {
			sum += utils.Reduce(nums[i], 0, func(x, acc int) int {
				return acc + x
			})
		}
	}
	return strconv.Itoa(sum), nil
}
