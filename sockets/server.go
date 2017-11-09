package main

import (
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
func creatServer(ip string, port string) *Server {
	server := &Server{ip: ip, port: port}
	return server
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

func (server *Server) Close() {
	server.listener.Close()
}

func main() {
	newServer := creatServer("", "3333")
	newServer.Listen()
	checkError(newServer.lErr)
	onClose(newServer.Close)
	fmt.Printf("Connected on port %s...\n", newServer.port)
	for {
		tmp := make([]byte, 256)
		newServer.Accept()
		checkError(newServer.cErr)
		newServer.Read(tmp)
		fmt.Println(fmt.Sprintf("Recieved: %s", string(tmp)))
	}
	newServer.Close()
}
