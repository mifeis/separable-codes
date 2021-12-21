package lib_main

import (
	"fmt"
	"log"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib_aux"
)

//funció que inicialitza i retorna l'array a combinar: {1,2,3,4,5,6,7,8,...}
func Init() []int {
	var c []int

	for i := 0; i < lib_aux.WORDS; i++ {
		c = append(c, i+1)
	}
	fmt.Println("Array:", c)
	return c
}

func GetTotal(c []int) int {
	cases := 1
	for i := 0; i < lib_aux.CASES; i++ {
		log.Println("CASE", i)
		cases *= getCases(c, i) / 2
		log.Println("Total cases:", cases)
	}
	return cases
}

//Funció que retorna els casos disjunts (List0) i els no disjunts (List1, List2)
//per a un array inicial de GROUP elements
func getCases(c []int, cas int) int {
	var total int
	var arraymap map[lib_aux.Combi][]lib_aux.Combi
	switch cas {
	case 0:
		arraymap = combinations.List0(c)
	case 1:
		arraymap = combinations.List1(c)
	case 2:
		arraymap = combinations.List2(c)
	}
	for _, combs := range arraymap {
		total += len(combs)
	}
	return total
}
