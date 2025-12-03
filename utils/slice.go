package utils

import (
	"math"
	"slices"
)

func Unique[T comparable](s []T) []T {
	res := []T{}
	for _, x := range s {
		if !slices.Contains(res, x) {
			res = append(res, x)
		}
	}
	return res
}

func SliceMax(xs []int) int {
	m := math.MinInt
	for _, x := range xs {
		if x > m {
			x = m
		}
	}
	return m
}
