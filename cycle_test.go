package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestCycle(t *testing.T) {
	{
		in := Numbers(3)
		expect := []int{0, 1, 2, 0, 1, 2, 0}
		helper.SliceEqual(t, expect, slices.Collect(Take(Cycle(in), 7)))
	}
	{
		in := Numbers(3)
		expect := []int{0, 1}
		helper.SliceEqual(t, expect, slices.Collect(Take(Cycle(in), 2)))
	}
	{
		in := Numbers(3)
		expect := []int{}
		helper.SliceEqual(t, expect, slices.Collect(Take(Cycle(in), 0)))
	}
}
