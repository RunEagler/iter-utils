package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestSkip(t *testing.T) {
	{
		in := slices.Values([]int{1, 2, 3, 4, 5})
		expect := []int{3, 4, 5}
		var actual []int
		for v := range Skip(in, 2) {
			actual = append(actual, v)
		}
		helper.SliceEqual(t, expect, actual)
	}
	{
		in := slices.Values([]int{1, 2, 3, 4, 5})
		var expect []int
		var actual []int
		for v := range Skip(in, 7) {
			actual = append(actual, v)
		}
		helper.SliceEqual(t, expect, actual)
	}
}
