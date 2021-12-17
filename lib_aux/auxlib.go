package lib_aux

const (

	/* casos totals (disjunts i no disjunts)
	 * Tipus {1,2,3}|{4,5,6}, {1,2,3}|{5,6,7}, {4,7,8}|{1,2,3}, ...
	 * Tipus {1,2,3}|{1,4,5}, {1,2,3}|{4,2,5}, {1,2,3}|{4,5,3}, ...
	 * Tipus {1,2,3}|{1,2,5}, {1,2,3}|{1,4,3}, {1,2,3}|{4,2,3}, ...
	 */

	CASES = 1

	WORDS = 8
	GROUP = 3
)

//Estructura que conté el grup de GROUP elements i un random id
//per saber de quina combinació es tracta i fer mes entendible l'arxiu resultant
type Combi struct {
	Id    int
	Group [lib_aux.GROUP]int
	Value [lib_aux.GROUP]int
}

func RemoveSlice(original []int, g []int) []int {
	var remaining []int
	remaining = append(remaining, original[:]...)

	for i, elem := range g {
		RemoveIndex(remaining, elem-i-1)
	}

	return remaining[:WORDS-GROUP]
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func Separable(group1 [lib_aux.GROUP]int, group2 [lib_aux.GROUP]int) bool {
	first := make(map[int]int)
	second := make(map[int]int)

	for _, v := range group1 {
		first[v] = v
	}

	for _, v := range group2 {
		second[v] = v
	}

	return len(first) != len(second)

}
