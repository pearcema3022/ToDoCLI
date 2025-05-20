package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define CLI flags
	indexFlag := flag.Int("index", -1, "Index to insert the task at")
	taskFlag := flag.String("task", "", "Task to add")
	flag.Parse()

	// Example task list
	tasks := []string{"Task 1", "Task 2", "Task 3"}

	// Validate inputs
	if *indexFlag < 0 || *indexFlag > len(tasks) {
		fmt.Println("Error: Invalid index. Please provide a valid index.")
		os.Exit(1)
	}
	if *taskFlag == "" {
		fmt.Println("Error: Task cannot be empty.")
		os.Exit(1)
	}

	// Insert task at the specified index
	tasks = append(tasks[:*indexFlag], append([]string{*taskFlag}, tasks[*indexFlag:]...)...)

	// Print updated task list
	fmt.Println("Updated Task List:")
	for i, task := range tasks {
		fmt.Printf("%d: %s\n", i, task)
	}
}
