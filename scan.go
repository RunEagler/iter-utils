package it

import "iter"

func Scan[I SeqT[V], S, V any](it I, sum S, f func(S, V) S) iter.Seq[S] {
	return func(yield func(S) bool) {
		for v := range it {
			sum = f(sum, v)
			if !yield(sum) {
				return
			}
		}
	}
}
