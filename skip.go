package it

import "iter"

func Skip[I SeqT[V], V any](it I, n int) I {
	return func(yield func(V) bool) {
		next, stop := iter.Pull(iter.Seq[V](it))
		defer stop()
		for range n {
			_, ok := next()
			if !ok {
				return
			}
		}
		for {
			v, ok := next()
			if !ok || !yield(v) {
				return
			}
		}
	}
}
