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
