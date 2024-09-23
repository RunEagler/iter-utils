package it

import (
	"slices"
	"testing"

	"iter-utils/internal/helper"
)

func TestEnumerate(t *testing.T) {
	{
		in := slices.Values([]string{"aa", "bb", "cc"})
		expect := []IndexedValue[string]{
			{0, "aa"}, {1, "bb"}, {2, "cc"},
		}
		helper.SliceEqual(t, expect, slices.Collect(Enumerate(in)))
	}
	{
		in := slices.Values([]string{})
		var expect []IndexedValue[string]
		helper.SliceEqual(t, expect, slices.Collect(Enumerate(in)))
	}
}
