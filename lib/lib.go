package lib

import (
	"math"
)

const (
	WORDS = 16
	GROUP = 4
	REPS  = GROUP
)

// Map struct has two attributes. The first one is an array
// containing elements from 1 to WORDS, which would be the rows of the
// code picked in a group of size from 1 to GROUP.
// The second object is a map that has as value an array with all the
// other possible groups of combinations given the first element of
// the pair and filtered by a key which is the size of the second
// group.

type Map struct {
	First   []int
	Seconds map[int][][]int // For example: [1 2 3]->[3:[[1 5 6][4 7 8]...]2:[[3 5],[4 6]...]1:[[1][2][6][7][8]...]]

}

// Init returns an array initialised with the values from one to
// WORDS
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

// Classifies the arraymaps append by the combinations method by
// the length of the first group. No matter how many repetitions in
// the pairs are. Returns a map with the length as a key and an
// array of the distinct objects with a first group and all its
// combinations.
func Sort(arraymap []Map) map[int][]Map {
	arraymaps := make(map[int][]Map)

	for _, m := range arraymap {
		arraymaps[len(m.First)] = append(arraymaps[len(m.First)], m)
	}
	return arraymaps
}

// Struct that defines binary values to a set of rows
type Code struct {
	Row    []int
	Values []int
}

// GetDefaultValues returns all binary possible combinations for a
// depending on the length l passed by argument
func GetDefaultValues(l int) [][]int {
	var values [][]int

	len := int(math.Exp2(float64(l)))
	for t := 0; t < len; t++ {
		var slice []int
		for i := 0; i < l; i++ {
			ijk := t / (len / int(math.Exp2(float64(i+1))))
			slice = append(slice, ijk%2)
		}
		values = append(values, slice[:])
	}
	//	fmt.Println("Possible binari values for a group of", l, "elements:", values)
	return values
}

// Function that assignes values from first to second if the rows
// are the same. If not it assignes the default number 2 to the
// leaving ones to distinguish them.
func SetValues(first Code, second *Code) {
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

//Separable says if the two arrays are separable or not
func Separable(group1 []int, group2 []int) bool {
	first := make(map[int]int)
	second := make(map[int]int)

	//	fmt.Print("Comparing ", group1, " with ", group2)

	// Bool that assign the key 0 the zero value if is on group1
	// and the 1 if applies the same way. The map’s length will be
	// 2 or 1 then.
	for _, v := range group1 {
		first[v] = v
	}

	for _, v := range group2 {
		second[v] = v
	}

	var isSep bool // Separable bool variable
	if len(first) == len(second) {
		// Find out about special cases like:
		// Non separables-> (0,0,0)|(0,0,0) and (1,1,1)|(1,1,1)
		// Separables-> (0,0,0)|(1,1,1) and (1,1,1)|(0,0,0)

		_, z1 := first[0]  // Pretended zero in group 1
		_, z2 := second[0] // Pretended zero in group 2
		if (z1 && z2) || (!z1 && !z2) {
		} else {
			isSep = true
		}
	} else {
		// Compare cases where the length is clearly distinct
		// or equal
		// Non separables-> (0,0,1)|(1,0,0), (0,0,0)|(0,0,0)...
		// Separables-> (0,0,0)|(1,0,0), (1,1,1)|(0,1,0)...

		isSep = len(first) != len(second)
	}
	//	fmt.Println(" -> Separables:", isSep)
	return isSep
}

// Says if the two arrays passed by argument are dependent or not
func Dependent(array1 []int, array2 []int) bool {
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v1 == v2 {
				return true
			}
		}
	}
	return false
}

// Returns a bool that is true if an array passed by argument is
// already located in another array of pairs also passed.

func InversAlreadyInArray(arraypairs [][]int, pair []int, reps int) bool {
	var length int

	for _, p := range arraypairs {
		length = CompareArrays(p, pair)
		if (len(pair) + reps*2) == length {
			return true
		}
	}

	return false
}

// Compare two arrays
func CompareArrays(array1 []int, array2 []int) int {
	var l int
	//	fmt.Println("comparing len:", array1, array1)
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v1 == v2 {
				l++
			}
		}
	}
	return l
}
