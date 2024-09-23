package it

import (
"fmt"
"slices"
"unique"
)

func main() {
	x := []string{"aa", "bb", "cc", "aa"}

	a := make([]unique.Handle[string], len(x))

	for i, c := range x {
		makeString(c)
		a[i] = unique.Make(c)
	}
	fmt.Println()

	numbers := slices.Values([]int{2, 4, 6, 7})

	for v := range Filter(numbers, func(i int) bool {
		return i%2 == 0
	}) {
		fmt.Println(v)
	}
}

func makeString(s string) {
	unique.Make(s)
}
