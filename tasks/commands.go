package main

import (
	"fmt"
	"strconv"
	"strings"
)

func handleCommand(input string) {
    parts := strings.Fields(input)
    if len(parts) == 0 {
        return
    }

    command := parts[0]
    args := parts[1:]
	_ = args

	
    switch command {
    case "add":
        name := args[0]
		priority, err := strconv.Atoi(args[1])
		if err != nil{
			fmt.Println("Error: second argument must be an integer")
			break
		}
		task, err := AddTask(name, priority)
		if err != nil {
			fmt.Printf("Error: %s", err)
			break
		}
        fmt.Printf("Added task: %s (Priority %d)", task.Name, task.Priority)
    case "list":
        listTasks() // already implemented as a helper in task_manager.go
    case "done":
        index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: argument must be an integer")
			break
		}
		err = MarkTaskDone(index)
		if err != nil {
			fmt.Printf("Error: %s", err)
			break
		}
        fmt.Println("Task marked complete")
    case "remove":
        index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: argument must be an integer")
			break
		}
		err = RemoveTask(index)
		if err != nil {
			fmt.Printf("Error: %s", err)
			break
		}
        fmt.Println("Task removed")
    default:
        fmt.Println("Unknown command:", command)
    }
}
