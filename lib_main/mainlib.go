package lib_main

import (
	"fmt"
	"log"

	"github.com/mifeis/Separable-Codes/combinations"
)

const (

	/* casos totals (disjunts i no disjunts)
	 * Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
	 * Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}, ...
	 * Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}, ...
	 */

	CASES = 3

	WORDS = 8
	GROUP = 3
)

//funció que inicialitza i retorna l'array a combinar: {1,2,3,4,5,6,7,8,...}
func Init() []int {
	var c []int

	for i := 0; i < WORDS; i++ {
		c = append(c, i+1)
	}
	fmt.Println("Array:", c)
	return c
}

func GetTotal(c []int) int {
	cases := 1
	for i := 0; i < CASES; i++ {
		log.Println("CASE", i)
		cases *= getCases(c, i) / 2
		log.Println("Total cases:", cases)
	}
	return cases
}

//Funció que retorna els casos disjunts (List0) i els no disjunts (List1, List2)
//per a un array inicial de GROUP elements
func getCases(c []int, cas int) int {
	var total int
	var arraymap map[[GROUP]int][][GROUP]int
	switch cas {
	case 0:
		arraymap = combinations.List0(c)
	case 1:
		arraymap = combinations.List1(c)
	case 2:
		arraymap = combinations.List2(c)
	default:
		return 0
	}
	for _, combs := range arraymap {
		total += len(combs)
	}
	return total
}
