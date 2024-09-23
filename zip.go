package it

import "iter"

type Tuple[V1, V2 any] struct {
	V1 V1
	V2 V2
}

func Zip[V1, V2 any](it1 iter.Seq[V1], it2 iter.Seq[V2]) iter.Seq[Tuple[V1, V2]] {
	return func(yield func(t Tuple[V1, V2]) bool) {
		next1, stop1 := iter.Pull(it1)
		defer stop1()
		next2, stop2 := iter.Pull(it2)
		defer stop2()
		for {
			v1, ok1 := next1()
			v2, ok2 := next2()
			if !ok1 || !ok2 {
				return
			}
			if !yield(Tuple[V1, V2]{v1, v2}) {
				break
			}
		}
	}
}
