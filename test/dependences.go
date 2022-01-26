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

// The local function getArraypairs returns a map with an array filled
// by pairs of groups where the key is a string composed by the number
// of repetitions in the pair and its dimensions. One example would be
// the key = 233 for a pair like [1 2 3],[1 2 5]. In resume, all the
// pairs in the array would be the same type.

func getArrayPairs(initial []int) map[string][][]int {
	// Local variables initialization
	var array []map[int][]lib.Map
	arraypairs := make(map[string][][]int)

	// Use of the combinations.List function to get the differents
	// arraymaps with the combinations depending on the repetitions.
	// Sort() categorises the arraymap by the length of the top group.
	for i := 0; i < lib.REPS; i++ {
		arraymap := combinations.List(initial, i)
		array = append(array, lib.Sort(arraymap))
	}

	// The variable array is an array of maps. The position 0
	// belongs to the map with elements of zero repetitions in the
	// pairs, the position 1 to the groups that have one repetition
	// between them, etc...

	for reps, arraymaps := range array {
		// arraymaps is a map whose key is the first group length. It
		// is the object returned previously by the function Sort().
		// The loop goes through all the combinations for a fixed first
		// group length
		for k1, am := range arraymaps {
			// Set a Map with an initial group and its combinations
			// sort also by length
			for _, m := range am {
				// Set an array from the m.Seconds map
				for k2, ss := range m.Seconds {
					// Go through the array of the seconds groups for
					// a fixed initial group and length
					for _, s := range ss {
						var first, second, group []int
						key := strconv.Itoa(reps) + strconv.Itoa(k1) + strconv.Itoa(k2)
						first = append(first, m.First...)
						second = append(second, s...)
						group = append(first, second...)
						// Add the pair if it is not in the array
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
	// Excel initialization
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

	// Local variables initialization
	var results []string
	chartdata := make(map[string][][]int)

	for k, arraypair := range arraypairs {
		// Pick one array from the arraypair each key type
		array := arraypair[0]
		var deps, nodeps, total int
		// Compare the pair chosen with al the other pairs
		for keyToCompare, arrayToCompare := range arraypairs {
			var init, fin int
			// If the pair is from the same array (same type key),
			// the length should be len(arrayToCompare)-1
			if k == keyToCompare {
				init, fin = 1, 1
			} else {
				init, fin = 0, 0
			}
			// Loop to compare each pack of groups with the main pair.
			// It is applied to the suitable length.
			for i := 0; i < len(arrayToCompare)-fin; i++ {
				index := (init + i) % (len(arrayToCompare))
				// The function Dependent returns a bool that is true
				// if the two arrays passed by argument are related.
				if lib.Dependent(array, arrayToCompare[index]) {
					deps++
				} else {
					nodeps++
				}
			}
			// Total sums all pairs (events)
			total += len(arrayToCompare)
		}

		// The data is printed here for each type of pairs
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

	// Add chart
	lib.AddDependenceChart(xlsx, chartdata)
	// Save excel
	lib.SaveExcel(xlsx, 3, 0)
}
