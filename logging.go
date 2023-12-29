package main

import (
	"context"
	"log"
	"net"
	"time"
)

type LoggingService struct {
	next Socket	
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
