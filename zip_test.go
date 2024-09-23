package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestZip(t *testing.T) {
	{
		x := slices.Values([]int{1, 2, 3, 4})
		y := slices.Values([]int{5, 6, 7, 8})
		expect := []Tuple[int, int]{
			{1, 5},
			{2, 6},
			{3, 7},
			{4, 8},
		}
		helper.SliceEqual(t, expect, slices.Collect(Zip(x, y)))
	}
	{
		x := slices.Values([]int{1, 2, 3, 4})
		y := slices.Values([]int{5, 6, 7})
		expect := []Tuple[int, int]{
			{1, 5},
			{2, 6},
			{3, 7},
		}
		helper.SliceEqual(t, expect, slices.Collect(Zip(x, y)))
	}
	{
		x := slices.Values([]int{1, 2, 3})
		y := slices.Values([]int{5, 6, 7, 8})
		expect := []Tuple[int, int]{
			{1, 5},
			{2, 6},
			{3, 7},
		}
		helper.SliceEqual(t, expect, slices.Collect(Zip(x, y)))
	}
	{
		x := slices.Values([]int{1, 2, 3})
		y := slices.Values([]string{"x", "ab", "tmt"})
		expect := []Tuple[int, string]{
			{1, "x"},
			{2, "ab"},
			{3, "tmt"},
		}
		helper.SliceEqual(t, expect, slices.Collect(Zip(x, y)))
	}
	{
		x := slices.Values([]int{})
		y := slices.Values([]string{})
		var expect []Tuple[int, string]
		helper.SliceEqual(t, expect, slices.Collect(Zip(x, y)))
	}
}
