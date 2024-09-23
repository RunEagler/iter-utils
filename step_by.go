package it

func StepBy[I SeqT[V], V any](it I, n int) I {
	return func(yield func(V) bool) {
		i := 0
		for v := range it {
			if i%n == 0 {
				if !yield(v) {
					return
				}
			}
			i++
		}
	}
}
