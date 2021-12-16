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
	fmt.Println("Connessione riuscita:")
	return prod
}

func Grep(prod *rpc.Client) {
	var wordToSearch string
	var result *string
	fmt.Scanf("%s", &wordToSearch)
	prod.Call("API.MasterAction", wordToSearch, result)
}

func main() {
	con := ClientConnection()
	Grep(con)
}
