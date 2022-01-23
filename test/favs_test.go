package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//casos totals (favorables i no favorables)
//Fav: {0,0,0}|{0,0,1}, {0,0,1}|{1,1,1}, {1,1,1}|{0,0,0}, ...
//desFav: {0,0,0}|{0,0,0}, {0,0,1}|{1,0,1}, ...

func TestFavs(t *testing.T) {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words is small")
	}
	initial := lib.Init(0, lib.WORDS)
	for i := 0; i < lib.REPS; i++ {
		fmt.Println("\n...Getting favorable and desfavorable cases for", i, "element repetitions")
		/*favs, nofavs := */ getFavs(initial, i)
		//lib.LogFavs(favs, nofavs)
	}
}

//Funció que retorna els casos favorables i no favorables tenint en compte totes les possibles combinacions
//de grups disjunts (List0) i no disjunts (List1, List2) per a un array inicial de WORDS paraules i grups de GROUP elements
func getFavs(initial []int, reps int) /*(int, int)*/ {
	var first, second lib.Code
	var c int

	arraymap := combinations.List(initial, reps)
	arraymaps := lib.Sort(arraymap)

	xlsx := excelize.NewFile()
	title, _ := xlsx.NewStyle(lib.TITLE)
	subtitle, _ := xlsx.NewStyle(lib.SUBTITLE)
	text, _ := xlsx.NewStyle(lib.TEXT)
	bold, _ := xlsx.NewStyle(lib.BOLD)

	xlsx.SetSheetName("Sheet1", "Summary")
	intro := "FAVORABLE AND DESFAVORABLE COMBINATIONS FOR " + strconv.Itoa(reps) + " ELEMENT REPETITIONS"

	xlsx.SetCellValue("Summary", "A1", intro)
	colf := excelize.ToAlphaString(lib.GROUP*2 + 2)
	xlsx.MergeCell("Summary", "A1", colf+"1")
	xlsx.SetCellStyle("Summary", "A1", colf+"1", title)

	xlsx.SetCellValue("Summary", "A3", "Favourable")
	xlsx.SetCellValue("Summary", "A4", "Unfavourable")
	xlsx.SetCellValue("Summary", "A5", "Total")
	xlsx.MergeCell("Summary", "A3", "B3")
	xlsx.MergeCell("Summary", "A4", "B4")
	xlsx.MergeCell("Summary", "A5", "B5")
	xlsx.SetCellStyle("Summary", "A3", "A4", text)
	xlsx.SetCellStyle("Summary", "A5", "A5", subtitle)

	for k, am := range arraymaps {
		//Set a combination

		for _, m := range am {
			first.Row = m.First //rows
			for k2, s := range m.Seconds {
				var favs, nofavs int
				second.Row = s[0] //rows

				length := strconv.Itoa(k) + "x" + strconv.Itoa(k2)
				xlsx.NewSheet(length)

				fmt.Println()
				fmt.Println("->", first.Row, "|", second.Row)

				//retornar a la mateixa funcio
				defaultvalues1 := lib.GetDefaultValues(len(first.Row))
				defaultvalues2 := lib.GetDefaultValues(len(second.Row))
				for i := 0; i < len(defaultvalues1); i++ {
					first.Values = defaultvalues1[i]
					fmt.Println()

					for j := 0; j < (len(defaultvalues2) / (reps + 1)); j++ {
						//Set the repe elements of group
						lib.SetValues(first, &second)
						for l, v := range second.Values {
							//Set the leaving values
							if v == 2 {
								second.Values[l] = defaultvalues2[j][l]
							}
						}
						fil := strconv.Itoa(j + 1)
						xlsx.SetCellValue(length, "A"+fil, first.Values)
						xlsx.SetCellValue(length, "B"+fil, "<->")
						xlsx.SetCellValue(length, "C"+fil, second.Values)
						xlsx.SetCellStyle(length, "A"+fil, "C"+fil, text)

						if lib.Separable(first.Values, second.Values) {
							xlsx.SetCellValue(length, "D"+fil, "Separable")
							favs++
						} else {
							xlsx.SetCellValue(length, "D"+fil, "NO Separable")
							nofavs++
						}
						xlsx.MergeCell(length, "D"+fil, "E"+fil)
						xlsx.SetCellStyle(length, "D"+fil, "E"+fil, bold)
					}
				}
				col := excelize.ToAlphaString(2 + c)
				xlsx.SetCellValue("Summary", col+"2", strconv.Itoa(k)+"x"+strconv.Itoa(k2))
				xlsx.SetCellStyle("Summary", col+"2", col+"2", subtitle)
				xlsx.SetCellValue("Summary", col+"3", favs)
				xlsx.SetCellValue("Summary", col+"4", nofavs)
				xlsx.SetCellValue("Summary", col+"5", favs+nofavs)
				xlsx.SetCellStyle("Summary", col+"3", col+"4", text)
				xlsx.SetCellStyle("Summary", col+"5", colf+"5", bold)
				c++
				lib.LogFavs(favs, nofavs)
			}
			break
		}
	}
	index := xlsx.GetSheetIndex("Summary")
	xlsx.SetActiveSheet(index)

	path, _ := os.Getwd()
	split := strings.Split(path, "\\")
	var n string
	for i := 0; i < len(split)-1; i++ {
		n += split[i] + "\\"
	}
	filename := "favorables" + strconv.Itoa(lib.WORDS) + "x" + strconv.Itoa(lib.GROUP) + "_" + strconv.Itoa(reps) + "_" + "element_repetitions.xlsx"
	xlsx.SaveAs(n + "out/test/favs/" + filename)

	//	return favs, nofavs
}
