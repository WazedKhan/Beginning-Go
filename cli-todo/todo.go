package main

import (
	"cli-todo/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type Task struct {
	Description string
	Completed   bool
}

// createTask prompts the user for a task description and adds it to the task list.
func createTask(tasks []Task) []Task {
	scanner := utils.NewScanner()
	fmt.Print("Enter task description: ")
	scanner.Scan()
	// Trimming leading/trailing spaces
	taskDescription := strings.TrimSpace(scanner.Text())

	// Validate that input is not empty
	if taskDescription == "" {
		fmt.Println("Task description cannot be empty.")
		return tasks
	}

	// Append the new task to the list and confirm creation
	tasks = append(tasks, Task{Description: taskDescription, Completed: false})
	color.Green("Task created successfully!")
	return tasks
}

// showTasks displays the list of tasks, formatted with task numbers.
func showTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks available!")
		return
	}

	// Print all tasks with their index
	fmt.Println("Your tasks:")
	for idx, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[X]"
		}
		fmt.Printf("[%d]: %s %s\n", idx+1, status, task.Description)
	}
}

func removeLastTask(tasks []Task) []Task {
	if len(tasks) == 0 {
		fmt.Println("No tasks to remove.")
		return tasks
	}

	lastTask := tasks[len(tasks)-1]
	tasks = tasks[:len(tasks)-1]
	fmt.Printf("Removed last task: %s\n", lastTask.Description)

	return tasks
}

func markTask(tasks []Task) []Task {
	if len(tasks) == 0 {
		fmt.Println("No tasks to mark as complete.")
		return tasks
	}

	scanner := utils.NewScanner()
	fmt.Print("Enter the task number to mark as complete: ")
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())

	// Convert input to integer
	taskNumber, err := strconv.Atoi(input)
	if err != nil || taskNumber < 1 || taskNumber > len(tasks) {
		fmt.Println("Invalid task number!")
		return tasks
	}
	// mark the task
	tasks[taskNumber-1].Completed = true
	fmt.Printf("Task [%d] marked as complete\n", taskNumber)
	return tasks
}

// main is the entry point where the CLI interacts with the user.
func main() {
	// Initialize an empty slice for tasks
	var tasks []Task
	scanner := utils.NewScanner()

	fmt.Println("Welcome to the CLI To-Do List!")

	// Main loop for accepting commands
	for {
		fmt.Print("Enter a command (add, list, mark, exit): ")
		scanner.Scan()
		command := strings.ToLower(strings.TrimSpace(scanner.Text()))

		switch command {
		case "add":
			tasks = createTask(tasks)

		case "list":
			showTasks(tasks)

		case "mark":
			markTask(tasks)

		case "remove":
			tasks = removeLastTask(tasks)

		case "exit":
			// Exit the loop and program
			fmt.Println("Exiting the application.")
			return

		default:
			fmt.Println("Invalid command. Please use 'add', 'list', 'mark' or 'exit'.")
		}
	}
}
