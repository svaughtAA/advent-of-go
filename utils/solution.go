package utils

import (
	"fmt"
	"time"
)

type Calculator func(string) (string, error)

type Solution struct {
	Year, Day, Part int
	Calculator      Calculator
}

func (s Solution) Name() string {
	return fmt.Sprintf("%d Day %d Pt %d", s.Year, s.Day, s.Part)
}

func (s Solution) Calculate() (string, error) {
	start := time.Now()
	inputPath := fmt.Sprintf("private/inputs/%d/day%d.txt", s.Year, s.Day)
	contents, e := GetFileContents(inputPath)
	if e != nil {
		return "", fmt.Errorf("error getting contents of input file for %s: %w", s.Name(), e)
	}
	result, err := s.Calculator(contents)
	elapsed := time.Since(start)
	c := "\033[36m"
	if elapsed > time.Millisecond {
		c = "\033[35m"
	}
	if elapsed > time.Millisecond*10 {
		c = "\033[31m"
	}
	fmt.Printf("%s[Y %d][d %02d][p %02d] Execution time: %s\n\033[0m", c, s.Year, s.Day, s.Part, elapsed)
	return result, err
}
