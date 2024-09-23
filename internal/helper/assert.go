package helper

import (
	"slices"
	"testing"
)

func SliceEqual[V comparable](t *testing.T, a, b []V) {
	t.Helper()
	if !slices.Equal(a, b) {
		t.Errorf("expected `%v` to equal `%v`", a, b)
	}
}
