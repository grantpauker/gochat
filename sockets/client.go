package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Client struct {
	ip         string
	port       string
	connection net.Conn
	errConn    error
	errSend    error
}

func inputStream(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
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
	client.errConn = send(str, n, client.connection)
	if client.errConn != nil {
		log.Fatalln(client.errConn)
	}
}
func main() {
	var (
		str string
	)
	for {
		newClient := createClient("127.0.0.1", "3333")
		newClient.Connect()
		str = inputStream("Send: ")
		newClient.Send(str, 255)
	}

}
