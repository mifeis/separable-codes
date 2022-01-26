package test

import (
	"fmt"

	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

// Cases counted:
// Disjoint 3x3, 3x2, 3x1, 2x2, 2x1 and 1x1
// For example: {1,2,3}|{4,5,6}, {1,2,3}|{5,6}, {4,7}|{1,2}
// Non disjoint (1 element repetition) 3x3, 3x2, 3x1, 2x2, 2x1
// For example: {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {2,3}|{3}
// Non disjoint (2 elements repetitions) 3x3, 3x2
// For example: {1,2,3}|{1,2,5}, {1,2,3}|{1,3}, {1,2,3}|{2,3}

// The theoretical method gives the total combinations number
// for each case (disjoint and non disjoint). This means making
// a loop that gives the result for 0,1 and 2 elements
// repetitions
func Theoretical(n int, k int) {

	/* Combinations generates all of the combinations of k elements from a set of size n.
	 * The returned slice has length Binomial(n,k) and each inner slice has length k.
	 * n and k must be non-negative with n >= k, otherwise Combinations will panic.
	 * CombinationGenerator may alternatively be used to generate the combinations iteratively instead of collectively,
	 * or IndexToCombination for random access.
	 */

	for i := 0; i < lib.REPS; i++ {
		fmt.Println(i, "elements REPETITIONS:")

		// The following loop fixes the size of the first group
		// given a certain number of repetitions. It must be
		// greater than variable i then.
		for l1 := lib.GROUP; l1 > i; l1-- {
			var total int // Total combinations for each first group length

			for l2 := l1; l2 > 0; l2-- {
				var all int // All cases for a fixed dimension type, such as 3x2 or 2x2

				// Avoid cases like 2x1 with 2 repetitions
				if i > l2 && l2 < l1 {
					break
				}
				fmt.Print(l1, " x ", l2, " -> ")

				// Use of the combin library to apply the theory
				i1 := lib.GROUP - l1
				i2 := lib.GROUP - l2
				// First groups combinations
				list := combin.Combinations(n, k-i1)
				// First * seconds combinations
				combinations := len(combin.Combinations(n-(k-i1), k-i2-i)) * len(combin.Combinations(k-i1, i))

				// Take into account that cases might be duplicated.
				// For example: {1 2 3}{4 5 6} and {4 5 6}{1 2 3}
				if l2 == l1 && l1 != i {
					all = len(list) * combinations / 2
				} else {
					all = len(list) * combinations
				}

				fmt.Println(all)
				total += all
			}
			fmt.Println("Total:\t", total)
		}
	}
}
