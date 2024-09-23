package it

import (
	"testing"

	"iter-utils/internal/helper"
)

func TestIterator_Collect(t *testing.T) {
	s := Slice([]int{1, 2, 3, 4, 5})

	helper.SliceEqual(t, []int{1, 2}, s.Take(2).ToCollect())
}
