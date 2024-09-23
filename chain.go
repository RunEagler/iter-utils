package it

func Chain[I SeqT[V], V any](it1 I, it2 I) I {
	return func(yield func(V) bool) {
		for _, it := range []I{it1, it2} {
			for v := range it {
				if !yield(v) {
					return
				}
			}
		}
	}
}
