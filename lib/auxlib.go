package lib

import (
	"fmt"
	"math"
)

const (

	/* CASOS TOTALS
	 * disjunts:
	 * Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
	 * no disjunts:
	 * Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}, ...
	 * Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}, ...
	 * inclomplerts:
	 * Tipus {1,2,3}|{4},
	 * Tipus {1,2,3}|{4,5}
	 */

	REPS = 2 //2*GROUP - 1 //disjunts+no disjunts+ inclomplerts (lib.GROUP+lib.GROUP-1)
	//- when (WORDS<2GROUP)

	WORDS = 6
	GROUP = 3
)

//Estructura que conté el primer element d'un tamany desde GROUP elements fins a 1
//i un segon array que consta de totes les combinacions possibles per aquest primer grup
type Map struct {
	First   []int
	Seconds [][]int
}

//funció que inicialitza i retorna l'array a combinar: {1,2,3,4,5,6,7,8,...}
func Init(init int, len int) []int {
	var initial []int

	for i := init; i < len; i++ {
		initial = append(initial, i+1)
	}
	//	fmt.Println("Initial array:", initial)
	return initial
}

//Removes the slice from the original
func RemoveSlice(original []int, slice []int) []int {
	var remaining []int
	remaining = append(remaining, original[:]...)

	for i, elem := range slice {
		RemoveIndex(remaining, elem-i-1)
	}

	//	fmt.Println("Remaining array from", slice, ":", remaining)
	return remaining[:WORDS-GROUP]
}

//Removes the index from the slice
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

//Estructura que defineix valors (0/1) per a un grup d'elements
type Code struct {
	Row    []int
	Values []int
}

//Retorna totes les combinacions de valors (0/1) d'un array de GROUP elements
func GetDefaultValues() [][]int {
	var slice []int
	var values [][]int
	//s'haura de passar len per argument depenent del GROUP del moment (3,2,1...)
	len := int(math.Exp2(GROUP))
	for t := 0; t < len; t++ {
		for i := range slice {
			ijk := t / (len / int(math.Exp2(float64(i+1))))
			slice[i] = ijk % 2
		}
		values = append(values, slice)
	}
	fmt.Println("Possible binari values for a group of", GROUP, "elements:", values)
	return values
}

//Function assign values from first to second if the columns are the same
//If not assignes the number 2 to the leaving columnes
func SetValues(first Code, second *Code) {
	//comprobar
	for i := 0; i < GROUP; i++ {
		second.Values[i] = 2
	}
	//	if GROUP == 4 {
	//		second.Value[GROUP-1] = 2
	//	}
	for m, v1 := range first.Row {
		for n, v2 := range second.Row {
			if v1 == v2 {
				second.Values[m] = first.Values[n]
			}
		}
	}
}

//Says if the two arrays are separable or not
func Separable(group1 []int, group2 []int) bool {
	first := make(map[int]int)
	second := make(map[int]int)

	fmt.Print("Comparing ", group1, " with ", group2)
	for _, v := range group1 {
		first[v] = v
	}

	for _, v := range group2 {
		second[v] = v
	}

	var isSep bool
	if len(first) == len(second) {
		//Comprobar casos especials en que lengths iguals:
		//No separables-> (0,0,0) i (0,0,0), (1,1,1) i (1,1,1)
		//Separables-> (0,0,0) i (1,1,1), (1,1,1) i (0,0,0)

		_, z1 := first[0]
		_, z2 := second[0]
		if (z1 && z2) || (!z1 && !z2) {
		} else {
			isSep = true
		}
	} else {
		//Altres casos
		//No separables-> (0,0,1) i (1,0,0), (0,0,0) i (0,0,0)
		//Separables-> (0,0,0) i (1,0,0), (1,1,1) i (0,1,0)

		isSep = len(first) != len(second)
	}
	fmt.Println(" -> Separables:", isSep)
	return isSep
}
