package day3

import (
	"advent-of-go/utils"
	"strconv"
	"strings"
)

func Pt1() utils.Solution {
	return utils.Solution{
		Year:       2020,
		Day:        3,
		Part:       1,
		Calculator: pt1,
	}
}

// example
// ..##.......
// #...#...#..
// .#....#..#.
// ..#.#...#.#
// .#...##..#.
// ..#.##.....
// .#.#.#....#
// .#........#
// #.##...#...
// #...##....#
// .#..#...#.#
func ParseIntoGrid(input string) [][]rune {
	lines := strings.Split(input, "\n")
	trees := [][]rune{}
	for _, line := range lines {
		trees = append(trees, []rune(line))
	}
	return trees
}

func pt1(input string) (string, error) {
	trees := ParseIntoGrid(input)
	xoffset := 0
	hit := 0
	for _, row := range trees {
		if row[xoffset%len(row)] == '#' {
			hit += 1
		}
		xoffset += 3
	}
	return strconv.Itoa(hit), nil
}
