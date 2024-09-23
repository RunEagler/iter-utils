package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestTake(t *testing.T) {
	{
		helper.SliceEqual(t, []int{0, 1}, slices.Collect(Take(Numbers(10), 2)))
	}
	{
		helper.SliceEqual(t, []int{3, 4, 5, 6}, slices.Collect(Take(Range(3, 10), 4)))
	}
	{
		helper.SliceEqual(t, []int{0, 1, 2, 3, 4}, slices.Collect(Take(Numbers(5), 10)))
	}
}
