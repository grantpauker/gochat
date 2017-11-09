package main

import (
	"fmt"
)

func main() {
	me := "Grant"
	clear()
	dir := getMessageDir()
	file := pickMessage(dir)
	path := getMessageDir() + "/" + file
	clear()
	m := createChatMap(path)
	var msg string
	newClient := createClient("127.0.0.1", "3333")
	for {
		newClient.Connect()

		fmt.Printf("--%s--\n", file)
		printChatMap(m)
		msg = inputStream(fmt.Sprintf("Message (%s): ", file))
		if msg == "" {
			clear()
			continue
		}
		m = append(m, newMessage(me, msg))
		appendFile(me, msg, path)
		newClient.Send(fmt.Sprintf("%s%s", me, msg), 255)
		clear()
		buff := make([]byte, 255)
		newClient.Read(buff)
		fmt.Printf("Receive: %s\n", string(buff))

	}

}
