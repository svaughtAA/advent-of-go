package utils

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
