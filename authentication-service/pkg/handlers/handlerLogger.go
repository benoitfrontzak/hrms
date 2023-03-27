package handlers

import (
	"net/rpc"
)

// logItemViaRPC logs an item by making an RPC call to the logger microservice
func (rep *Repository) LogItemViaRPC(l RPCPayload) error {
	// Dial up RPC server
	client, err := rpc.Dial("tcp", logger)
	if err != nil {
		return err
	}

	// Call RPC server to run LogInfo and return result
	var result string
	err = client.Call("RPCServer.LogInfo", l, &result)
	if err != nil {
		return err
	}

	return nil
}
