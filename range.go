package it

import "iter"

func Range(start, end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := start; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func Numbers(end int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := 0; i < end; i++ {
			if !yield(i) {
				return
			}
		}
	}
}
