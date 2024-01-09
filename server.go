package main

import (
	"context"
	"fmt"
)

type IServer interface {
	Start()
}

type Server struct {
	sender ClientConfig
	receiver ClientConfig

}

type ClientConfig struct {
	protocol string
	address string
	port string
}

func (s *Server) Start() {
	ctx := context.Background()
	sock1, err := NewSocketConn(
		s.sender.protocol,
		s.sender.address,
		s.sender.port,
	)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	sock2, err := NewSocketConn(
		s.receiver.protocol,
		s.receiver.address,
		s.receiver.port,
	)
	if err != nil {
		panic(fmt.Errorf(err.Error()))
	}
	sock1.Subscribe(ctx, s.receiver.address + s.receiver.port)
	sock2.Subscribe(ctx, s.receiver.address + s.receiver.port)
	return 
}

func NewCoreServer() IServer {
	client1 := ClientConfig{
		"tcp",
		"localhost",
		"9988",
	}
	client2 := ClientConfig{
		"tcp",
		"localhost",
		"9987",
	}

	return &Server{
		sender: client1,
		receiver: client2,
	}
}
