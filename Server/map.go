package main

import (
	"ProgettoGo/utils"
	"log"
	"strconv"
	"strings"
)

func work(input []string, word string) utils.Result {
	values := mapPhase(input, word)
	shuffleAndSort := shuffleAndSortPhase(values)
	reduce := reducePhase(shuffleAndSort)

	countValues := mapPhaseCount(input, word)
	shuffleAndSortCount := shuffleAndSortPhaseCount(countValues)
	reduceCount := reducePhaseCount(shuffleAndSortCount)

	output := reduce[word]
	outCount := reduceCount[word]
	result := utils.Result{Line: output, Num: outCount}
	return result
}

// scrive in una map tutte le occorrenze della parola cercata assegnandole il valore 1
//Questa è la base per contare quante volte quella parola ricorre nel testo
func mapPhaseCount(input []string, find string) [][]string {
	var out [][]string
	for _, row := range input {
		// words è un array di stringhe, contiene tutte le parole di una riga
		words := strings.Split(row, " ")
		for _, w := range words {
			if w == find {
				var elem []string
				elem = append(elem, w)
				elem = append(elem, "1")
				out = append(out, elem)
			}
		}
	}
	return out
}

func shuffleAndSortPhaseCount(values [][]string) map[string][]string {
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

func reducePhaseCount(values map[string][]string) map[string]int {
	out := make(map[string]int)
	outCount := 0
	for key, value := range values {
		for _, num := range value {
			conv, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Error while converting string to int %s", err)
			}
			outCount += conv
		}
		out[key] = outCount
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
