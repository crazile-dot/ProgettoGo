package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Result struct {
	Line string
	Num  int
}

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
	var res Result
	fmt.Print("Inserire la parola da cercare:\n")
	fmt.Scanf("%s", &wordToSearch)
	err := prod.Call("API.MasterAction", wordToSearch, &res)
	//gestiamo l'errore da parte del server, sarebbe da inserire anche un timeout
	if err != nil {
		print(err.Error())
		return
	}
	print("la parola che cerchi appare nelle seguenti righe: \n")
	print(res.Line)

}

func main() {
	con := ClientConnection()
	Grep(con)
}
