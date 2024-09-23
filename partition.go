package it

func Partition[I SeqT[V], V any](it I, f func(V) bool) ([]V, []V) {
	var left []V
	var right []V
	for v := range it {
		if f(v) {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}
	return left, right
}
