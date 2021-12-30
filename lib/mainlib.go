package lib

import (
	"fmt"
)

//funció que inicialitza i retorna l'array a combinar: {1,2,3,4,5,6,7,8,...}
func Init() []int {
	var initial []int

	for i := 0; i < WORDS; i++ {
		initial = append(initial, i+1)
	}
	fmt.Println("Initial array:", initial)
	return initial
}
