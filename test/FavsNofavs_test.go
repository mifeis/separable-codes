package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib_aux"
	"github.com/mifeis/Separable-Codes/lib_main"
)

//casos totals (favorables i no favorables)
//Fav: {0,0,0}|{0,0,1}, {0,0,1}|{1,1,1}, {1,1,1}|{0,0,0}, ...
//desFav: {0,0,0}|{0,0,0}, {0,0,1}|{1,0,1}, ...

func TestFavsNofavs(t *testing.T) {
	var favs, nofavs int
	c := lib_main.Init()
	for i := 0; i < CASES; i++ {
		totalfavs, totalnofavs := getFavs(c, i)
		favs = +totalfavs
		nofavs = +totalnofavs
	}

	fmt.Println("Total desfavorable cases:", nofavs)
	fmt.Println("Total favorable cases:", favs)
}

//Funció que retorna els casos favorables i no favorables tenint en compte totes les possibles combinacions
//de grups disjunts (List0) i no disjunts (List1, List2) per a un array inicial de WORDS paraules i grups de GROUP elements
func getFavs(c []int, cas int) (int, int) {
	var totalfavs, totalnofavs int

	switch cas {
	case 0:
		totalfavs, totalnofavs = GetFavsNofavs(c)
	case 1:
		totalfavs, totalnofavs = combinations.GetFavsNofavs1(c)
	case 2:
		totalfavs, totalnofavs = combinations.GetFavsNofavs2(c)

	}
	return totalfavs, totalnofavs
}

func GetFavsNofavs(c []int) (int, int) {
	var favs, nofavs int
	var first, second lib_aux.Combi
	arraymap := combinations.List0(c)
	//Set a combination
	for k, g := range arraymap {
		first.Group = k.Group
		//		first.Value = [lib_aux.GROUP]int{2, 2, 2}
		second.Group = g[0].Group
		break
	}

	defaultvalues := lib_aux.GetDefaultValues()
	log.Println(first.Group, "|", second.Group)

	//Tipus 1: {1,2,3}{4,5,6}; Tipus 2: {1,2,3}{1,4,5}; Tipus 3: {1,2,3}{1,2,4}
	for i := 0; i < len(defaultvalues); i++ {
		first.Value = defaultvalues[i]
		//Set the repe elements of group
		assign(first, second)
		//Set the leaving values
		//contabilitzar d'alguna manera els casos repetits-> recorre l'array fins a GROUP-elements cops?
		for j := 0; j < len(defaultvalues); j++ {
			for l, v2 := range second.Value {
				if v2 == 2 {
					log.Println("entra")
					second.Value[l] = defaultvalues[j][l]
					log.Println("second value: ", second.Value[l], "default value", defaultvalues[j][l])
				}
			}
			if lib_aux.Separable(first.Value, second.Value) {
				favs++
			} else {
				nofavs++
			}
			assign(first, second)
		}
	}

	return favs, nofavs
}

//Function assign assignes values from first to second if the columns are the same
//If not assignes the number 2 to the leaving columnes
func assign(first lib_aux.Combi, second lib_aux.Combi) {
	second.Value = [lib_aux.GROUP]int{2, 2, 2}
	for _, v1 := range first.Group {
		for l, v2 := range second.Group {
			if v1 == v2 {
				second.Value[l] = v1
			}
		}
	}
}
