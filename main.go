package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Id      int
	Content string
	Done    bool
}

var todos []Todo
var nextId = 1

func main() {
	reader := bufio.NewReader(os.Stdin)
	help()
	for {
		fmt.Print("Choose: ")

		selected, _ := reader.ReadString('\n')
		selected = strings.TrimSpace(selected)

		fmt.Println(selected)

		switch selected {
		case "0":
			help()
		case "2":
			add(reader)
		default:
			fmt.Println("You select wrong input")
		}
	}
}

func help() {
	fmt.Println("\n===========================")
	fmt.Println("       TODO CLI")
	fmt.Println("===========================")
	fmt.Println("0. Help")
	fmt.Println("1. List tasks")
	fmt.Println("2. Add task")
	fmt.Println("3. Mark task done")
	fmt.Println("4. Delete task")
	fmt.Println("5. Exit")
}

func add(reader *bufio.Reader) {
	fmt.Println("Add text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	// validate

	id := nextId
	todos = append(todos, Todo{
		Id:      id,
		Content: text,
	})
	nextId++

	fmt.Printf("Todo Added! with content (%d:%s)\n", id, text)
}
