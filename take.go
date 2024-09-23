package it

func Take[I SeqT[V], V any](it I, n int) I {
	return func(yield func(V) bool) {
		num := n
		for v := range it {
			if num == 0 {
				return
			}
			if !yield(v) {
				return
			}
			num--
		}
	}
}
