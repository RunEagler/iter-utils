package it

import (
	"iter"
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestRepeat(t *testing.T) {
	{
		helper.SliceEqual(t, []int{1, 1, 1}, slices.Collect(Take(Repeat[iter.Seq[int]](1), 3)))
	}
	{
		helper.SliceEqual(t, []int{}, slices.Collect(Take(Repeat[iter.Seq[int]](5), 0)))
	}
}
