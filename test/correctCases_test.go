package main

import (
	"fmt"
	"strconv"
	"testing"

	"gonum.org/v1/gonum/stat/combin"
)

//casos no disjunts
const (
	CASES = 1
)

func TestRead(t *testing.T) {

	//	_, err := model.ReadConf(READ_SETTINGS)
	//	if err {
	//		t.Errorf("FAILED")
	//	}

	totalCases := 1
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
	for i := 1; i <= CASES; i++ {
		//		disjunt := len(combin.Combinations(n-k, k))
		nodisjunt := len(combin.Combinations(n-k, k-i)) * k
		fmt.Println("len combis no disjuntes", nodisjunt)
		totalCases *= len(list) * nodisjunt / 2
	}

	fmt.Println("Total cases for a code of "+strconv.Itoa(n)+" words: ", totalCases)
}
