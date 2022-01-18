package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Valid per:
//casos disjunts i no disjunts
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}, ...
//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}, ...

func TestTeoric(t *testing.T) {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words is small")
	}
	var res int
	n := lib.WORDS
	k := lib.GROUP

	/* Combinations generates all of the combinations of k elements from a set of size n.
	 * The returned slice has length Binomial(n,k) and each inner slice has length k.
	 * n and k must be non-negative with n >= k, otherwise Combinations will panic.
	 * CombinationGenerator may alternatively be used to generate the combinations iteratively instead of collectively,
	 * or IndexToCombination for random access.
	 */
	list := combin.Combinations(n, k)
	//	fmt.Println("First group possible combinations:", len(list))

	for i := 0; i < lib.REPS; i++ {
		var total, all int
		fmt.Println(i, "elements repetitions:")
		for l := 0; l < lib.GROUP; l++ {
			if i > k-l {
				break
			}
			fmt.Println("max length:", lib.GROUP-l)
			combinations := len(combin.Combinations(n-k, (k-l)-i)) * len(combin.Combinations(k, i))
			if k-l == lib.GROUP {
				total = len(list) * combinations / 2
			} else {
				total = len(list) * combinations
			}
			fmt.Println(total)
			all += total
			res += total
		}
		fmt.Println("Total (", i, "elements repetitions )", all)
	}

	fmt.Println("Total cases for a code of "+strconv.Itoa(n)+" words in elements of", lib.GROUP, ":", res)
}
