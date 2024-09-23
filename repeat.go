package it

func Repeat[I SeqT[V], V any](v V) I {
	return func(yield func(V) bool) {
		for {
			if !yield(v) {
				return
			}
		}
	}
}
