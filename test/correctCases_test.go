package main

import (
	"fmt"
	"strconv"
	"testing"

	"gonum.org/v1/gonum/stat/combin"
)

//casos no disjunts
const (
	CASES = 3
)

func TestRead(t *testing.T) {

	//	_, err := model.ReadConf(READ_SETTINGS)
	//	if err {
	//		t.Errorf("FAILED")
	//	}

	total := 1
	n := 8
	k := 3

	/* Combinations generates all of the combinations of k elements from a set of size n.
	 * The returned slice has length Binomial(n,k) and each inner slice has length k.
	 * n and k must be non-negative with n >= k, otherwise Combinations will panic.
	 * CombinationGenerator may alternatively be used to generate the combinations iteratively instead of collectively,
	 * or IndexToCombination for random access.
	 */

	list := combin.Combinations(n, k)
	fmt.Println("len first groups", len(list))
	for i := 2; i < CASES; i++ {
		cas := len(combin.Combinations(n-k, k-i)) * len(combin.Combinations(k, i))
		fmt.Println("len combis no disjuntes (cas 1)", cas)
		total *= len(list) * cas / 2
	}

	fmt.Println("Total cases for a code of "+strconv.Itoa(n)+" words: ", total)
}
