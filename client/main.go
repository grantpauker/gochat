package main

import (
	"fmt"
)

var me = "Grant"

func main() {
	clear()

	newClient := createClient("108.70.148.10", "3333")
	dir := getMessageDir()
	file := pickMessage(dir)
	path := getMessageDir() + "/" + file
	clear()
	m := createChatMap(path)
	var msg string
	fmt.Printf("--%s--\n", file)
	printChatMap(m)

	for {
		newClient.Connect()
		msg = inputStream(fmt.Sprintf("Message (%s): ", file))
		fmt.Printf("\033[1A\033[2K")
		if msg == "" {
			continue
		}
		m = append(m, newMessage(me, msg))
		appendFile(me, msg, path)
		printMessage(newMessage(me, msg))
		newClient.Send(fmt.Sprintf("%s%s", me, msg), 255)

		// buff := make([]byte, 255)
		// newClient.Read(buff)
		// fmt.Printf("Receive: %s\n", string(buff))

	}

}
