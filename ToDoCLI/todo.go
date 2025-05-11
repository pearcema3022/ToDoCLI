package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// create stuct for headers in cli
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

// add function to define data types
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*todos = append(*todos, todo)
}

// checks removal of methods is valid
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

// delete item and return new list without item removed
func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {

	}

	*todos = append(t[:index], t[index+1:]...)

	return nil

}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed

	if !isCompleted {
		completedTime := time.Now()
		t[index].CompletedAt = &completedTime

	}
	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Completed

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")
	for index, t := range *todos {
		completed := ""
		CompletedAT := ""

		if t.Completed {
			completed = ""
			if t.CompletedAt != nil {
				CompletedAT = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123, CompletedAT))
	}

	// print table
	table.render
}
