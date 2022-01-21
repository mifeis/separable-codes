package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//Test que retorna el numero de combinacions dependents

func TestDeps(t *testing.T) {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words must be smaller than 2 * group elements")
	}
	initial := lib.Init(0, lib.WORDS)
	fmt.Println("\n...Getting dependence")
	getDeps(initial)
}

func getDeps(initial []int) {
	var array []map[int][]lib.Map
	arraypairs := make(map[string][][]int)
	var key string

	for i := 0; i < lib.REPS; i++ {
		arraymap := combinations.List(initial, i)
		array = append(array, lib.Sort(arraymap))
	}

	for reps, arraymaps := range array {
		//Set a combination
		for k1, am := range arraymaps {
			for _, m := range am {
				for k2, ss := range m.Seconds {
					for _, s := range ss {
						var first, second, group []int
						key = strconv.Itoa(k1) + strconv.Itoa(k2)
						first = append(first, m.First...)
						second = append(second, s...)
						group = append(first, second...)
						if k1 == k2 {
							if !inversAlreadyInArray(arraypairs[key], group, reps) {
								//la key pot ser array de reps + group1xgroup2
								arraypairs[key] = append(arraypairs[key], group)
							}
						} else {
							arraypairs[key] = append(arraypairs[key], group)
						}
					}
				}
			}
		}
	}
	//	fmt.Println(arraypairs)
	for k, arraypair := range arraypairs {
		for o, v := range arraypair {
			var deps int

			for key, arrayToCompare := range arraypairs {
				var init, fin int
				//				fmt.Println(arrayToCompare)
				if k == key {
					init = o + 1
					fin = 1
				} else {
					init, fin = 0, 0
				}
				for i := 0; i < len(arrayToCompare)-fin; i++ {
					index := (init + i) % (len(arrayToCompare))
					if lib.Dependent(v, arrayToCompare[index]) {
						deps++
					}
				}
			}
			lib.LogDeps(k, v, deps)
		}
	}

}

func inversAlreadyInArray(arraypairs [][]int, pair []int, reps int) bool {
	var length int
	//	fmt.Print(pair)

	for _, p := range arraypairs {
		length = compareArrays(p, pair)
		//		fmt.Print(" len:", length)
		if (len(pair) + reps*2) == length {
			//			fmt.Println(" repe")
			return true
		}
	}
	//	fmt.Println()

	return false
}

/*
[3 5 7 3 5 2]
repe [3 5 7 3 7 4]

repe [3 5 7 5 7 4]
repe [3 5 7 5 7 6]
[3 5 7 3 7 2]
repe [3 5 7 3 5 8]
repe [3 5 7 5 7 8]
repe [3 5 7 3 7 1]
repe [3 5 7 3 5 1]
repe [3 5 7 3 7 6]
repe [3 5 7 3 5 6]
[3 5 7 5 7 2]
repe [3 5 7 3 5 4]
repe [3 5 7 3 7 8]
repe [3 5 7 5 7 1]

map[33:[[3 5 7 3 5 2] [3 5 7 3 7 2] [3 5 7 5 7 2]]]
*/
func compareArrays(array1 []int, array2 []int) int {
	var l int
	//	fmt.Println("comparing len:", array1, array1)
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v1 == v2 {
				l++
			}
		}
	}
	return l
}

//		[3 5 7 3 5 2]
//repe 	[3 5 7 3 7 4]
