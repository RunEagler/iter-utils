package it

func Count[I SeqT[V], V any](it I) int {
	count := 0
	for _ = range it {
		count++
	}
	return count
}
