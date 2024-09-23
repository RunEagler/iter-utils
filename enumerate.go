package it

import "iter"

type IndexedValue[V any] struct {
	Index int
	Value V
}

func Enumerate[I SeqT[V], V any](it I) iter.Seq[IndexedValue[V]] {
	return func(yield func(v IndexedValue[V]) bool) {
		i := 0
		for v := range it {
			if !yield(IndexedValue[V]{i, v}) {
				return
			}
			i++
		}
	}
}
