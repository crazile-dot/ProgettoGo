package main

import (
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
	var result *string
	fmt.Print("Inserire la parola da cercare:\n")
	fmt.Scanf("%s", &wordToSearch)
	err := prod.Call("API.MasterAction", wordToSearch, result)
	//gestiamo l'errore da parte del server, sarebbe da inserire anche un timeout
	if err != nil {
		return
	}
}

func main() {
	con := ClientConnection()
	Grep(con)
}
