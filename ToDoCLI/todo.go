package main

import (
	"errors"
	"time"
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
func (todos *Todos) add(Title string) {
	todo := Todo{
		Title:       Title,
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
		fmt.println.err
		return err
	}

	return nil
}

// delete item and return new list without item removed
func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}
	//return new list
	*todos = append(t[:index], t[index+1]...)

	return nil
}
