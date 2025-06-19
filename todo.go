package main
import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID int
	Title string
	Description string
	Completed bool
}

// Golobal variable to hold tasks
var tasks []Task
var idCounter int = 1

// add task function to add a new task
func add_task(title string, description string){
	task := Task{
		ID: idCounter,
		Title: title,
		Description: description,
		Completed: false,
	}
	tasks = append(tasks, task)
	idCounter++
	fmt.Println("Task added successfully")
}

// View task function to view a specific task by ID
func view_task(task_id int){
	fmt.Println()
	if task_id < 1 || task_id > len(tasks) {
		fmt.Println("Invalid task ID.")
		return
	}
	task := tasks[task_id-1]
	fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %t\n", task.ID, task.Title, task.Description, task.Completed)
}

// View all tasks function to list all tasks
func view_tasks(){
	fmt.Println()
	if len(tasks) == 0{
		fmt.Println("No tasks available.")
		fmt.Println()
		return
	}
	for _, task := range tasks{
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %t\n", task.ID, task.Title, task.Description, task.Completed)
	}
}

// Mark task as completed function to mark a specific task as completed
func mark_task_completed(task_id int){
	fmt.Println()
	if task_id < 1 || task_id > len(tasks){
		fmt.Println("Invalid task ID.")
		return
	}
	tasks[task_id-1].Completed = true
	fmt.Println("Task marked as completed successfully")
}

// Remove task 
func remove_task(task_id int){
	fmt.Println()
	if task_id < 1 || task_id > len(tasks){
		fmt.Println("Invalid task ID.")
		return
	}
	tasks = append(tasks[:task_id-1], tasks[task_id:]...)
	fmt.Println("Task removed successfully")
}


func main(){
	scanner := bufio.NewScanner(os.Stdin)
	for {
		// display welcome message
		fmt.Println()
		fmt.Println("Welcome to the Task Manager!")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Task")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Mark Task as Completed")
		fmt.Println("5. Remove Task")
		fmt.Println("6. Exit")
		fmt.Println()
		fmt.Print("Please select an option (1-6):")

		var choice int
		fmt.Scanln(&choice)

		// Handle user choice. Check first is it wihtin range
		switch choice{
		case 1:
			fmt.Println()
			var title, description string
			
			// read the title
			fmt.Print("Enter task title:")
			scanner.Scan()
			title = scanner.Text()

			// read the description
			fmt.Print("Enter task description:")
			scanner.Scan()
			description = scanner.Text()

			add_task(title, description)
		case 2:
			fmt.Println()
			var task_id int
			fmt.Print("Enter task ID to view:")
			fmt.Scanln(&task_id)
			view_task(task_id)
		case 3:
			view_tasks()
		case 4:
			fmt.Println()
			var task_id int
			fmt.Print("Enter task ID to mark as completed:")
			fmt.Scanln(&task_id)
			mark_task_completed(task_id)
		case 5:
			fmt.Println()
			var task_id int
			fmt.Print("Enter task ID to remove:")
			fmt.Scanln(&task_id)
			remove_task(task_id)
		case 6:
			fmt.Println("Exiting the Task Manager. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option (1-6).")
		}
	}
}
