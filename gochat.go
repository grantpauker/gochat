package main

import "fmt"

func main() {
	// str := "3,3,3100,3123,123123,31123"
	// intArray := str2IntArray(str)
	// fmt.Println(intArray)
	// fmt.Println(intArray2Str(intArray))
	me := newUser("Grant", createRGB(255, 0, 100))
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
