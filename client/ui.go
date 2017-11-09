package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func clear() {
	fmt.Printf("\033c")
}

func inputStream(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}

func pickMessage(dir string) string {
	messages := filesInDir(dir)
	fmt.Println("Pick a chat:")

	for i, message := range messages {
		fmt.Printf(" [%d] %s\n", i+1, message)
	}
	var err error
	var choice int
	for {
		choice, err = strconv.Atoi(inputStream("Choice: "))
		if err == nil && 1 <= choice && choice <= len(messages) {
			break
		}
		fmt.Printf("Please pick a number between 1 and %d\n", len(messages))
	}
	return messages[choice-1]
}
