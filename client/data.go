package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Message struct {
	user string
	text string
}

func getMessageDir() string {
	_, dir, _, _ := runtime.Caller(0)
	dir = filepath.Dir(dir) + "/messages"
	return dir
}
func filesInDir(dir string) []string {
	files, _ := ioutil.ReadDir(dir)
	var fileNames []string

	for _, item := range files {
		fileNames = append(fileNames, item.Name())
	}
	return fileNames
}

func newMessage(user string, text string) Message {
	return Message{user: user, text: text}
}

func appendFile(user string, text string, fileName string) {
	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)

	f.WriteString(user + "" + text + "\n")

	f.Close()
}
func createChatMap(filename string) []Message {
	var tmp []string
	var m []Message

	b, _ := os.Open(filename)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		tmp = strings.Split(scanner.Text(), "")
		if len(tmp) > 1 {
			m = append(m, newMessage(tmp[0], tmp[1]))
		}
	}
	return m
}
func printChatMap(m []Message) {
	for _, theMessage := range m {
		fmt.Println(theMessage.user + ": " + theMessage.text)
	}
}
