package main

import "fmt"

var tasks []Task

func AddTask(name string, priority int) (Task, error) {
	task := Task{Name: name, Priority: priority}
	tasks = append(tasks, task)
	return task, nil
}

func MarkTaskDone(index int) error {
	err := validateIndex(index)
	if err != nil {
		return err
	}
	
	tasks[index].MarkDone()
	return nil
}

func RemoveTask(index int) error {
	err := validateIndex(index)
	if err != nil {
		return err
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
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
