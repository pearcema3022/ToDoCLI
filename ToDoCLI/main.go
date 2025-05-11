package main

func main() {
	todos := Todos{}
	todos.add("read documentation")
	todos.add("add flags")
	todos.toggle(0)
	todos.print()
}
