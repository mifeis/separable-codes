package combinations

import (
	"fmt"

	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Canviar
func GetCombinations(c []int) {
	fmt.Println("\n...Getting combinations ")
	for reps := 0; reps < lib.REPS; reps++ {
		arraymap := List(c, reps)
		arraymaps := lib.Sort(arraymap)
		lib.WriteCombinations(arraymaps, reps)
	}
	fmt.Println("Done! Check /out/combinations folder")
}

// List returns an array that contains the different Maps for a given
// number of repetitions in the elements of a pair of groups. The
// initial array passed by argument is the one with the number of
// each row in a code of WORDS words.
// In this case it would be [1 2 3 4 5 6 7 8]
func List(initial []int, reps int) []lib.Map {
	// Local variables initialization
	var combins map[int][]int
	arraymap := []lib.Map{}

	//First combinations of GROUP elements
	groups := getCombins(initial, []int{}, reps, 0)
	// Loop through all the different first combinations from size
	// GROUP to 1 (If repetitions are 0, if not they will be from
	// GROUP to the minimum size value)
	for _, g := range groups {
		list := make(map[int][][]int)
		// Get all second combinations for the given first group
		// Remove slice is an auxiliary function described
		// afterwards
		combins = getCombins(lib.RemoveSlice(initial, g[:]), g, 0, reps)
		// Sort the second arrays by its length in a map
		for _, c := range combins {
			list[len(c)] = append(list[len(c)], c)
		}
		m := lib.Map{
			First:   g,
			Seconds: list,
		}
		// Array with the distinct Maps objects for each first
		// group
		arraymap = append(arraymap, m)
	}
	return arraymap
}

// getCombins returns in a map the different combos for a remaining
// array passed by argument and a given first group of the pair. If
// the second parameter is empty, we are looking for the leading
// groups and the remaining array is the initial.
func getCombins(remaining []int, g []int, init int, reps int) map[int][]int {
	// Local variables initialization
	var key int // In order to sort the map. The value is needless
	maxLen := lib.GROUP
	combins := make(map[int][]int, 1000)

	if len(g) != 0 {
		maxLen = len(g)
	}

	// The combin library is used next to define the number of
	// iterations to do in case the repetitions are bigger than 0,
	// so the elements placed from the first group in the second
	// combination swap. In other case its length is 1 due to the 0
	// passed by argument as reps when the method is invoked
	// initially
	in := combin.Combinations(maxLen, reps)
	for p := 0; p < len(in); p++ {
		var slice1 []int
		// Loop to assign non disjoint values to the second group.
		for r := 0; r < reps; r++ {
			slice1 = append(slice1, g[in[p][r]])
		}
		if len(slice1) != 0 {
			combins[key] = slice1[:]
			key++
		}

		//Loop to allocate the leaving disjoint values
		for l := maxLen - reps; l > init; l-- {
			indexes := combin.Combinations(len(remaining), l)
			// Use of the combin library to obtain the indexes to go
			// through the remaining array passed by argument and
			// obtain its different combinations.
			for _, index := range indexes {
				var slice2 []int
				// Fix the previous disjoint values if required
				slice2 = append(slice2, slice1...)
				for _, v := range index {
					slice2 = append(slice2, remaining[v])
				}
				combins[key] = slice2[:]
				key++
			}
		}
	}
	return combins
}
