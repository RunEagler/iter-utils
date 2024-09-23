package it

import "iter"

func Map[I, O any](it iter.Seq[I], f func(I) O) iter.Seq[O] {
	return func(yield func(O) bool) {
		for s := range it {
			if !yield(f(s)) {
				return
			}
		}
	}
}
