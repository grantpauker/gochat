package main

import (
	"fmt"
)

func main() {
	newServer := createServer("", "3333")
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
		newServer.Send(string(tmp), 270)
	}
}
