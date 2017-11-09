package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

type Server struct {
	ip         string
	port       string
	listener   net.Listener
	lErr       error
	connection net.Conn
	cErr       error
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func onClose(toRun interface{}) {

	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTSTP)
	go func() {
		<-c
		fn := reflect.ValueOf(toRun)
		fn.Call([]reflect.Value{})
		os.Exit(1)
	}()
}
func createServer(ip string, port string) *Server {
	server := &Server{ip: ip, port: port}
	return server
}
func send(text string, n int, conn net.Conn) error {
	var err error
	if len(text) < n {
		_, err = conn.Write([]byte(text))
	} else {
		err = errors.New(fmt.Sprintf("String must be less than length %d", n+1))
	}
	return err
}
func (server *Server) Listen() {
	server.listener, server.lErr = net.Listen("tcp", fmt.Sprintf("%s:%s", server.ip, server.port))
}
func (server *Server) Accept() {
	server.connection, server.cErr = server.listener.Accept()
}
func (server *Server) Read(buffer []byte) {
	server.connection.Read(buffer)
}
func (server *Server) Send(str string, n int) {
	server.cErr = send(str, n, server.connection)
	if server.cErr != nil {
		log.Fatalln(server.cErr)
	}
}
func (server *Server) Close() {
	server.listener.Close()
}
