package it

import (
	"slices"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestFindQ(t *testing.T) {
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{1, 2, 3, 4, 5})
		assert.Equal(t, mo.Some(2), FindQ(in, isEven))
	}
	{
		isEven := func(v int) bool {
			return v%2 == 0
		}
		in := slices.Values([]int{1, 3, 5})
		assert.Equal(t, mo.None[int](), FindQ(in, isEven))
	}
}
