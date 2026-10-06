package main

import (
	"errors"
	"fmt"
	"os"
	"strconv" // Added missing import
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t := *todos
	// Remove the element at 'index' by appending the slice after it to the slice before it
	*todos = append(t[:index], t[index+1:]...)

	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil
}

// Fixed receiver capitalization (Todos)
func (todos *Todos) print() {
	// Renamed variable to 'tbl' so it doesn't shadow the 'table' package
	tbl := table.New(os.Stdout) 
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Title", "Completed", "CreatedAt", "CompletedAt")
	
	for index, t := range *todos {
		completed := "✗"
		completedAt := ""

		if t.Completed {
			completed = "✓"
			if t.CompletedAt != nil {
				// Fixed field capitalization (CompletedAt)
				completedAt = t.CompletedAt.Format(time.RFC1123) 
			}
		}

		tbl.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}

	tbl.Render()
}
