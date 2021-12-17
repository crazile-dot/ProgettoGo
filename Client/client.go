package main

import (
	"ProgettoGo/utils"
	"fmt"
	"log"
	"net/rpc"
)

var tcp = "tcp"
var host = "localhost:4040"

func ClientConnection() *rpc.Client {
	prod, err := rpc.DialHTTP(tcp, host)
	if err != nil {
		log.Fatal("connection error", err)
	}
	fmt.Println("Connesso")
	return prod
}

func Grep(prod *rpc.Client) {
	var wordToSearch string
	var res utils.Result
	fmt.Println("Inserire la parola da cercare: ")
	fmt.Scanf("%s", &wordToSearch)
	err := prod.Call("API.MasterAction", wordToSearch, &res)
	//gestiamo l'errore da parte del server, sarebbe da inserire anche un timeout
	if err != nil {
		print(err.Error())
		return
	}
	fmt.Println("La parola che cerchi appare ", res.Num, " volte nelle seguenti righe: ")
	fmt.Println(res.Line)

}

func main() {
	con := ClientConnection()
	Grep(con)
}
