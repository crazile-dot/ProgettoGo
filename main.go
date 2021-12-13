package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func somma(ch chan int, uno int, due int) chan int {
	ch <- uno + due
	return ch
}

func main() {
	var tot int = 0
	var chans [4]chan string
	for i := range chans {
		chans[i] = make(chan string)
	}
	file, err := os.Open("C:\\Users\\Ilenia\\Desktop\\prova.txt.txt")
	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	var splittedFile [1000]string
	for fileScanner.Scan() {
		splittedFile[tot] = fileScanner.Text()
		tot = tot + 1
	}
	for i := 0; i < 4; i++ {
		go worker(chans[i], splittedFile.
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	fmt.Println("in totale il file ha", tot, "righe")
	file.Close()
}
