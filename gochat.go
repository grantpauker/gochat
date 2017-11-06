package main

import (
	"fmt"
)

func main() {
	rgb := [3]int{100, 0, 255}
	me := newUser("Grant", rgb)
	clear()
	dir := getMessageDir()
	file := pickMessage(dir)
	path := getMessageDir() + "/" + file
	clear()
	m := createChatMap(path)
	var msg string
	for {

		fmt.Printf("--%s--\n", file)
		printChatMap(m)
		msg = inputStream(fmt.Sprintf("Message (%s): ", file))
		if msg == "" {
			clear()
			continue
		}
		m = append(m, newMessage(me, msg))
		appendFile(me, msg, path)
		clear()
	}

}
