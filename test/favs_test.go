package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//casos totals (favorables i no favorables)
//Fav: {0,0,0}|{0,0,1}, {0,0,1}|{1,1,1}, {1,1,1}|{0,0,0}, ...
//desFav: {0,0,0}|{0,0,0}, {0,0,1}|{1,0,1}, ...

func TestFavs(t *testing.T) {
	if lib.WORDS <= lib.GROUP {
		log.Fatal("Can't combine elements because num of words is smaller than group elements")
	}
	initial := lib.Init(0, lib.WORDS)
	for i := 1; i < lib.REPS; i++ {
		lib.LogTipus(i)
		favs, nofavs := getFavs(initial, i)
		fmt.Println("Total favorable cases:", favs)
		fmt.Println("Total desfavorable cases:", nofavs)
	}

	//	fmt.Println("Total favorable cases:", favs)
}

//Funció que retorna els casos favorables i no favorables tenint en compte totes les possibles combinacions
//de grups disjunts (List0) i no disjunts (List1, List2) per a un array inicial de WORDS paraules i grups de GROUP elements
func getFavs(initial []int, tipus int) (int, int) {
	var favs, nofavs int
	var first, second lib.Code

	arraymap := combinations.List(initial, tipus)
	fmt.Print("...Getting favorable and desfavorable cases for the type ", tipus)
	//Set a combination
	for _, m := range arraymap {
		first.Row = m.First       //rows
		second.Row = m.Seconds[0] //rows
		break
	}
	fmt.Println("->", first.Row, "|", second.Row)

	defaultvalues := lib.GetDefaultValues()

	// Tipus 1: {1,2,3}{4,5,6}
	// Tipus 2: {1,2,3}{1,4,5}
	// Tipus 3: {1,2,3}{1,2,4}
	for i := 0; i < len(defaultvalues); i++ {
		first.Values = defaultvalues[i]
		fmt.Println()
		//contabilitzar d'alguna manera els casos repetits-> recorre l'array fins a GROUP-elements cops?
		for j := 0; j < len(defaultvalues)/(tipus); j++ {
			//Set the repe elements of group
			lib.SetValues(first, &second)
			for l, v := range second.Values {
				//Set the leaving values
				if v == 2 {
					second.Values[l] = defaultvalues[j][l]
				}
			}
			if lib.Separable(first.Values, second.Values) {
				favs++
			} else {
				nofavs++
			}
		}
	}

	return favs, nofavs
}
