package it

func ForEach[I SeqT[V], V any](it I, f func(V)) {
	for v := range it {
		f(v)
	}
}
