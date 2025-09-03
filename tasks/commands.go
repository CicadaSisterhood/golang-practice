package main

import (
	"fmt"
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
        // TODO: parse args and call AddTask
        fmt.Println("Add task not implemented yet")
    case "list":
        listTasks() // already implemented as a helper in task_manager.go
    case "done":
        // TODO: parse args and call MarkTaskDone
        fmt.Println("Mark task done not implemented yet")
    case "remove":
        // TODO: parse args and call RemoveTask
        fmt.Println("Remove task not implemented yet")
    default:
        fmt.Println("Unknown command:", command)
    }
}
