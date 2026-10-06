package main

func main() {
	todos := Todos{}
	
	// Initialize storage
	storage := NewStorage[Todos]("todos.json")
	
	// Load existing todos from the file. 
	// (Note: It's okay if this fails on the very first run because the file doesn't exist yet)
	storage.Load(&todos)
	
	// Parse the command line flags
	cmdFlags := NewCmdFlags()
	
	// Execute the requested command (add, edit, list, etc.)
	// Fixed 'cmdflags' to 'cmdFlags'
	cmdFlags.Execute(&todos)
	
	// Save the updated list back to the file
	storage.Save(todos)
}