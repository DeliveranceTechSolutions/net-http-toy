package main

import (
	"net"
	"net/http"

	"github.com/google/uuid"
)

type IServer struct {
	NewServer(string)
	ServeHTTP(interface{})
	RequestedClientConnection(uuid.UUIDs)

}

type Server struct {
	name string
	prot string
	port string
	url string
	pooledConnections map[uuid.UUID]uuid.UUID
	//Heartbeat so server kills the sockets not the client
	// will add later
}

func NewServer(name, prot, port, url string) {
	s := &Server{
		name: name, 
		prot: prot,
		port: port,
		url:  url,
		pooledConnections: make(map[uuid.UUID]uuid.UUID),
	}

	s = NewLoggingService(s)
	s.startServer()
	return
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var handler http.Handler
	http.Handler.ServeHTTP(handler, w, r)
	return
}

func (s *Server) startServer() { 
	http.ListenAndServe(s.url + ":" + s.port, s) 
	return
}

func (s *Server) RequestedClientConnection(requester, recipient uuid.UUID) (bool, error) {

}