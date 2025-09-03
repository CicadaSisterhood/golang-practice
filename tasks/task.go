package main

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

// TODO: implement MarkDone so it updates the task's status
func (t *Task) MarkDone() {
    // Your code here
}

// TODO: implement Summary to return a string like "[❌] Name (Priority X)"
func (t *Task) Summary() string {
    // Your code here
    return ""
}
