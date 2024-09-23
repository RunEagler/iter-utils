package it

import (
	"iter"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"

	"iter-utils/internal/helper"
)

func TestFactorial(t *testing.T) {
	expect := []int{1, 2, 6, 24, 120, 720}
	next, stop := iter.Pull(Factorial[int]())
	defer stop()
	for i := range len(expect) {
		v, ok := next()
		assert.True(t, ok)
		assert.Equal(t, expect[i], v)
	}
}

func TestSquired(t *testing.T) {
	expect := []int{1, 4, 9, 16, 25, 36}
	next, stop := iter.Pull(Squired[int]())
	defer stop()
	for i := range len(expect) {
		v, ok := next()
		assert.True(t, ok)
		assert.Equal(t, expect[i], v)
	}
}

func TestPrime(t *testing.T) {
	{
		helper.SliceEqual(t, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}, slices.Collect(Take(Prime[int](100), 10)))
	}
	{
		helper.SliceEqual(t, []int{}, slices.Collect(Take(Prime[int](100), 0)))
	}
	{
		helper.SliceEqual(t, []int{2, 3, 5, 7}, slices.Collect(Take(Prime[int](10), 100)))
	}
}
