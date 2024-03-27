package main

import (
	"context"
	"log"
	"log-service/data"
)

type RpcServer struct {
}
type RpcPayload struct {
	Name string
	Data string
}

func (r *RpcServer) LogInfo(payload RpcPayload, resp *string) error {
	log.Println("Received new log rpc call")
	collection := data.GetLoggingCollection()

	_, err := collection.InsertOne(context.TODO(), data.LogEntry{
		Name: payload.Name,
		Data: payload.Data,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	*resp = "Processed payload via RPC: " + payload.Data
	return nil
}
