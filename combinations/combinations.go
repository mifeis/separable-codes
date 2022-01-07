package combinations

import (
	"github.com/mifeis/Separable-Codes/lib"
	"gonum.org/v1/gonum/stat/combin"
)

//Funció que retorna els elements disjunts en grups de GROUP elements de l'array inicial
func List(initial []int, t int) map[[lib.GROUP]int][][lib.GROUP]int {
	var combins [][lib.GROUP]int
	arraymap := make(map[[lib.GROUP]int][][lib.GROUP]int)

	//First combinations of GROUP elements
	groups := GetGroups(initial, [lib.GROUP]int{}, 1)
	for _, g := range groups {
		var list [][lib.GROUP]int
		combins = GetGroups(initial, g, t)
		list = append(list, combins...)
		arraymap[g] = list
	}
	return arraymap
}

//treure els tipus automaticaments
func GetGroups(initial []int, g [lib.GROUP]int, t int) [][lib.GROUP]int {
	var combins [][lib.GROUP]int
	var remaining []int
	slice := [lib.GROUP]int{}

	if g != [lib.GROUP]int{} {
		remaining = lib.RemoveSlice(initial, g[:]) //Remaining array
	} else {
		remaining = initial
	}

	switch t {
	case 1:
		//Casos disjunts
		//Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
		for i := 0; i < len(remaining); i++ {
			for j := i + 1; j < len(remaining); j++ {
				for k := j + 1; k < len(remaining); k++ {
					slice = [lib.GROUP]int{remaining[i], remaining[j], remaining[k]}
					/*
					 *	If pel cas de GROUP ==4
					 *	A arrel d'aixo tenir en compte els casos on 2n es menor a k
					 *	llavors les combinacions s'haurien de tenir en compte (for) per crear-
					 * 	les igualment (Parlem del tipus 4 a lib/apunts.txt)
					 */

					if lib.GROUP == 4 {
						//Tipus {1,2,3,4}|{5,6,7,8}, ...
						init := 2
						for l := k + 1; l < len(remaining); l++ {
							slice[init+1] = remaining[l]
							combins = append(combins, slice)
						}
					} else {
						combins = append(combins, slice)
					}
				}
			}
		}
	case 2:
		//Casos NO disjunts
		//Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}
		var init int
		for p := 0; p < lib.GROUP; p++ {
			init = 0
			for r := init; r < t-1; r++ {
				slice[r] = g[(r+p)%lib.GROUP]
				init++
			}
			for i := 0; i < len(remaining); i++ {
				slice[init] = remaining[i]
				for j := i + 1; j < len(remaining); j++ {
					slice[init+1] = remaining[j]

					if lib.GROUP == 4 {
						//Tipus {1,2,3,4}|{1,5,6,7}, {1,2,3,4}|{2,5,6,7}
						for l := j + 1; l < len(remaining); l++ {
							slice[init+2] = remaining[l]
							combins = append(combins, slice)
						}
					} else {
						combins = append(combins, slice)
					}
				}
			}
		}

	case 3:
		//Casos NO disjunts
		//Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}
		var init int
		//GROUP=3->3; GROUP=4->6; GROUP=5->10;
		for p := 0; p < len(combin.Combinations(lib.GROUP, t-1)); p++ {
			init = 0
			for r := init; r < t-1; r++ {
				//no va
				slice[r] = g[(r+p)%lib.GROUP]
				//				slice[r] = g[(r+p*init)%lib.GROUP]
				init++
			}
			for j := 0; j < len(remaining); j++ {
				slice[init] = remaining[j]
				if lib.GROUP == 4 {
					//Tipus {1,2,3,4}|{1,2,5,6}, {1,2,3,4}|{1,4,5,6}, {1,2,3,4}|{2,3,6,7}
					for l := j + 1; l < len(remaining); l++ {
						//canviar el 1/index per GROUP(t)
						slice[init+1] = remaining[l]
						combins = append(combins, slice)
					}
				} else {
					combins = append(combins, slice)
				}

			}
		}
	}

	return combins
}
