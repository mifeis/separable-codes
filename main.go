package main

import (
	"fmt"
	"strconv"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//Main retorna el numero total de totes les combinacions possibles en grup de GROUP elements
//d'entre un array c de WORDS
func main() {
	initial := lib.Init()
	allcases := getAllCases(initial)

	fmt.Println("Total cases (", lib.CASES, "types ) for a code of "+strconv.Itoa(lib.WORDS)+" words: ", allcases)
}

func getAllCases(c []int) int {
	//	fmt.Println("First group possible combinations:", len(list))
	all := 1
	for i := 0; i < lib.CASES; i++ {
		//Total pren el valor dels casos disjunts ó no disjunts
		//per a un array inicial de GROUP elements
		var total int

		lib.LogType(i + 1)

		arraymap := combinations.List(c, i+1)

		fmt.Println("Combinations:")
		for g, combs := range arraymap {
			total += lib.LogCombinations(g, combs)
		}

		fmt.Println("Total cases:", total/2)
		//ó *
		all += total / 2
	}
	return all
}
