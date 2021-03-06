package main

import (
	"ProgettoGo/utils"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type API int

//esempio di RPC
func (a *API) MasterAction(arg1 string, reply *utils.Result) error {
	var temp = startSplit(arg1)
	reply.Line = temp.Line
	reply.Num = temp.Num
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
