package utils

import (
	"fmt"
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

func Remove[T any](s []T, index int) []T {
	fmt.Printf("Removing index %d from slice of len %d\n", index, len(s))
	return append(s[:index], s[index+1:]...)
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

func Reduce(xs []int, acc int, cb func(x, acc int) int) int {
	_acc := acc
	for _, x := range xs {
		_acc = cb(x, _acc)
	}
	return _acc
}
