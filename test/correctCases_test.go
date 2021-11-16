package main

import (
	"fmt"
	"strconv"
	"testing"

	"gonum.org/v1/gonum/stat/combin"
)

const (
	CASES = 1
)

func getCase(c int, n int, k int) int {
	switch c {
	case 1:
		return len(combin.Combinations(n-k, k))
	case 2:
		return len(combin.Combinations(n-k, k-1))
	case 3:
		return len(combin.Combinations(n-k, k-2))
	default:
		return 0
	}
}

func TestRead(t *testing.T) {

	//	_, err := model.ReadConf(READ_SETTINGS)
	//	if err {
	//		t.Errorf("FAILED")
	//	}

	totalCases := 1
	n := 16
	k := 3

	/* Combinations generates all of the combinations of k elements from a set of size n.
	 * The returned slice has length Binomial(n,k) and each inner slice has length k.
	 * n and k must be non-negative with n >= k, otherwise Combinations will panic.
	 * CombinationGenerator may alternatively be used to generate the combinations iteratively instead of collectively,
	 * or IndexToCombination for random access.
	 */

	list := combin.Combinations(n, k)
	//	log.Println(list)
	for i := 0; i < CASES; i++ {
		totalCases *= len(list) * getCase(i+1, n, k) / 2
	}

	fmt.Println("Total cases for a code of "+strconv.Itoa(n)+" words: ", totalCases)
}
