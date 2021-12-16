package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

type Result struct {
	Line string
	Num  int
}

//esempio di RPC
func (a *API) MasterAction(arg1 string, reply *Result) error {
	reply.Line = startSplit(arg1)
	return nil
}

func ServerConnection() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registering API ", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listening error ", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving ", err)
	}
}
