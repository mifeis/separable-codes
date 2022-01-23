package lib

import (
	"fmt"
	"math"
)

const (
	REPS = 4

	WORDS = 16
	GROUP = 4
)

//Estructura que conté el primer element d'un tamany desde GROUP elements fins a 1
//i un segon array que consta de totes les combinacions possibles per aquest primer grup
type Map struct {
	First   []int
	Seconds map[int][][]int
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
	return remaining[:WORDS-len(slice)]
}

//Removes the index from the slice
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

//ordena els arraymaps segons el tamany del primer grup
func Sort(arraymap []Map) map[int][]Map {
	arraymaps := make(map[int][]Map)

	for _, m := range arraymap {
		arraymaps[len(m.First)] = append(arraymaps[len(m.First)], m)
	}
	return arraymaps
}

//Estructura que defineix valors (0/1) per a un grup d'elements
type Code struct {
	Row    []int
	Values []int
}

//Retorna totes les combinacions de valors (0/1) per un array de length l
func GetDefaultValues(l int) [][]int {
	var values [][]int
	//s'haura de passar len per argument depenent del GROUP del moment (3,2,1...)
	len := int(math.Exp2(float64(l)))
	for t := 0; t < len; t++ {
		var slice []int
		for i := 0; i < l; i++ {
			ijk := t / (len / int(math.Exp2(float64(i+1))))
			slice = append(slice, ijk%2)
		}
		values = append(values, slice[:])
	}
	fmt.Println("Possible binari values for a group of", l, "elements:", values)
	return values
}

//Function assign values from first to second if the columns are the same
//If not assignes the number 2 to the leaving columnes
func SetValues(first Code, second *Code) {
	//comprobar
	second.Values = []int{}
	for i := 0; i < len(second.Row); i++ {
		second.Values = append(second.Values, 2)
	}
	//	log.Println(first, second)

	for m, v1 := range first.Row {
		for n, v2 := range second.Row {
			if v1 == v2 {
				second.Values[n] = first.Values[m]
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

//Says if the two arrays are dependent or not
func Dependent(array1 []int, array2 []int) bool {
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v1 == v2 {
				//			fmt.Println(array1, " dependent with ", array2)
				//				fmt.Println(" -> Dependents")
				return true
			}
		}
	}
	//	fmt.Println(" -> NO Dependents")
	return false
}
