package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

func somma(ch chan int, uno int, due int) chan int {
	ch <- uno + due
	return ch
}

func main() {
	var N int = 4
	var j int = 0
	var tot int = 0
	var chans []chan string
	for i := 0; i < N; i++ {
		chans = append(chans, make(chan string))
	}
	file, err := os.Open("C:\\Users\\Ilenia\\GolandProjects\\ProgettoGo\\prova.txt.txt")
	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	fileScanner := bufio.NewScanner(file)
	var splittedFile []string
	for fileScanner.Scan() {
		splittedFile = append(splittedFile, fileScanner.Text())
		tot++
	}

	for i := 0; i < N; i++ {
		var batch []string
		len := j + ((len(splittedFile) - j) / (N - i))
		time.Sleep(1000)
		batch = splittedFile[j:len]
		go worker(chans[i], batch, i)
		j = len
	}

	//handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	//fmt.Println("in totale il file ha", tot, "righe")
	file.Close()
}

func worker(c chan string, splitted []string, num int) {
	//fmt.Println("Sono il worker numero ", num, " ed ho queste righe: ", splitted)
	work(splitted)
}
