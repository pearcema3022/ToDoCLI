package main

import "fmt"

func main() {
	todos := Todos{}
	todos.add("read documentation")
	todos.add("include flag")
	todos.add("add traceid")
	fmt.Printf("%+v\n\n", todos)
	todos.Delete(0)
	fmt.Printf("%+v", todos)
}
