package main

import "fmt"

func main() {
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
}
