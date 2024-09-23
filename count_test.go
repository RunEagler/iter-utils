package it

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	{
		assert.Equal(t, 5, Count(Range(0, 5)))
	}
	{
		assert.Equal(t, 0, Count(slices.Values([]int{})))
	}
}
