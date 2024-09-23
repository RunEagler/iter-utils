package it

import "github.com/samber/mo"

func Position[I SeqT[V], V any](it I, f func(V) bool) int {
	i := 0
	for v := range it {
		if f(v) {
			return i
		}
		i++
	}
	return -1
}

func PositionQ[I SeqT[V], V any](it I, f func(V) bool) mo.Option[int] {
	i := 0
	for v := range it {
		if f(v) {
			return mo.Some(i)
		}
		i++
	}
	return mo.None[int]()
}
