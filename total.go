package main

import (
	"fmt"
	"strconv"

	"github.com/mifeis/Separable-Codes/lib_main"
)

//Main retorna el numero total de totes les combinacions possibles en grup de GROUP elements
//d'entre un array c de WORDS
func main() {
	c := lib_main.Init()
	cases := lib_main.GetTotal(c)

	fmt.Println("Total cases for a code of "+strconv.Itoa(lib_main.WORDS)+" words: ", cases)
}
