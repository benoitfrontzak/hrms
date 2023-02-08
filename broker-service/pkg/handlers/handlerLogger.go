package handlers

import (
	"log"
	"net/http"
	"net/rpc"
)

var logger = "logger-service:5001"

// logItemViaRPC logs an item by making an RPC call to the logger microservice
func (rep *Repository) LogItemViaRPC(w http.ResponseWriter, l RPCPayload) error {
	// Dial up RPC server
	client, err := rpc.Dial("tcp", logger)
	if err != nil {
		log.Println("Client failed to dial")
		// app.errorJSON(w, err)
		return err
	}

	// Call RPC server to run LogInfo and return result
	var result string
	err = client.Call("RPCServer.LogInfo", l, &result)
	if err != nil {
		log.Println("error is here", err)
		// app.errorJSON(w, err)
		return err
	}

	log.Println("Mongo logged notification successfully", result)
	return nil
}
