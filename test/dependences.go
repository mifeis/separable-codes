package test

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/mifeis/Separable-Codes/combinations"
	"github.com/mifeis/Separable-Codes/lib"
)

//Test que retorna el numero de combinacions dependents
func GetDependences(initial []int) {
	fmt.Println("\n...Getting dependence")
	arraypairs := getArrayPairs(initial)
	getDependences(arraypairs)
	fmt.Println("Done! Check /out/test/deps folder")
}

func getArrayPairs(initial []int) map[string][][]int {
	var array []map[int][]lib.Map
	arraypairs := make(map[string][][]int)

	for i := 0; i < lib.REPS; i++ {
		arraymap := combinations.List(initial, i)
		array = append(array, lib.Sort(arraymap))
	}

	for reps, arraymaps := range array {
		//Set a combination
		for k1, am := range arraymaps {
			for _, m := range am {
				for k2, ss := range m.Seconds {
					for _, s := range ss {
						var first, second, group []int
						key := strconv.Itoa(reps) + strconv.Itoa(k1) + strconv.Itoa(k2)
						first = append(first, m.First...)
						second = append(second, s...)
						group = append(first, second...)
						if k1 == k2 {
							if !lib.InversAlreadyInArray(arraypairs[key], group, reps) {
								arraypairs[key] = append(arraypairs[key], group)
							}
						} else {
							arraypairs[key] = append(arraypairs[key], group)
						}
					}
				}
			}
		}
	}
	return arraypairs
}

func getDependences(arraypairs map[string][][]int) {
	xlsx := excelize.NewFile()
	title, subtitle, text, bold, fill := lib.GetExcelStyles(xlsx)
	lib.SetExcelIntro(xlsx, "DEPENDENCE BETWEEN EVENTS", 18, title)

	xlsx.SetCellValue("Summary", "A2", "Array")
	xlsx.SetCellValue("Summary", "D2", "Type")
	xlsx.SetCellValue("Summary", "G2", "Element repetitions")
	xlsx.SetCellValue("Summary", "J2", "Dependent")
	xlsx.SetCellValue("Summary", "M2", "No dependent")
	xlsx.SetCellValue("Summary", "P2", "Total")

	xlsx.SetCellStyle("Summary", "J2", "P2", subtitle)
	xlsx.SetCellStyle("Summary", "A2", "P2", bold)

	xlsx.MergeCell("Summary", "A2", "C2")
	xlsx.MergeCell("Summary", "D2", "F2")
	xlsx.MergeCell("Summary", "G2", "I2")
	xlsx.MergeCell("Summary", "J2", "L2")
	xlsx.MergeCell("Summary", "M2", "O2")
	xlsx.MergeCell("Summary", "P2", "S2")

	var results []string
	chartdata := make(map[string][][]int)

	for k, arraypair := range arraypairs {
		array := arraypair[0]
		var deps, nodeps, total int

		for keyToCompare, arrayToCompare := range arraypairs {
			var init, fin int
			if k == keyToCompare {
				init, fin = 1, 1
			} else {
				init, fin = 0, 0
			}
			for i := 0; i < len(arrayToCompare)-fin; i++ {
				index := (init + i) % (len(arrayToCompare))
				if lib.Dependent(array, arrayToCompare[index]) {
					deps++
				} else {
					nodeps++
				}
			}
			//Events (combinacions d'arrays) totals
			total += len(arrayToCompare)
		}

		results = append(results, k)
		strings.SplitAfter(k, "")

		reps, _ := strconv.Atoi(string(k[0]))
		k1 := string(k[1])
		k2 := string(k[2])

		//	lib.LogDeps(k1, k2, reps, v, deps)
		filf := strconv.Itoa(len(results) + 2)
		xlsx.SetCellValue("Summary", "A"+filf, array)
		xlsx.SetCellValue("Summary", "D"+filf, k1+"x"+k2)
		xlsx.SetCellValue("Summary", "G"+filf, reps)
		xlsx.SetCellValue("Summary", "J"+filf, deps)
		xlsx.SetCellValue("Summary", "M"+filf, nodeps)
		xlsx.SetCellValue("Summary", "P"+filf, total)

		xlsx.SetCellStyle("Summary", "A3", "I"+filf, text)
		xlsx.SetCellStyle("Summary", "J3", "J"+filf, fill)
		xlsx.SetCellStyle("Summary", "M3", "P"+filf, text)

		xlsx.MergeCell("Summary", "A"+filf, "C"+filf)
		xlsx.MergeCell("Summary", "D"+filf, "F"+filf)
		xlsx.MergeCell("Summary", "G"+filf, "I"+filf)
		xlsx.MergeCell("Summary", "J"+filf, "L"+filf)
		xlsx.MergeCell("Summary", "M"+filf, "O"+filf)
		xlsx.MergeCell("Summary", "P"+filf, "S"+filf)

		data := []int{reps, deps}
		chartdata[k1+"x"+k2] = append(chartdata[k1+"x"+k2], data)
	}
	lib.AddDependenceChart(xlsx, chartdata)
	lib.SaveExcel(xlsx, 3, 0)
}
