package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
)

type ILogger interface {
	ISocket | IClient | IServer
}

type LoggingService struct {
	next 
}

// Publish implements ISocket.
func (ls *LoggingService) Publish(ctx context.Context, msg string) (err error) {
	defer func(start time.Time) {
		log.Printf("%s Error: %#v", msg, err)
	}(time.Now())

	return ls.next.Publish(ctx, msg)
}

// Retry implements ISocket.
func (*LoggingService) Retry(context.Context) {
	panic("unimplemented")
}

func NewLoggingService(i) *LoggingService {
	return &LoggingService{
		next: t,
	}
}

func (ls *LoggingService) Subscribe(ctx context.Context) (conn net.Conn, err error) {
	defer func(start time.Time) {
		log.Printf("Connection: %#v, Error: %#v", conn, err)
	}(time.Now())

	return ls.next.Subscribe(ctx)
}

func (ls *LoggingService) Disconnect(ctx context.Context) (confirm string, err error) {
	defer func(start time.Time) {
		log.Printf("%s Error: %#v", confirm, err)
	}(time.Now())

	return ls.next.Disconnect(ctx)
}

func (ls *LoggingService) RequestedClientConnection(requester, recipient uuid.UUID) (success bool, err error) {
	defer func(start time.Time) {
		log.Printf("%s Error: %#v", requester, recipient, success, err)
	}(time.Now())

	return ls.next.RequestedClientConnection(requester, recipient)
}
