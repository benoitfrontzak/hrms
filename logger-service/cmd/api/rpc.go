package main

import (
	"context"
	"fmt"
	"log"
	"logger/pkg/mongoDB"
	"net"
	"net/rpc"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RPCServer is the type for our RPC Server. Methods that take this as a receiver are available
// over RPC, as long as they are exported.
type RPCServer struct{}

// RPCPayload is the type for data we receive from RPC
type RPCPayload struct {
	Collection string
	Name       string
	Data       string
	CreatedAt  string
	CreatedBy  string
}

// LogInfo writes our payload to mongo DB
func (r *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	collection := mongoDB.Client.Database(mongoDB.DB).Collection(payload.Collection)
	result, err := collection.InsertOne(context.TODO(), mongoDB.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: payload.CreatedAt,
		CreatedBy: payload.CreatedBy,
	})

	if err != nil {
		log.Println("error writing to mongo", err)
		return err
	}

	// convert *mongo.InsertOneResult (result) to primitive.ObjectID (oid)
	oid, _ := result.InsertedID.(primitive.ObjectID)
	log.Println("rpc server successfully logged the authentication with id: ", oid.Hex())

	// resp is the message sent back to the RPC caller
	*resp = "Processed payload via RPC:" + payload.Name
	return nil
}

// rpcListen announces on the local network address (TCP).
// it listens on all available unicast and anycast IP addresses of the local system.
func rpcListen() error {

	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%s", rpcPort))
	if err != nil {
		return err
	}
	defer listen.Close()

	log.Println("RPC server listening on port ", rpcPort)

	for {
		rpcConn, err := listen.Accept()
		if err != nil {
			continue
		}
		go rpc.ServeConn(rpcConn)
	}

}
