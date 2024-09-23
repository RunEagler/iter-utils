package it

import (
	"slices"
	"testing"

	"github.com/samber/mo"
	"github.com/stretchr/testify/assert"
)

func TestNthQ(t *testing.T) {
	{
		in := slices.Values([]int{0, 2, 4, 6, 8, 10})
		assert.Equal(t, mo.Some(6), NthQ(in, 3))
	}
	{
		in := slices.Values([]int{})
		assert.Equal(t, mo.None[int](), NthQ(in, 5))
	}
}

func TestLastQ(t *testing.T) {
	{
		in := slices.Values([]int{0, 2, 4, 6, 8, 10})
		assert.Equal(t, mo.Some(10), LastQ(in))
	}
	{
		in := slices.Values([]int{})
		assert.Equal(t, mo.None[int](), LastQ(in))
	}
}
