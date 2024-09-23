package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestPartition(t *testing.T) {
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{1, 2, 3, 4, 5, 6, 7})
		expect1 := []int{2, 4, 6}
		expect2 := []int{1, 3, 5, 7}
		a1, a2 := Partition(in, isEven)
		helper.SliceEqual(t, expect1, a1)
		helper.SliceEqual(t, expect2, a2)
	}
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{1, 3, 5, 7})
		expect1 := []int{}
		expect2 := []int{1, 3, 5, 7}
		a1, a2 := Partition(in, isEven)
		helper.SliceEqual(t, expect1, a1)
		helper.SliceEqual(t, expect2, a2)
	}
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{2, 4, 6})
		expect1 := []int{2, 4, 6}
		expect2 := []int{}
		a1, a2 := Partition(in, isEven)
		helper.SliceEqual(t, expect1, a1)
		helper.SliceEqual(t, expect2, a2)
	}
}
