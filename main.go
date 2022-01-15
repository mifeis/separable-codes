package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//Main retorna el numero total de totes les combinacions possibles en grup de GROUP elements
//d'entre un array c de WORDS
func main() {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words must be smaller than 2 * group elements")
	}
	initial := lib.Init(0, lib.WORDS)
	all := getAllCombinations(initial)

	fmt.Println("Total cases (", lib.REPS, "types ) for a code of "+strconv.Itoa(lib.WORDS)+" words:", all)
}

//Canviar
func getAllCombinations(c []int) int {
	//	fmt.Println("First group possible combinations:", len(list))
	var all int
	for i := 1; i < lib.REPS; i++ {
		//Total pren el valor dels casos disjunts ó no disjunts
		//per a un array inicial de GROUP elements
		var total int

		lib.LogTipus(i)
		arraymap := combinations.List(c, i)

		fmt.Println("Combinations:")
		for _, m := range arraymap {
			total += lib.LogCombinations(m)
		}
		fmt.Println("Total cases:", total/2)
		log.Println(total / 2)

		all += total / 2
	}
	return all
}
