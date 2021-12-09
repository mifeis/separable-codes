package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//casos totals (disjunts i no disjunts)
//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}, ...
//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}, ...

const (
	CASES = 1
)

//Main retorna el numero total de totes les combinacions possibles en grup de GROUP elements
//d'entre un array c de WORDS
func main() {
	c := lib.Init()
	total := getTotal(c)

	fmt.Println("Total cases for a code of "+strconv.Itoa(lib.WORDS)+" words: ", total)
}

func getTotal(c []int) int {
	total := 1
	for i := 0; i < CASES; i++ {
		log.Println("CASE", i)
		total *= getCase(c, i) / 2
		log.Println("Total:", total)
	}
	return total
}

//Funció que retorna els casos disjunts (List0) i els no disjunts (List1, List2)
//per a un array inicial de GROUP elements
func getCase(c []int, cas int) int {
	switch cas {
	case 0:
		combs := combinations.List0(c)
		log.Println(combs)
		return len(combinations.List0(c))
	case 1:
		return len(combinations.List1(c))
	case 2:
		return len(combinations.List2(c))
	default:
		return 0
	}
}
