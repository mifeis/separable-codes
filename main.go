package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
	"github.com/mifeis/Separable-Codes/test"
)

//Main retorna el numero total de totes les combinacions possibles en grup de GROUP elements
//d'entre un array c de WORDS
func main() {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words is small")
	}
	fmt.Println("Intro test for " + strconv.Itoa(lib.WORDS) + "x" + strconv.Itoa(lib.GROUP) + ":")
	fmt.Println("1.\tCombinations\n2.\tFavourable and unfavourable cases\n3.\tDependence\n4.\tTheoretical")

	var t int
	fmt.Scanf("%d", &t)
	initial := lib.Init(0, lib.WORDS)

	switch t {
	case 1:
		combinations.GetCombinations(initial)
	case 2:
		test.GetFavourables(initial)
	case 3:
		test.GetDependences(initial)
	case 4:
		test.Theoretical(lib.WORDS, lib.GROUP)
	}
}
