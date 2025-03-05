package main

import (
	"bufio"
	"fmt"
	"github.com/0ne290/go-tasks/task4/internal"
	"os"
	"strings"
)

func main() {
	stdinReader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text separated by commas: ")
	source, err := stdinReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	source = strings.TrimSpace(source)

	stack := internal.NewStack[string]()
	queue := internal.NewQueue[string]()

	for _, text := range strings.Split(source, ",") {
		stack.Push(text)
		queue.Enqueue(text)
	}

	var queueString strings.Builder
	gh := ""// Как же я устал придумывать имена переменным
	if !queue.IsEmpty() {
		

		for !queue.IsEmpty() {
			text, _ := queue.Dequeue()
			queueString.WriteString(text + ", ")
		}

		gh = queueString.String()
		gh = gh[:len(gh)-2]
	}

	var stackString strings.Builder
	hg := ""// Как же я устал придумывать имена переменным
	if !stack.IsEmpty() {
		for !stack.IsEmpty() {
			text, _ := stack.Pop()
			stackString.WriteString(text + ", ")
		}

		hg = stackString.String()
		hg = hg[:len(hg)-2]
	}

	fmt.Println("Source text: " + source)
	fmt.Println("Queue iteration: " + gh)
	fmt.Println("Stack iteration: " + hg)
}
