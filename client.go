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
