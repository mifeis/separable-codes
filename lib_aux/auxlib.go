package lib_aux

import "github.com/mifeis/Separable-Codes/lib_main"

//Estructura que conté el grup de GROUP elements i un random id
//per saber de quina combinació es tracta i fer mes entendible l'arxiu resultant

type Combin struct {
	Id    int
	Group [lib_main.GROUP]int
}

func RemoveSlice(original []int, g []int) []int {
	var remaining []int
	remaining = append(remaining, original[:]...)

	for i, elem := range g {
		RemoveIndex(remaining, elem-i-1)
	}

	return remaining[:lib_main.WORDS-lib_main.GROUP]
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
