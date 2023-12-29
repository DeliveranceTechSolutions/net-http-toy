package main

import (
	"context"
	"log"
	"net"
	"time"
)

type LoggingService struct {
	next Interface
}

// Publish implements Interface.
func (ls *LoggingService) Publish(ctx context.Context, msg string) (err error) {
	defer func(start time.Time) {
		log.Printf("%s Error: %#v", msg, err)
	}(time.Now())

	return ls.next.Publish(ctx, msg)
}

// Retry implements Interface.
func (*LoggingService) Retry(context.Context) {
	panic("unimplemented")
}

func NewLoggingService(i Interface) *LoggingService {
	return &LoggingService{
		next: i,
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
