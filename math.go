package it

import (
	"iter"

	"golang.org/x/exp/constraints"
)

func Factorial[V constraints.Signed]() iter.Seq[V] {
	return func(yield func(V) bool) {
		i := V(1)
		sum := i
		for {
			if !yield(sum) {
				return
			}
			sum *= i + 1
			i++
		}
	}
}

func Prime[V constraints.Signed](n int) iter.Seq[V] {
	primeTable := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primeTable[i] = true
	}
	return func(yield func(V) bool) {
		for p := 2; p <= n; p++ {
			if !primeTable[p] {
				continue
			}
			if !yield(V(p)) {
				return
			}
			for q := p * 2; q <= n; q += p {
				primeTable[q] = false
			}
		}
	}
}

func Squired[V constraints.Signed]() iter.Seq[V] {
	return func(yield func(V) bool) {
		v := V(1)
		for {
			if !yield(v * v) {
				return
			}
			v++
		}
	}
}
