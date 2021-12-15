package main

import (
	"fmt"
	"strings"
)

func work(input []string) {
	find := "daniele"
	values := mapPhase(input, find)
	shuffleAndSort := shuffleAndSortPhase(values)
	reduce := reducePhase(shuffleAndSort)

	output := reduce[find]

	fmt.Print(output)
}

// scrive in una map tutte le occorrenze della parola cercata assegnandole il valore 1
//Questa è la base per contare quante volte quella parola ricorre nel testo
func mapPhaseDummy(input string, find string) map[string]int {
	out := make(map[string]int)
	word := strings.Split(input, " ")
	for _, s := range word {
		if s == find {
			out[s] = 1
		} else {
			continue
		}
	}
	return out
}

//MA a noi serve la grep, che ritorna le righe che contengono quella parola (o espressione)
// quindi per prima cosa il master mi deve dare una lista di righe
func mapPhase(input []string, find string) [][]string {
	//out := make(map[string]string)
	var out [][]string
	for _, row := range input {
		var elem []string
		// words è un array di stringhe, contiene tutte le parole di una riga
		words := strings.Split(row, " ")
		for _, w := range words {
			if w == find {
				elem = append(elem, w)
				elem = append(elem, row)
				out = append(out, elem)
				break
			}
		}
	}
	return out
}

func shuffleAndSortPhase(values [][]string) map[string][]string {
	out := make(map[string][]string)
	for _, elem := range values {
		key := elem[0]
		var slice []string
		if out[key] == nil {
			slice = append(slice, elem[1])
			out[key] = slice
		} else {
			out[key] = append(out[key], elem[1])
		}
	}
	return out
}

func reducePhase(values map[string][]string) map[string]string {
	out := make(map[string]string)
	outStr := ""
	for key, value := range values {
		for _, str := range value {
			outStr += "\n" + str
		}
		out[key] = outStr
	}
	return out
}
