package main

//main function which dictates what the table will do and the index's to be added
func main() {
	todos := Todos{}
	todos.add("read documentation")
	todos.add("add flags")
	todos.toggle(0)
	todos.print()
}
