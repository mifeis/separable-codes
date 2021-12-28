package lib_main

import (
	"fmt"

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

func GetAllCases(c []int) int {
	//	fmt.Println("First group possible combinations:", len(list))
	all := 1
	for i := 2; i < lib_aux.CASES; i++ {
		fmt.Println("Type", i+1)
		all *= getCasesByType(c, i) / 2
	}
	return all
}

//Funció que retorna els casos disjunts (List0) o els no disjunts (List1, List2)
//per a un array inicial de GROUP elements
func getCasesByType(c []int, t int) int {
	var total int
	arraymap := combinations.List(c, t)
	for _, combs := range arraymap {
		total += len(combs)
	}
	fmt.Println("Total cases:", total)
	return total
}
