package main

import (
	"fmt"

	"github.com/mifeis/Separable-Codes/combinations"
)

//Funció que busca totes les combinacions possibles en grup de 3 d'entre un grup de #WORDS
func main() {

	c := combinations.Init()
	totalCases := combinations.ListSeq(c)
	fmt.Println("Combinations found:", totalCases)
}
