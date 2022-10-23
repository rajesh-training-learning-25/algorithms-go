// bogosort.go
// description: Implementation of the bogosort algorithm
// details:
// A simple bogosort algorithm implementation. [Wikipedia](https://en.wikipedia.org/wiki/Bogosort)
// author(s): [Focusucof](https://github.com/Focusucof)
// see sorts_test.go for testing

package sort

import (
	"github.com/TheAlgorithms/Go/constraints"
	"sort"
	"math/rand"
)

// shuffles the given int array
func shuffle[T constraints.Ordered](arr *[]int) {
	rand.Shuffle(len(*arr), func(i, j int) {
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i] // swap values
	})
}

// loops until array is sorted
func BogoSort[T constraints.Ordered](arr []int) []int {
	for {
		if sort.IntsAreSorted(arr) { // if array is sorted, break loop and return it
			return arr
		} else {
			shuffle(&arr) // else, shuffle array again
		}
	}
}