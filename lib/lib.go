package lib

import "fmt"

const (
	WORDS = 8
	GROUP = 3
)

//Estructura que conté el grup de GROUP elements i un random id
//per saber de quina combinació es tracta i fer mes entendible l'arxiu resultant

type Combin struct {
	Id    int
	Group [GROUP]int
}

//funció que inicialitza i retorna l'array a combinar: {1,2,3,4,5,6,7,8,...}
func Init() []int {
	var c []int

	for i := 0; i < WORDS; i++ {
		c = append(c, i+1)
	}
	fmt.Println("Array:", c)
	return c
}

func RemoveSlice(original []int, g []int) []int {
	var remaining []int
	remaining = append(remaining, original[:]...)

	for i, elem := range g {
		RemoveIndex(remaining, elem-i-1)
	}

	return remaining[:WORDS-GROUP]
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
