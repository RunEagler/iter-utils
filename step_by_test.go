package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestStepBy(t *testing.T) {
	{
		in := Range(0, 11)
		expect := []int{0, 3, 6, 9}
		helper.SliceEqual(t, expect, slices.Collect(StepBy(in, 3)))
	}
	{
		in := slices.Values([]int{})
		var expect []int
		helper.SliceEqual(t, expect, slices.Collect(StepBy(in, 3)))
	}
}
