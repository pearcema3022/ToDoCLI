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
	deleteIndex := flag.Int("delete", -1, "Index to be deleted")
	flag.Parse()

	//check to see if delete flag is available
	if *deleteIndex == -1 {
		fmt.Println("you need to specify task to delete")
	}

	// task list
	tasks := []string{"Task 1", "Task 2", "Task 3"}

	// Validate inputs
	if *indexFlag < 0 || *indexFlag > len(tasks) {
		fmt.Println("Error: Invalid index. Please provide a valid index.")
		os.Exit(1)
	}
	// error output if you do not run correct request
	if *taskFlag == "" {
		fmt.Println("Error: Task cannot be empty.")
		os.Exit(1)
	}

	// Allow user to add flag and specified index
	tasks = append(tasks[:*indexFlag], append([]string{*taskFlag}, tasks[*indexFlag:]...)...)

	// Print updated task list.
	fmt.Println("Updated Task List:")
	for i, task := range tasks {
		fmt.Printf("%d: %s\n", i, task)
	}
}
