package main

import (
	"errors"
	"fmt"
	"log"
	"net"
)

/*
var (
	str string
)
for {
	newClient := createClient("127.0.0.1", "3333")
	newClient.Connect()
	str = inputStream("Send: ")
	newClient.Send(str, 255)
	buff := make([]byte, 255)
	newClient.Read(buff)
	fmt.Printf("Receive: %s\n", string(buff))
}
*/
type Client struct {
	ip         string
	port       string
	connection net.Conn
	errConn    error
	errSend    error
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
func createClient(ip string, port string) *Client {
	client := &Client{ip: ip, port: port}
	return client
}

func (client *Client) Connect() {
	client.connection, client.errConn = net.Dial("tcp", fmt.Sprintf("%s:%s", client.ip, client.port))
	if client.errConn != nil {
		log.Fatalln(client.errConn)
	}
}
func (client *Client) Send(str string, n int) {
	client.errSend = send(str, n, client.connection)
	if client.errSend != nil {
		log.Fatalln(client.errSend)
	}
}
func (client *Client) Read(buffer []byte) {
	client.connection.Read(buffer)
}
