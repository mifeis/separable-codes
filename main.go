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
	CASES = 3
)

//Funció que busca totes les combinacions possibles en grup de k d'entre un grup de WORDS
func main() {
	c := combinations.Init()
	total := 1
	for i := 0; i < CASES; i++ {
		log.Println("CASE", i)
		total *= getCase(c, i) / 2
		log.Println("Total:", total)
	}
	fmt.Println("Total cases for a code of "+strconv.Itoa(lib.WORDS)+" words: ", total)
}

//Funció que retorna els casos disjunts (List0) i els no disjunts (List1, List2)
//per a un array inicial de GROUP elements

func getCase(c []int, cas int) int {
	switch cas {
	case 0:
		return combinations.List0(c)
	case 1:
		return combinations.List1(c)
	case 2:
		return combinations.List2(c)
	default:
		return 0
	}
}
