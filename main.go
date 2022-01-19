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

	fmt.Println("Total combinations (til", lib.REPS-1, "elements repetitions) for a code of "+strconv.Itoa(lib.WORDS)+" words:", all)
}

//Canviar
func getAllCombinations(c []int) int {
	var all int

	for reps := 0; reps < lib.REPS; reps++ {
		arraymap := combinations.List(c, reps)
		all += lib.LogCombinations(arraymap, reps)
	}
	return all
}
