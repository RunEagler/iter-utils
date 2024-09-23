package it

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"iter-utils/internal/helper"
)

func TestFilter(t *testing.T) {
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		helper.SliceEqual(t, []int{0, 2, 4, 6, 8}, slices.Collect(Filter(Numbers(10), isEven)))
	}
	{
		isLenLTE := func(length int) func(v string) bool {
			return func(v string) bool {
				return length >= len(v)
			}
		}
		expect := []string{
			"banana",
			"apple",
			"orange",
			"peach",
			"cherry",
		}
		helper.SliceEqual(t, []string{expect[1], expect[3]}, slices.Collect(Filter(slices.Values(expect), isLenLTE(5))))
	}
}

func TestAll(t *testing.T) {
	isEven := func(v int) bool {
		return v%2 == 0
	}
	{
		assert.True(t, All(slices.Values([]int{0, 2, 4, 6, 8}), isEven))
	}
	{
		assert.False(t, All(slices.Values([]int{0, 2, 5, 6, 8}), isEven))
	}
	{
		assert.True(t, All(slices.Values([]int{}), isEven))
	}
}

func TestAny(t *testing.T) {
	isEven := func(v int) bool {
		return v%2 == 0
	}
	{
		assert.True(t, Any(slices.Values([]int{0, 2, 4, 6, 8}), isEven))
	}
	{
		assert.True(t, Any(slices.Values([]int{1, 3, 6, 7, 9}), isEven))
	}
	{
		assert.False(t, Any(slices.Values([]int{}), isEven))
	}
}
