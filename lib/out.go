package lib

import (
	"fmt"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	BOLD  = `{"alignment":{"horizontal":"center","vertical":"center"},"font":{"bold":true}}`
	TITLE = `{"alignment":{"horizontal":"center","vertical":"center"},
	"font":{"bold":true},"fill":{"type":"pattern","color":["#BDD7EE"],"pattern":1}}`
	SUBTITLE = `{"alignment":{"horizontal":"center","vertical":"center"},
	"font":{"bold":true},"fill":{"type":"pattern","color":["#F2F2F2"],"pattern":1}}`
	TEXT = `{"alignment":{"horizontal":"center","vertical":"center"}}`
	FILL = `{"alignment":{"horizontal":"center","vertical":"center"},
	"fill":{"type":"pattern","color":["#F2F2F2"],"pattern":1}}`
)

func LogTipus(k1 int, k2 int, v int) string {
	tipus := strconv.Itoa(k1) + " x " + strconv.Itoa(k2) + "-> "
	fmt.Println(tipus, v)
	return tipus
}

/*
func LogFavs(favs int, nofavs int) {
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Total favorable cases:", favs)
	fmt.Println("Total desfavorable cases:", nofavs)
	fmt.Println("--------------------------------------------------------------------------------")
}

func LogDeps(k1 string, k2 string, reps int, v []int, total int) {
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println(k1, "x", k2, "(", reps, "element repetitions ) ->", v)
	fmt.Println("Total dependent pairs:", total)
	fmt.Println("--------------------------------------------------------------------------------")
}
*/

//Returns the total number of combinations and creates an excel with the data
func WriteCombinations(arraymaps map[int][]Map, reps int) {
	xlsx := excelize.NewFile()
	title, subtitle, text, bold, _ := GetExcelStyles(xlsx)
	fmt.Println(strconv.Itoa(reps) + " elemens repetitions:")
	SetExcelIntro(xlsx, "COMBINATIONS FOR "+strconv.Itoa(reps)+" ELEMENT REPETITIONS", GROUP*2+2, title)

	for k, am := range arraymaps {
		total := make(map[int]int)
		length := strconv.Itoa(k) + " group elements"
		xlsx.NewSheet(length)
		for c, m := range am {
			coli := excelize.ToAlphaString(c*len(m.First) + c)
			colf := excelize.ToAlphaString((c*len(m.First) + len(m.First) - 1 + c))
			fil := "1"
			xlsx.SetCellValue(length, coli+fil, m.First)
			xlsx.MergeCell(length, coli+fil, colf+fil)
			xlsx.SetCellStyle(length, coli+fil, colf+fil, title)

			for k2 := range m.Seconds {
				//3x1, 3x2, 3x3
				tipus := strconv.Itoa(k) + "x" + strconv.Itoa(k2)
				colii := excelize.ToAlphaString(c*len(m.First) + c + (k2 - 1))
				fil = "2"
				xlsx.SetCellValue(length, colii+fil, tipus)
				xlsx.SetCellStyle(length, colii+fil, colii+fil, subtitle)
				for f, s := range m.Seconds[k2] {
					fil := strconv.Itoa(f + 3)
					xlsx.SetCellValue(length, colii+fil, s)
					xlsx.SetCellStyle(length, colii+fil, colii+fil, text)
					//					fmt.Println("\t\t", m.First, "|", s)
					total[k2]++
				}
			}
			//			fmt.Println()
		}
		var all int
		var coli, colf, fil string

		for l, v := range total {
			if l == k {
				v = v / 2
			}
			all += v
			coli = excelize.ToAlphaString(2 * (l - 1))
			colf = excelize.ToAlphaString((2 * (l - 1)) + 1)
			fil = strconv.Itoa(2 * k)
			xlsx.SetCellValue("Summary", coli+fil, strconv.Itoa(k)+"x"+strconv.Itoa(l)+": ")
			xlsx.SetCellStyle("Summary", coli+fil, coli+fil, bold)
			xlsx.SetCellValue("Summary", colf+fil, v)
			xlsx.SetCellStyle("Summary", colf+fil, colf+fil, text)
			LogTipus(k, l, v)
		}
		coli = excelize.ToAlphaString((2 * GROUP))
		colf = excelize.ToAlphaString((2 * GROUP) + 1)
		xlsx.SetCellValue("Summary", coli+fil, "Total")
		xlsx.MergeCell("Summary", coli+fil, colf+fil)
		colf = excelize.ToAlphaString((2 * GROUP) + 2)
		xlsx.SetCellValue("Summary", colf+fil, all)
		xlsx.SetCellStyle("Summary", coli+fil, colf+fil, subtitle)

		fmt.Println("Total:\t", all)
	}
	SaveExcel(xlsx, 1, reps)
}

func GetExcelStyles(xlsx *excelize.File) (int, int, int, int, int) {
	title, _ := xlsx.NewStyle(TITLE)
	subtitle, _ := xlsx.NewStyle(SUBTITLE)
	text, _ := xlsx.NewStyle(TEXT)
	bold, _ := xlsx.NewStyle(BOLD)
	fill, _ := xlsx.NewStyle(FILL)

	return title, subtitle, text, bold, fill
}
func SetExcelIntro(xlsx *excelize.File, intro string, c int, style int) {
	xlsx.SetSheetName("Sheet1", "Summary")
	xlsx.SetCellValue("Summary", "A1", intro)
	col := excelize.ToAlphaString(c)
	xlsx.MergeCell("Summary", "A1", col+"1")
	xlsx.SetCellStyle("Summary", "A1", col+"1", style)
}

func SaveExcel(xlsx *excelize.File, t int, reps int) {
	var sheet, filename string
	switch t {
	case 1:
		sheet = "Summary"
		filename = "/out/combinations/" + "combinations" + strconv.Itoa(WORDS) + "x" + strconv.Itoa(GROUP) + "_" + strconv.Itoa(reps) + "_" + "element_repetitions.xlsx"
	case 2:
		sheet = "Summary"
		filename = "/out/test/favs/" + "favorables" + strconv.Itoa(WORDS) + "x" + strconv.Itoa(GROUP) + "_" + strconv.Itoa(reps) + "_" + "element_repetitions.xlsx"
	case 3:
		sheet = "Graph"
		filename = "/out/test/deps/" + "dependence" + strconv.Itoa(WORDS) + "x" + strconv.Itoa(GROUP) + ".xlsx"
	case 4:
	}
	index := xlsx.GetSheetIndex(sheet)
	xlsx.SetActiveSheet(index)
	path, _ := os.Getwd()
	xlsx.SaveAs(path + filename)
}

func AddDependenceChart(xlsx *excelize.File, chartdata map[string][][]int) {
	xlsx.NewSheet("Graph")
	var serie []string
	f := 2

	xlsx.SetColVisible("Graph", "A", false)
	for c := 0; c < GROUP; c++ {
		col := excelize.ToAlphaString(c + 1)
		xlsx.SetCellValue("Graph", col+"1", c)
		xlsx.SetColVisible("Graph", col, false)
	}
	coli := excelize.ToAlphaString(1)
	colf := excelize.ToAlphaString(GROUP)

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

	colf = excelize.ToAlphaString(GROUP + 2)
	xlsx.AddChart("Graph", colf+"2",
		`{
	"type":"col",
	"series":`+`[`+series+`]`+`,
	"title":
		{"name":"Dependence between events"}
	}`)
}
