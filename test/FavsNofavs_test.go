package main

import (
	"fmt"
	"testing"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib_aux"
	"github.com/mifeis/Separable-Codes/lib_main"
)

//casos totals (favorables i no favorables)
//Fav: {0,0,0}|{0,0,1}, {0,0,1}|{1,1,1}, {1,1,1}|{0,0,0}, ...
//desFav: {0,0,0}|{0,0,0}, {0,0,1}|{1,0,1}, ...

func TestFavs(t *testing.T) {
	var favs, nofavs int
	c := lib_main.Init()
	for i := 1; i < CASES; i++ {
		totalfavs, totalnofavs := getFavs(c, i)
		favs = +totalfavs
		nofavs = +totalnofavs
		fmt.Println("Total desfavorable cases:", nofavs)
		fmt.Println("Total favorable cases:", favs)
		favs = 0
		nofavs = 0

	}

	//	fmt.Println("Total desfavorable cases:", nofavs)
	//	fmt.Println("Total favorable cases:", favs)
}

//Funció que retorna els casos favorables i no favorables tenint en compte totes les possibles combinacions
//de grups disjunts (List0) i no disjunts (List1, List2) per a un array inicial de WORDS paraules i grups de GROUP elements
func getFavs(c []int, cas int) (int, int) {
	var favs, nofavs int
	var first, second lib_aux.Combi
	var arraymap map[lib_aux.Combi][]lib_aux.Combi

	switch cas {
	case 0:
		arraymap = combinations.List0(c)
		cas = 1
	case 1:
		arraymap = combinations.List1(c)
	case 2:
		arraymap = combinations.List2(c)
	}

	fmt.Println("Getting favorable and desfavorable cases for the case", cas)
	//Set a combination
	for k, g := range arraymap {
		first.Group = k.Group
		second.Group = g[0].Group
		break
	}

	defaultvalues := lib_aux.GetDefaultValues()
	fmt.Println("\n->", first.Group, "|", second.Group)

	//Tipus 1: {1,2,3}{4,5,6}; Tipus 2: {1,2,3}{1,4,5}; Tipus 3: {1,2,3}{1,2,4}
	for i := 0; i < len(defaultvalues); i++ {
		first.Value = defaultvalues[i]
		maxlen := len(defaultvalues)
		//contabilitzar d'alguna manera els casos repetits-> recorre l'array fins a GROUP-elements cops?
		for j := 0; j < maxlen/cas; j++ {
			//Set the repe elements of group
			setValues(first, &second)
			for l, v2 := range second.Value {
				//Set the leaving values
				if v2 == 2 {
					second.Value[l] = defaultvalues[j][l]
				} else {
					maxlen--
				}
			}
			if lib_aux.Separable(first.Value, second.Value) {
				favs++
			} else {
				nofavs++
			}
		}
		fmt.Println()
	}

	return favs, nofavs
}

//Function assign values from first to second if the columns are the same
//If not assignes the number 2 to the leaving columnes
func setValues(first lib_aux.Combi, second *lib_aux.Combi) {
	second.Value = [lib_aux.GROUP]int{2, 2, 2}
	for m, v1 := range first.Group {
		for n, v2 := range second.Group {
			if v1 == v2 {
				second.Value[m] = first.Value[n]
			}
		}
	}
}
