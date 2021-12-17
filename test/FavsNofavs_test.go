package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib_aux"
	"github.com/mifeis/Separable-Codes/lib_main"
)

//casos totals (favorables i no favorables)
//Fav: {0,0,0}|{0,0,1}, {0,0,1}|{1,1,1}, {1,1,1}|{0,0,0}, ...
//desFav: {0,0,0}|{0,0,0}, {0,0,1}|{1,0,1}, ...

func TestFavsNoFavs(t *testing.T) {
	var favs, nofavs int
	c := lib_main.Init()
	for i := 0; i < CASES; i++ {
		totalfavs, totalnofavs := getFavs(c, i)
		favs = +totalfavs
		nofavs = +totalnofavs
	}

	fmt.Println("Total desfavorable cases for a code of "+strconv.Itoa(lib_aux.WORDS)+" words: ", nofavs)
	fmt.Println("Total favorable cases for a code of "+strconv.Itoa(lib_aux.WORDS)+" words: ", favs)

	if favs+nofavs != lib_main.GetTotal(c) {
		t.Failed()
	}
}

//Funció que retorna els casos favorables i no favorables tenint en compte totes les possibles combinacions
//de grups disjunts (List0) i no disjunts (List1, List2) per a un array inicial de WORDS paraules i grups de GROUP elements
func getFavs(c []int, cas int) (int, int) {
	var totalfavs, totalnofavs int
	switch cas {
	case 0:
		totalfavs, totalnofavs = combinations.GetFavsNofavs0(c)
	case 1:
		totalfavs, totalnofavs = combinations.GetFavsNofavs1(c)
	case 2:
		totalfavs, totalnofavs = combinations.GetFavsNofavs2(c)
	default:
		return 0, 0
	}
	return totalfavs / 2, totalnofavs / 2
}
