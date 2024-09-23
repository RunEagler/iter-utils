package it

import (
	"github.com/samber/mo"
)

func NthQ[I SeqT[V], V any](it I, n int) mo.Option[V] {
	i := 0
	for s := range it {
		if i < n {
			i++
			continue
		}
		return mo.Some(s)
	}
	return mo.None[V]()
}

func LastQ[I SeqT[V], V any](it I) mo.Option[V] {
	var l V
	i := 0
	for v := range it {
		l = v
		i++
	}
	if i == 0 {
		return mo.None[V]()
	}
	return mo.Some[V](l)
}
