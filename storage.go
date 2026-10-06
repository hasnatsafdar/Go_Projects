package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](filename string) *Storage[T] {
	// Fixed 't' to 'T'
	return &Storage[T]{FileName: filename}
}

func (s *Storage[T]) Save(data T) error {
	// Replaced non-breaking spaces with standard spaces
	fileData, err := json.MarshalIndent(data, "", "    ")

	if err != nil {
		return err
	}

	return os.WriteFile(s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	// Fixed 'Filename' to 'FileName'
	fileData, err := os.ReadFile(s.FileName)

	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, data)
}