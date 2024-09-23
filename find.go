package it

import "github.com/samber/mo"

func FindQ[I SeqT[V], V any](it I, f func(V) bool) mo.Option[V] {
	for v := range it {
		if f(v) {
			return mo.Some(v)
		}
	}
	return mo.None[V]()
}
