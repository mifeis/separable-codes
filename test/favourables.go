package test

import (
	"fmt"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//Fav: {0,0,0}|{0,0,1}, {0,0,1}|{1,1,1}, {1,1,1}|{0,0,0}, ...
//desFav: {0,0,0}|{0,0,0}, {0,0,1}|{1,0,1}, ...

func GetFavourables(initial []int) {
	fmt.Println("\n...Getting favorable and desfavorable cases")
	for i := 0; i < lib.REPS; i++ {
		arraymap := combinations.List(initial, i)
		arraymaps := lib.Sort(arraymap)
		GetFavs(arraymaps, i)
	}
	fmt.Println("Done! Check /out/test/favs folder")
}

//GetFavs gets the favourable and unfavourable cases of each type
func GetFavs(arraymaps map[int][]lib.Map, reps int) {
	// Local variables initialization
	var first, second lib.Code
	var c int
	// Excel initialization
	xlsx := excelize.NewFile()
	title, subtitle, text, bold, _ := lib.GetExcelStyles(xlsx)
	lib.SetExcelIntro(xlsx, "FAVORABLE AND DESFAVORABLE COMBINATIONS FOR "+strconv.Itoa(reps)+" ELEMENT REPETITIONS", lib.GROUP*2+3, title)

	xlsx.SetCellValue("Summary", "A3", "Favourable")
	xlsx.SetCellValue("Summary", "A4", "Unfavourable")
	xlsx.SetCellValue("Summary", "A5", "Total")
	xlsx.MergeCell("Summary", "A3", "B3")
	xlsx.MergeCell("Summary", "A4", "B4")
	xlsx.MergeCell("Summary", "A5", "B5")
	xlsx.SetCellStyle("Summary", "A3", "A4", text)
	xlsx.SetCellStyle("Summary", "A5", "A5", subtitle)

	// Start loop going through all type of combinations
	for k, am := range arraymaps {
		//Set a combination
		for _, m := range am {
			first.Row = m.First // Assign the rows numbers
			for k2, s := range m.Seconds {
				var favs, nofavs int
				second.Row = s[0] // Assign the second group rows

				length := strconv.Itoa(k) + "x" + strconv.Itoa(k2)
				xlsx.NewSheet(length)

				//	fmt.Println()
				//	fmt.Println("->", first.Row, "|", second.Row)

				// defaultvalues contains an array with the binary
				// combos returned by the auxiliary function
				defaultvalues1 := lib.GetDefaultValues(len(first.Row))
				defaultvalues2 := lib.GetDefaultValues(len(second.Row))
				f := 1 // Excel variable
				// The loop gives assign each possible value in the
				// array returned before
				for i := 0; i < len(defaultvalues1); i++ {
					first.Values = defaultvalues1[i]
					//	fmt.Println()

					for j := 0; j < (len(defaultvalues2) / (reps + 1)); j++ {
						// Set the repeated elements of group from the
						// first to the second pack
						lib.SetValues(first, &second)
						for l, v := range second.Values {
							// Set the values not assigned to the
							// second group. If they are placed in 2,
							// then they need to be setted.
							if v == 2 {
								second.Values[l] = defaultvalues2[j][l]
							}
						}
						// Excel printing
						fil := strconv.Itoa(f)
						xlsx.SetCellValue(length, "A"+fil, first.Values)
						xlsx.SetCellValue(length, "B"+fil, "<->")
						xlsx.SetCellValue(length, "C"+fil, second.Values)
						xlsx.SetCellStyle(length, "A"+fil, "C"+fil, text)
						f++
						// Separable auxiliary function returns a bool
						if lib.Separable(first.Values, second.Values) {
							xlsx.SetCellValue(length, "D"+fil, "Separable") // Excel printing
							favs++
						} else {
							xlsx.SetCellValue(length, "D"+fil, "NO Separable") // Excel printing
							nofavs++
						}
						// Excel printing
						xlsx.MergeCell(length, "D"+fil, "E"+fil)
						xlsx.SetCellStyle(length, "D"+fil, "E"+fil, bold)
					}
				}
				// Excel printing
				col := excelize.ToAlphaString(c + 2)

				xlsx.SetCellValue("Summary", col+"2", strconv.Itoa(k)+"x"+strconv.Itoa(k2))
				xlsx.SetCellStyle("Summary", col+"2", col+"2", subtitle)
				xlsx.SetCellValue("Summary", col+"3", favs)
				xlsx.SetCellValue("Summary", col+"4", nofavs)
				xlsx.SetCellValue("Summary", col+"5", favs+nofavs)
				xlsx.SetCellStyle("Summary", col+"3", col+"4", text)
				xlsx.SetCellStyle("Summary", col+"5", col+"5", bold)
				c++
				//	lib.LogFavs(favs, nofavs)
			}
			// Break because only needs one pair of each type to know
			// its favourable and unfavourable cases.
			break
		}
		// Excel saving
		lib.SaveExcel(xlsx, 2, reps)
	}
}
