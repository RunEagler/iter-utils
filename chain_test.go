package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestChain(t *testing.T) {
	{
		x := slices.Values([]int{1, 2, 3})
		y := slices.Values([]int{4, 5, 6})
		expect := []int{1, 2, 3, 4, 5, 6}
		helper.SliceEqual(t, expect, slices.Collect(Chain(x, y)))
	}
	{
		x := slices.Values([]int{})
		y := slices.Values([]int{4, 5, 6})
		expect := []int{4, 5, 6}
		helper.SliceEqual(t, expect, slices.Collect(Chain(x, y)))
	}
	{
		x := slices.Values([]int{1, 2, 3})
		y := slices.Values([]int{})
		expect := []int{1, 2, 3}
		helper.SliceEqual(t, expect, slices.Collect(Chain(x, y)))
	}
	{
		x := slices.Values([]int{})
		y := slices.Values([]int{})
		expect := []int{}
		helper.SliceEqual(t, expect, slices.Collect(Chain(x, y)))
	}
}
