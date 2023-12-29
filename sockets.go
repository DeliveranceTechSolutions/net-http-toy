package main

import (	
	"context"
	"fmt"
	"net"
	"os"
)	

type Interface interface {
	Subscribe(context.Context) (conn net.Conn, err error)
	Publish(context.Context, string) (err error)
	Disconnect(context.Context) (confirm string, err error)
	Retry(context.Context) 	
}
	
type Socket struct {
	conn net.Conn
	listener net.Listener
}

func NewSocketConn(prot, host, port string) (Interface, error) {
	listener, err := net.Listen(prot, host + ":" + port)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	conn, err := net.Dial(prot, host + ":" + port)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	
	return &Socket{
		conn: conn,
		listener: listener,

	}, nil
}

func (s *Socket) Subscribe(ctx context.Context) (conn net.Conn, err error) {
	return s.listener.Accept()
}

func (s *Socket) Publish(ctx context.Context, msg string) (err error) {
	_, err = os.Stdout.Write([]byte(msg + "\n"))
	if err != nil {
		fmt.Println(fmt.Errorf("error: Publish() - %s", err))
	}	

	return
}

func (s *Socket) Disconnect(ctx context.Context) (confirm string, err error){
	if err := s.listener.Close(); err != nil {
		return "", err	
	}

	return "Successfully closed socket", nil
}

func (s *Socket) Retry(ctx context.Context) {
	return 
}
