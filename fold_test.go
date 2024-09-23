package it

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFold(t *testing.T) {
	{
		in := slices.Values([]int{1, 2, 3, 4, 5})
		expect := 15
		actual := Fold(in, 0, func(sum int, v int) int {
			sum += v
			return sum
		})
		assert.Equal(t, expect, actual)
	}
	{
		in := slices.Values([]int{})
		actual := Fold(in, 0, func(sum int, v int) int {
			sum += v
			return sum
		})
		assert.Equal(t, 0, actual)
	}
}
