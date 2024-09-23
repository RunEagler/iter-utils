package it

func Cycle[I SeqT[V], V any](s I) I {
	return func(yield func(V) bool) {
		for {
			for v := range s {
				if !yield(v) {
					return
				}
			}
		}
	}
}
