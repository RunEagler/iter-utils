package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestMap(t *testing.T) {
	{
		in := slices.Values([]int{1, 2, 3, 4, 5})
		expect := []int{1, 4, 9, 16, 25}
		toSquare := func(v int) int {
			return v * v
		}
		helper.SliceEqual(t, expect, slices.Collect(Map(in, toSquare)))
	}
	{
		in := slices.Values([]string{"a", "bb", "ccc", "dddd", "eeeee"})
		expect := []int{1, 2, 3, 4, 5}
		toLen := func(v string) int {
			return len(v)
		}
		helper.SliceEqual(t, expect, slices.Collect(Map(in, toLen)))
	}
	{
		in := slices.Values([]string{})
		var expect []int
		toLen := func(v string) int {
			return len(v)
		}
		helper.SliceEqual(t, expect, slices.Collect(Map(in, toLen)))
	}
}
