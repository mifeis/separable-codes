package lib_aux

import (
	"fmt"
)

const (

	/* casos totals (disjunts i no disjunts)
	 * Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
	 * Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}, ...
	 * Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}, ...
	 */

	CASES = 1

	WORDS = 8
	GROUP = 3
)

//Estructura que conté el grup de GROUP elements i un random id
//per saber de quina combinació es tracta i fer mes entendible l'arxiu resultant
type Combi struct {
	Id    int
	Group [3]int
	Value [3]int
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

//Says if the two arrays are separable or not
func Separable(group1 [GROUP]int, group2 [GROUP]int) bool {
	first := make(map[int]int)
	second := make(map[int]int)

	fmt.Print("Comparing ", group1, " with ", group2)
	for _, v := range group1 {
		first[v] = v
	}

	for _, v := range group2 {
		second[v] = v
	}

	//Comprobar casos especials en que lengths iguals:
	//No separables-> (0,0,0) i (0,0,0), (1,1,1) i (1,1,1)
	//Separables-> (0,0,0) i (1,1,1), (1,1,1) i (0,0,0)
	if len(first) == len(second) {
		_, z1 := first[0]
		_, z2 := second[0]
		if (z1 && z2) || (!z1 && !z2) {
			fmt.Println(" -> No separables")
			return false
		}
		fmt.Println(" -> Separables")
		return true
	} else {
		fmt.Println(" -> Separables:", len(first) != len(second))
		return len(first) != len(second)
	}
}

//Retorna totes les combinacions de valors (0/1) d'un array de GROUP elements
func GetDefaultValues() [][3]int {
	var slice [GROUP]int
	var values [][GROUP]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				slice = [GROUP]int{i % 2, j % 2, k % 2}
				values = append(values, slice)
			}
		}
	}
	fmt.Println("Array possible binari values:", values)
	return values
}
