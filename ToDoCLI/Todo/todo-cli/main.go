package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define CLI flags
	indexFlag := flag.Int("index", -1, "Index to insert the task at")
	taskFlag := flag.String("task", "", "Task to add")
	flag.Parse()

	// task list
	tasks := []string{"Task 1", "Task 2", "Task 3"}

	//split string

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

	// Convert the list to JSON format
	jsonData, err := json.MarshalIndent(taskFlag, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Create or open a file to store the JSON data
	file, err := os.Create("task.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("List successfully written to task.json")
}
