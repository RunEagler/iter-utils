package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestScan(t *testing.T) {
	{
		in := slices.Values([]int{1, 2, 3, 4, 5})
		expect := []int{1, 3, 6, 10, 15}
		actual := slices.Collect(Scan(in, 0, func(sum int, v int) int {
			sum += v
			return sum
		}))
		helper.SliceEqual(t, expect, actual)
	}
	{
		in := slices.Values([]int{})
		var expect []int
		actual := slices.Collect(Scan(in, 0, func(sum int, v int) int {
			sum += v
			return sum
		}))
		helper.SliceEqual(t, expect, actual)
	}
}
