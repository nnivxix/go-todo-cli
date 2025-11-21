package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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
		case "1":
			list()
		case "2":
			add(reader)
		case "3":
			toggleMarkAsDone(reader)
		case "4":
			deleteTodo(reader)
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

func list() {
	if len(todos) == 0 {
		fmt.Println("No Todos Avalaible")
		return
	}
	fmt.Println("Todos:")

	for _, todo := range todos {
		status := " "
		if todo.Done {
			status = "x"
		}

		fmt.Printf("[%s]%d - %s\n", status, todo.Id, todo.Content)
	}

}

func add(reader *bufio.Reader) {
	fmt.Println("Add TODO: ")
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

func toggleMarkAsDone(reader *bufio.Reader) {
	fmt.Println("Select TODO id")

	// get id
	idStr, _ := reader.ReadString('\n')
	fmt.Println(idStr)
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for i := range todos {
		if todos[i].Id == id {
			todos[i].Done = !todos[i].Done
			todo := todos[i]
			if todo.Done {
				fmt.Printf("Todo marked as done (%d:%s)\n", todo.Id, todo.Content)
			} else {
				fmt.Printf("Todo marked as not done (%d:%s)\n", todo.Id, todo.Content)
			}
			return
		} else {
			fmt.Printf("Todo not found\n")
		}
	}

}

func deleteTodo(reader *bufio.Reader) {
	fmt.Println("Select TODO id")

	// get id
	idStr, _ := reader.ReadString('\n')
	fmt.Println(idStr)
	idStr = strings.TrimSpace(idStr)
	id, err := strconv.Atoi(idStr)

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	for i := range todos {
		if todos[i].Id == id {
			todos = slices.Delete(todos, i, i+1)
			fmt.Printf("Todo deleted on ID: %d \n", id)
			return
		}
	}

	fmt.Println("Todo not found")

}
