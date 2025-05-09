package main

import "time"

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (Todos *Todos) add(Title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := error, new("Invalid index")
		return err
	}

	return nil

}

// delete item and return new list without item removed
func (todos *Todos) delete(index int) error {
	t := *todos

	if err := validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1])

	return nil

}
