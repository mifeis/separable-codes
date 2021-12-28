package main

import (
	"fmt"
	"strconv"

	"github.com/mifeis/Separable-Codes/lib_aux"
	"github.com/mifeis/Separable-Codes/lib_main"
)

//Main retorna el numero total de totes les combinacions possibles en grup de GROUP elements
//d'entre un array c de WORDS
func main() {
	c := lib_main.Init()
	allcases := lib_main.GetAllCases(c)

	fmt.Println("Total cases (", lib_aux.CASES, "types ) for a code of "+strconv.Itoa(lib_aux.WORDS)+" words: ", allcases)
}
