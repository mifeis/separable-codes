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
	fmt.Println("Initial array:", c)
	return c
}

func GetAllCases(c []int) int {
	//	fmt.Println("First group possible combinations:", len(list))
	all := 1
	for i := 0; i < lib_aux.CASES; i++ {
		var total int

		fmt.Println("\n------------------------------------------")
		fmt.Println("\t\t\t\tType", i+1)
		fmt.Println("------------------------------------------")

		arraymap := combinations.List(c, i+1)
		//Total pren el valor dels casos disjunts ó no disjunts per a un array inicial de GROUP elements
		for _, combs := range arraymap {
			total += len(combs)
		}

		fmt.Println("Total cases:", total)
		all *= total / 2
	}
	return all
}
