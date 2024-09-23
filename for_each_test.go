package it

import (
	"testing"

	"iter-utils/internal/helper"
)

func TestForEach(t *testing.T) {
	{
		expect := []int{0, 1, 4, 9, 16}
		var actual []int
		ForEach(Numbers(5), func(v int) {
			actual = append(actual, v*v)
		})
		helper.SliceEqual(t, expect, actual)
	}
}
