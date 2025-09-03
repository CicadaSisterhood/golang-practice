package main

import "fmt"

var tasks []Task

// TODO: implement AddTask to append a task to the tasks slice
func AddTask(name string, priority int) (Task, error) {
	// Your code here
	return Task{}, nil
}

// TODO: implement MarkTaskDone to mark a task as done using its index
func MarkTaskDone(index int) error {
	// Your code here
	return nil
}

// TODO: implement RemoveTask to safely remove a task from the slice
func RemoveTask(index int) error {
	// Your code here
	return nil
}

// listTasks is implemented for you so you can see tasks
func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
		return
	}
	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task.Summary())
	}
}
