package main

import (
	"testing"
)

func TestAddTask(t *testing.T) {
    tasks = []Task{} // reset slice before test

    taskName := "Write tests"
    priority := 2

    task, err := AddTask(taskName, priority)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if len(tasks) != 1 {
        t.Fatalf("Expected 1 task, got %d", len(tasks))
    }

    if task.Name != taskName || task.Priority != priority || task.Done != false {
        t.Errorf("Task not added correctly: %+v", task)
    }
}

func TestMarkTaskDone(t *testing.T) {
    tasks = []Task{{Name: "Write code", Priority: 1, Done: false}}

    err := MarkTaskDone(0)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if !tasks[0].Done {
        t.Errorf("Expected task to be done, got %+v", tasks[0])
    }
}

func TestRemoveTask(t *testing.T) {
    tasks = []Task{
        {Name: "Task 1", Priority: 1, Done: false},
        {Name: "Task 2", Priority: 2, Done: false},
    }

    err := RemoveTask(0)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if len(tasks) != 1 || tasks[0].Name != "Task 2" {
        t.Errorf("Task not removed correctly, tasks: %+v", tasks)
    }
}

func TestInvalidIndices(t *testing.T) {
    tasks = []Task{{Name: "Only task", Priority: 1, Done: false}}

    if err := MarkTaskDone(5); err == nil {
        t.Errorf("Expected error for invalid index in MarkTaskDone")
    }

    if err := RemoveTask(-1); err == nil {
        t.Errorf("Expected error for invalid index in RemoveTask")
    }
}

func TestTaskSummary(t *testing.T) {
    task := Task{Name: "Example", Priority: 3, Done: false}
    expected := "[❌] Example (Priority 3)"
    if task.Summary() != expected {
        t.Errorf("Expected summary '%s', got '%s'", expected, task.Summary())
    }

    task.MarkDone()
    expectedDone := "[✅] Example (Priority 3)"
    if task.Summary() != expectedDone {
        t.Errorf("Expected summary '%s', got '%s'", expectedDone, task.Summary())
    }
}

func TestAddMultipleTasks(t *testing.T) {
    tasks = []Task{}

    taskNames := []string{"Task A", "Task B", "Task C"}
    priorities := []int{1, 2, 3}

    for i := 0; i < len(taskNames); i++ {
        _, err := AddTask(taskNames[i], priorities[i])
        if err != nil {
            t.Fatalf("Unexpected error adding task %d: %v", i, err)
        }
    }

    if len(tasks) != len(taskNames) {
        t.Errorf("Expected %d tasks, got %d", len(taskNames), len(tasks))
    }

    for i, task := range tasks {
        if task.Name != taskNames[i] || task.Priority != priorities[i] {
            t.Errorf("Task %d mismatch: got %+v, want name=%s priority=%d", i, task, taskNames[i], priorities[i])
        }
    }
}

func TestTaskableInterface(t *testing.T) {
    var tInterface Taskable = &Task{Name: "Interface Test", Priority: 1, Done: false}

    // Initially not done
    expected := "[❌] Interface Test (Priority 1)"
    if tInterface.Summary() != expected {
        t.Errorf("Expected summary '%s', got '%s'", expected, tInterface.Summary())
    }

    tInterface.MarkDone()
    expectedDone := "[✅] Interface Test (Priority 1)"
    if tInterface.Summary() != expectedDone {
        t.Errorf("Expected summary '%s', got '%s'", expectedDone, tInterface.Summary())
    }
}
