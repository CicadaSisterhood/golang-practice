package main

import "fmt"

// Task represents a to-do item
type Task struct {
    Name     string
    Priority int
    Done     bool
}

// Taskable interface
type Taskable interface {
    MarkDone()
    Summary() string
}

// TimedTask embeds Task for potential future features
type TimedTask struct {
    Task
    DueDate string
}

func (t *Task) MarkDone() {
    t.Done = true
}

func (t *Task) Summary() string {
	var status string
	if t.Done {
		status = "✅"
	} else {
		status = "❌"
	}
	return fmt.Sprintf("[%s] %s (Priority %d)", status, t.Name, t.Priority)
}
