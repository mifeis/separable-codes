package main

import (
	"fmt"
	"strconv"

	"github.com/mifeis/Separable-Codes/combinations"
)

const (
	CASES = 1
)

func getCase(c []int, cas int) int {
	switch cas {
	case 0:
		return combinations.ListSeq(c)
	case 1:
		return combinations.ListSeq1(c)
	case 2:
		return combinations.ListSeq2(c)
	default:
		return 0
	}
}

//Funció que busca totes les combinacions possibles en grup de 3 d'entre un grup de #WORDS
func main() {
	c := combinations.Init()
	total := 1
	for i := 0; i < CASES; i++ {
		total *= total * getCase(c, i) / 2
	}
	fmt.Println("Total cases for a code of "+strconv.Itoa(combinations.WORDS)+" words: ", total)
}
