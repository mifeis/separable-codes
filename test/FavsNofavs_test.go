package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/mifeis/Separable-Codes/lib"
)

//casos totals (favorables i no favorables)
//Fav: {0,0,0}|{0,0,1}, {0,0,1}|{1,1,1}, {1,1,1}|{0,0,0}, ...
//desFav: {0,0,0}|{0,0,0}, {0,0,1}|{1,0,1}, ...

func TestFavsNoFavs(t *testing.T) {
	var favs int
	var nofavs int

	fmt.Println("Total desfavorable cases for a code of "+strconv.Itoa(lib.WORDS)+" words: ", nofavs)
	fmt.Println("Total favorable cases for a code of "+strconv.Itoa(lib.WORDS)+" words: ", favs)

	if favs+nofavs == main.GetTotal(lib.Init()) {

	}

}
