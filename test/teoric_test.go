package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//casos disjunts i no disjunts
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}, ...
//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}, ...

func TestTeoric(t *testing.T) {

	total := 1
	all := 1
	n := lib.WORDS
	k := lib.GROUP

	/* Combinations generates all of the combinations of k elements from a set of size n.
	 * The returned slice has length Binomial(n,k) and each inner slice has length k.
	 * n and k must be non-negative with n >= k, otherwise Combinations will panic.
	 * CombinationGenerator may alternatively be used to generate the combinations iteratively instead of collectively,
	 * or IndexToCombination for random access.
	 */
	list := combin.Combinations(n, k)
	fmt.Println("First group possible combinations:", len(list))

	for i := 0; i < lib.REPS; i++ {
		combinations := len(combin.Combinations(n-k, k-i)) * len(combin.Combinations(k, i))
		fmt.Println("Type", i+1, "-> Combinations:", combinations)
		total = len(list) * combinations / 2
		fmt.Println("Total cases:", total)
		all += total

	}
	fmt.Println("Total cases (", lib.REPS, "types ) for a code of "+strconv.Itoa(n)+" words: ", all)
}
