package main

import (
	"encoding/json"
	"os"
)

// storage type todo
type storage[T any] struct {
	FileName string
}

func NewStorage[T any](FileName string) *storage[T] {
	return &storage[T]{FileName: FileName}
}

// converting data into json and checking for error
func (s *storage[T]) Save(data T) error {
	filedata, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}
	//permission to everyone can read file
	return os.WriteFile(s.FileName, filedata, 0644)
}

func (s *storage[T]) Load(data *T) error {
	filedata, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}
	//convert json and populate
	return json.Unmarshal(filedata, data)
}
