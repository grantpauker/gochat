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

type User struct {
	name string
}
type Message struct {
	user User
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
func newUser(user string) User {
	return User{name: user}
}

func newMessage(user User, text string) Message {
	return Message{user: user, text: text}
}

func appendFile(user User, text string, fileName string) {
	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)

	f.WriteString(user.name + "" + text + "\n")

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
			m = append(m, newMessage(newUser(tmp[0]), tmp[1]))
		}
	}
	return m
}
func printChatMap(m []Message) {
	for _, theMessage := range m {
		fmt.Println(theMessage.user.name + ": " + theMessage.text)
	}
}
