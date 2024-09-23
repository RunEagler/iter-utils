package it

func Fold[I SeqT[V], V any](it I, sum V, f func(V, V) V) V {
	for v := range it {
		sum = f(sum, v)
	}

	return sum
}
