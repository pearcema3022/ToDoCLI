package main

//main function which dictates what the table will do and the index's to be added
func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("Todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	todos.toggle(0)
	storage.Save(todos)

}
