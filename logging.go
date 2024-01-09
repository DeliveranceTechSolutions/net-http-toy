package main

import (
	"context"
	"log"
	"net"
	"time"
)

// Any new interfaces will be embedded into ILogger
type ILogger interface {
	IClient
	IServer
	ISocket
}

type LoggingService struct {
	next ILogger
}

// use a descriptive name for each individual LogginService, to identify the "base" type
func (socket *LoggingService) Subscribe(ctx context.Context, target string) (conn net.Conn, err error) {
	defer func(start time.Time) {
		log.Printf("Connection: %#v, Error: %#v", conn, err)	
	}(time.Now())
		
	return socket.next.Subscribe(ctx, target)
}

// it could be lowercase, whatever seems right to yourself
func (socket *LoggingService) Disconnect(ctx context.Context) (confirm string, err error) {
	defer func(start time.Time) {
		log.Printf("%s Error: %#v", confirm, err)
	}(time.Now())
	
	return socket.next.Disconnect(ctx)
}
