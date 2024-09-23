package it

import (
	"iter"

	"github.com/samber/mo"
)

type SeqT[V any] interface {
	~func(yield func(V) bool)
}

type Iterator[V any] iter.Seq[V]

func Slice[V any](list []V) Iterator[V] {
	return func(yield func(V) bool) {
		for _, v := range list {
			if !yield(v) {
				return
			}
		}
	}
}

func (it Iterator[V]) Take(n int) Iterator[V] {
	return Take(it, n)
}

func (it Iterator[V]) Count() int {
	return Count(it)
}

func (it Iterator[V]) Last() mo.Option[V] { return LastQ(it) }

func (it Iterator[V]) Nth(i int) mo.Option[V] {
	return NthQ(it, i)
}

func (it Iterator[V]) Cycle() Iterator[V] { return Cycle(it) }

func (it Iterator[V]) Find(f func(V) bool) mo.Option[V] {
	return FindQ(it, f)
}

func (it Iterator[V]) All(f func(V) bool) bool { return All(it, f) }

func (it Iterator[V]) Any(f func(V) bool) bool { return Any(it, f) }

func (it Iterator[V]) Filter(f func(V) bool) { Filter(it, f) }

func (it Iterator[V]) ForEach(f func(V)) { ForEach(it, f) }

func (it Iterator[V]) Partition(f func(V) bool) ([]V, []V) { return Partition(it, f) }

func (it Iterator[V]) Position(f func(V) bool) mo.Option[int] { return PositionQ(it, f) }

func (it Iterator[V]) Skip(i int) Iterator[V] { return Skip(it, i) }

func (it Iterator[V]) StepBy(i int) Iterator[V] { return StepBy(it, i) }

func (it Iterator[V]) Chain(it2 Iterator[V]) Iterator[V] { return Chain(it, it2) }

func (it Iterator[V]) ToCollect() []V {
	var ret []V
	for s := range it {
		ret = append(ret, s)
	}
	return ret
}
