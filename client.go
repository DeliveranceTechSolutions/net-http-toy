package main

import (
	"github.com/google/uuid"
)

type IClient interface {
	NewCoreClient()
}

type Client struct {
	// db
	id uuid.UUID
	socket ISocket
	connection bool
	partner uuid.UUID
}

func NewCoreClient(is ISocket) *Client {
	return &Client{
		id: uuid.New(),
		socket: is,
		connection: false,
		partner: uuid.Nil,
	}
}