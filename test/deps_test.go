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

//Test que retorna el numero de combinacions dependents
func TestDeps(t *testing.T) {
	if lib.WORDS < 2*lib.GROUP {
		log.Fatal("num of words must be smaller than 2 * group elements")
	}
	initial := lib.Init(0, lib.WORDS)
	fmt.Println("\n...Getting dependence")
	getDeps(initial)
}

func getDeps(initial []int) {
	var array []map[int][]lib.Map
	arraypairs := make(map[string][][]int)
	var key string

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
						key = strconv.Itoa(reps) + strconv.Itoa(k1) + strconv.Itoa(k2)
						first = append(first, m.First...)
						second = append(second, s...)
						group = append(first, second...)
						if k1 == k2 {
							if !inversAlreadyInArray(arraypairs[key], group, reps) {
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

	xlsx := excelize.NewFile()
	title, _ := xlsx.NewStyle(lib.TITLE)
	subtitle, _ := xlsx.NewStyle(lib.SUBTITLE)
	text, _ := xlsx.NewStyle(lib.TEXT)
	bold, _ := xlsx.NewStyle(lib.BOLD)
	fill, _ := xlsx.NewStyle(lib.FILL)

	xlsx.SetSheetName("Sheet1", "Summary")
	intro := "DEPENDENCE BETWEEN EVENTS"

	xlsx.SetCellValue("Summary", "A1", intro)
	colf := excelize.ToAlphaString(12 - 1)
	xlsx.MergeCell("Summary", "A1", colf+"1")
	xlsx.SetCellStyle("Summary", "A1", colf+"1", title)

	xlsx.SetCellValue("Summary", "A2", "Array")
	xlsx.SetCellValue("Summary", "D2", "Type")
	xlsx.SetCellValue("Summary", "G2", "Element repetitions")
	xlsx.SetCellValue("Summary", "J2", "Total")

	xlsx.SetCellStyle("Summary", "J2", "L2", subtitle)
	xlsx.SetCellStyle("Summary", "A2", "G2", bold)

	xlsx.MergeCell("Summary", "A2", "C2")
	xlsx.MergeCell("Summary", "D2", "F2")
	xlsx.MergeCell("Summary", "G2", "I2")
	xlsx.MergeCell("Summary", "J2", "L2")

	var results []string
	var counted bool
	chartdata := make(map[string][][]int)

	for k, arraypair := range arraypairs {
		for o, v := range arraypair {
			var deps int

			for key, arrayToCompare := range arraypairs {
				var init, fin int
				if k == key {
					init = o + 1
					fin = 1
				} else {
					init, fin = 0, 0
				}
				for i := 0; i < len(arrayToCompare)-fin; i++ {
					index := (init + i) % (len(arrayToCompare))
					if lib.Dependent(v, arrayToCompare[index]) {
						deps++
					}
				}
			}

			for _, v := range results {
				if v == k {
					counted = true
				} else {
					counted = false
				}
			}

			if !counted {
				results = append(results, k)
				strings.SplitAfter(k, "")

				reps, _ := strconv.Atoi(string(k[0]))
				k1 := string(k[1])
				k2 := string(k[2])

				lib.LogDeps(k1, k2, reps, v, deps)
				filf := strconv.Itoa(len(results) + 2)
				xlsx.SetCellValue("Summary", "A"+filf, v)
				xlsx.SetCellValue("Summary", "D"+filf, k1+"x"+k2)
				xlsx.SetCellValue("Summary", "G"+filf, reps)
				xlsx.SetCellValue("Summary", "J"+filf, deps)
				xlsx.SetCellStyle("Summary", "A3", "I"+filf, text)
				xlsx.SetCellStyle("Summary", "J3", "J"+filf, fill)
				xlsx.MergeCell("Summary", "A"+filf, "C"+filf)
				xlsx.MergeCell("Summary", "D"+filf, "F"+filf)
				xlsx.MergeCell("Summary", "G"+filf, "I"+filf)
				xlsx.MergeCell("Summary", "J"+filf, "L"+filf)
				data := []int{reps, deps}
				chartdata[k1+"x"+k2] = append(chartdata[k1+"x"+k2], data)
			}
		}
	}

	xlsx.NewSheet("Graph")
	var serie []string

	f := 2
	xlsx.SetColVisible("Graph", "A", false)
	for c := 0; c < lib.GROUP; c++ {
		col := excelize.ToAlphaString(c + 1)
		xlsx.SetCellValue("Graph", col+"1", c)
		xlsx.SetColVisible("Graph", col, false)
	}
	coli := excelize.ToAlphaString(1)
	colf = excelize.ToAlphaString(lib.GROUP)
	for dimension, d := range chartdata {
		fil := strconv.Itoa(f)
		xlsx.SetCellValue("Graph", "A"+fil, dimension)
		f++
		var col string
		for _, data := range d {
			col = excelize.ToAlphaString(data[0] + 1)
			xlsx.SetCellValue("Graph", col+fil, data[1])
		}
		serie = append(serie, `{"name":"Graph!$A$`+fil+`","categories":"Graph!$B$1:$`+colf+`$1","values":"Graph!$`+coli+`$`+fil+`:$`+colf+`$`+fil+`"}`)
	}
	var series string
	for i, s := range serie {
		if i != len(serie)-1 {
			series = series + s + `,`
		} else {
			series = series + s
		}
	}

	colf = excelize.ToAlphaString(lib.GROUP + 2)
	xlsx.AddChart("Graph", colf+"2",
		`{
	"type":"col",
	"series":`+`[`+series+`]`+`,
	"title":
		{"name":"Dependence between events"}
	}`)

	index := xlsx.GetSheetIndex("Graph")
	xlsx.SetActiveSheet(index)

	path, _ := os.Getwd()
	split := strings.Split(path, "\\")
	var n string
	for i := 0; i < len(split)-1; i++ {
		n += split[i] + "\\"
	}
	filename := "dependence" + strconv.Itoa(lib.WORDS) + "x" + strconv.Itoa(lib.GROUP) + ".xlsx"
	xlsx.SaveAs(n + "out/test/deps/" + filename)

}

func inversAlreadyInArray(arraypairs [][]int, pair []int, reps int) bool {
	var length int

	for _, p := range arraypairs {
		length = compareArrays(p, pair)
		if (len(pair) + reps*2) == length {
			return true
		}
	}

	return false
}

func compareArrays(array1 []int, array2 []int) int {
	var l int
	//	fmt.Println("comparing len:", array1, array1)
	for _, v1 := range array1 {
		for _, v2 := range array2 {
			if v1 == v2 {
				l++
			}
		}
	}
	return l
}
