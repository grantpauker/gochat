package main

import (
	"bufio"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type User struct {
	name  string
	color [3]int
}
type Message struct {
	user User
	text string
}

func createRGB(r int, g int, b int) [3]int {
	return [3]int{r, g, b}
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
func newUser(user string, rgb [3]int) User {
	return User{name: user, color: rgb}
}

func newMessage(user User, text string) Message {
	return Message{user: user, text: text}
}

func appendFile(user User, text string, fileName string) {
	f, _ := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	f.WriteString(user.name + "" + intArray2Str(user.color[:]) + "" + text + "\n")
	f.Close()
}
func createChatMap(filename string) []Message {
	var tmp []string
	var m []Message

	b, _ := os.Open(filename)
	scanner := bufio.NewScanner(b)
	var newIntArray [3]int
	for scanner.Scan() {
		tmp = strings.Split(scanner.Text(), "")
		if len(tmp) > 2 {
			intArray := str2IntArray(tmp[1])

			copy(newIntArray[:], intArray[:3])
			m = append(m, newMessage(newUser(tmp[0], newIntArray), tmp[2]))
		}
	}
	return m
}
func printChatMap(m []Message) {
	for _, theMessage := range m {
		printRGB(theMessage.user.name+": "+theMessage.text+"\n", theMessage.user.color)
	}
}
func str2IntArray(str string) []int {
	split := strings.Split(str, ",")
	var intArray []int
	var val int
	for _, item := range split {
		val, _ = strconv.Atoi(item)
		intArray = append(intArray, val)
	}
	return intArray
}
func intArray2Str(intArray []int) string {
	var str string
	for _, item := range intArray {
		str = str + strconv.Itoa(item) + ","
	}
	str = strings.TrimSuffix(str, ",")
	return str
}
