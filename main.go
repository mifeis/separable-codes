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
		log.Fatal("num of words is small")
	}
	initial := lib.Init(0, lib.WORDS)
	all := getAllCombinations(initial)

	fmt.Println("Total cases (", lib.REPS, "types ) for a code of "+strconv.Itoa(lib.WORDS)+" words:", all)
}

//Canviar
func getAllCombinations(c []int) int {
	var all int

	for k := 0; k < lib.GROUP; k++ {
		for reps := 0; reps < lib.REPS; reps++ {
			lib.LogTipus(reps, lib.GROUP-k)
			arraymap := combinations.List(c, k, reps)
			all += lib.LogCombinations(arraymap) / 2
		}
	}
	return all
}
