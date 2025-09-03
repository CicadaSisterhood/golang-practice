package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Mini Task Manager CLI")

    for {
        fmt.Print("\nEnter command (add/list/done/remove/exit): ")
        if !scanner.Scan() {
            break
        }
        input := strings.TrimSpace(scanner.Text())
        if input == "exit" {
            fmt.Println("Goodbye!")
            break
        }

        handleCommand(input) // implemented in commands.go
    }
}
