package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Task      string
	Completed bool
}

var tasks []Task

func addTask(task string) {
	newTask := Task{Task: task, Completed: false}
	tasks = append(tasks, newTask)

	fmt.Println("Task Added")
}

func listTasks() {
	for i, task := range tasks {
		status := "n"
		if task.completed {
			status = "d"
		}

		fmt.Println("%d. %s [%s]\n", i+1, task.Task, status)
	}
}

func markCompleted(index int) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].Completed = true
		fmt.Println("Task Completed")
	} else {
		fmt.Println("Invlid Index")
	}
}
func editTask(index int, newString string) {
	if index >= 1 && index <= len(tasks) {
		tasks[index-1].Task = newString
		fmt.Println("Invalid index")
	} else {
		fmt.Println("Invalid Index")
	}

}

func deleteTask(Index int) {
	if Index >= 1 && Index <= len(tasks) {
		tasks = append(tasks[:Index-1], tasks[Index:]...)
		fmt.Println("Task deleted")
	} else {
		fmt.Println("Invalid index")
	}
}

func main() {
	var indexInput, int
	var indexInput, newTaskInput string

	fmt.Println("Options")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks ")
	fmt.Println("3. Mark as complete")
	fmt.Println("4. Edit task")
	fmt.Println("5. Delete task")
	fmt.Println("6. Exit")

	scanner := bufio.NewScanner((os.Stdin))

	for {
		fmt.Println("Enter choice (1, 2, 3, 4, 5, 6): ")
		if err != nil {
			fmt.Println("Invalid choice")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Enter task")
			scanner.Scan()
			newTaskInput = scanner.Text()
			addTask(newTaskInput)
		case 2:
			listTasks()
		case 3:
			fmt.Println("Enter index: ")
			scanner.Scan()

			indexInput, _ = strconv.Atoi(scanner.Text())
			markCompleted(indexInput)
		case 4:
			fmt.Println("Enter index")
			scanner.Scan()
			newTaskInput = scanner.Text()
			editTask(indexInput, newTaskInput)
		case 6:
			os.Exit(0)
		defualt:
			fmt.Println("Invalid choice")
		}
	}

}
