package utils

import (
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
