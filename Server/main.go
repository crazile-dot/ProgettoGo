package main

import (
	"ProgettoGo/utils"
	"bufio"
	"log"
	"os"
	"time"
)

var path = "/home/ec2-user/ProgettoGo/Novecento_Un_monologo.txt"

func startSplit(word string) *utils.Result {
	var N int = 4
	var j int = 0
	var tot int = 0
	var countRes int = 0
	var strRes string = ""
	result := new(utils.Result)
	var chans []chan utils.Result

	for i := 0; i < N; i++ {
		chans = append(chans, make(chan utils.Result))
	}
	file, err := os.Open(path)

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
		time.Sleep(2000)
		batch = splittedFile[j:len]
		go worker(chans[i], batch, i, word)

		j = len
		partial := <-chans[i]
		countRes += partial.Num
		strRes += "\n" + partial.Line

		result.Line = strRes
		result.Num = countRes
	}

	//handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}
	file.Close()

	return result
}

func worker(c chan utils.Result, splitted []string, num int, word string) {
	result := work(splitted, word)
	c <- result
}

func main() {
	ServerConnection()
}
