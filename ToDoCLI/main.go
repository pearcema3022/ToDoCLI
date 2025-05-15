package main

import (
	"log/slog"
)

// main function which dictates what the table will do and the index's to be added, edited, deleted, toggled.
func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("Todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	todos.toggle(0)
	storage.Save(todos)
	slog.Debug("Debug message")
	slog.Info("Index added")
	slog.Error("Error occured")

}
