package it

import (
	"slices"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestPosition(t *testing.T) {
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{1, 3, 6, 8})
		assert.Equal(t, 2, Position(in, isEven))
	}
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{1, 3, 5, 7})
		assert.Equal(t, -1, Position(in, isEven))
	}
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{})
		assert.Equal(t, -1, Position(in, isEven))
	}
}

func TestPositionQ(t *testing.T) {
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{1, 3, 6, 8})
		assert.Equal(t, mo.Some(2), PositionQ(in, isEven))
	}
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{1, 3, 5, 7})
		assert.Equal(t, mo.None[int](), PositionQ(in, isEven))
	}
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{})
		assert.Equal(t, mo.None[int](), PositionQ(in, isEven))
	}
}
