package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// define flags for edit of items
type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

// function to define commands
func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add new title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit list to specify new title, id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Delete index in the todo list")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify which index to change")
	flag.BoolVar(&cf.List, "list", false, "list all indexes")

	flag.Parse()

	return &cf
}

// Use switch to execute commands
func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])

	case cf.Toggle != -1:
		todos.toggle(cf.Toggle)

	case cf.Del != -1:
		todos.delete(cf.Del)

		//print statement if no matches
	default:
		fmt.Println("Invalid command")
	}
}
