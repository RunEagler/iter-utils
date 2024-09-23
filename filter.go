package it

func Filter[I SeqT[V], V any](s I, f func(V) bool) I {
	return func(yield func(V) bool) {
		for v := range s {
			if f(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

func All[I SeqT[V], V any](s I, f func(V) bool) bool {
	for v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func Any[I SeqT[V], V any](s I, f func(V) bool) bool {
	for v := range s {
		if f(v) {
			return true
		}
	}
	return false
}
