package test

import (
	"fmt"

	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Valid per:
//casos disjunts i no disjunts
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}, ...
//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}, ...

func Theoretical(n int, k int) {

	/* Combinations generates all of the combinations of k elements from a set of size n.
	 * The returned slice has length Binomial(n,k) and each inner slice has length k.
	 * n and k must be non-negative with n >= k, otherwise Combinations will panic.
	 * CombinationGenerator may alternatively be used to generate the combinations iteratively instead of collectively,
	 * or IndexToCombination for random access.
	 */

	for i := 0; i < lib.REPS; i++ {
		fmt.Println(i, "elements REPETITIONS:")

		for l1 := lib.GROUP; l1 > i; l1-- {
			var total int

			for l2 := l1; l2 > 0; l2-- {
				var all int

				if i > l2 && l2 < l1 {
					break
				}
				fmt.Print(l1, " x ", l2, " -> ")
				i1 := lib.GROUP - l1
				i2 := lib.GROUP - l2
				list := combin.Combinations(n, k-i1)
				combinations := len(combin.Combinations(n-(k-i1), k-i2-i)) * len(combin.Combinations(k-i1, i))
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
