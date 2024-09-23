package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestRange(t *testing.T) {
	{
		actual := slices.Collect(Range(4, 10))
		helper.SliceEqual(t, []int{4, 5, 6, 7, 8, 9}, actual)
	}
	{
		actual := slices.Collect(Range(-3, 3))
		helper.SliceEqual(t, []int{-3, -2, -1, 0, 1, 2}, actual)
	}
	{
		actual := slices.Collect(Range(0, 1))
		helper.SliceEqual(t, []int{0}, actual)
	}
	{
		actual := slices.Collect(Range(0, 0))
		helper.SliceEqual(t, []int{}, actual)
	}
	{
		actual := slices.Collect(Range(1, 0))
		helper.SliceEqual(t, []int{}, actual)
	}
}

func TestNumbers(t *testing.T) {
	{
		actual := slices.Collect(Numbers(5))
		helper.SliceEqual(t, []int{0, 1, 2, 3, 4}, actual)
	}
	{
		actual := slices.Collect(Numbers(0))
		helper.SliceEqual(t, []int{}, actual)
	}
	{
		actual := slices.Collect(Numbers(1))
		helper.SliceEqual(t, []int{0}, actual)
	}
}
