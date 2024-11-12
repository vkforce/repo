package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var input, habitId int
	fmt.Println("Select from the menu to add, list or update habit. \n1. Add New Habit\n2. List All Habits\n3. Update a Habit")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ = strconv.Atoi(scanner.Text())
	fmt.Print("\nGive/Add the Habit Id:")
	scanner.Scan()
	habitId, _ = strconv.Atoi(scanner.Text())

	switch input {
	case 1:
		var habitName, habitDescription string
		for {
			fmt.Print("Enter the habit name (or 'q' to quit): ")
			scanner.Scan()
			habitName = scanner.Text()
			if habitName == "q" {
				break
			}
			fmt.Print("\nEnter the habit description: ")
			scanner.Scan()
			habitDescription = scanner.Text()
			fmt.Printf("Habit: %s\nDescription: %s\n\n", habitName, habitDescription)
			if habitName != "" && habitDescription != "" {
				myhabit := habits{Name: habitName, Description: habitDescription, CreatedDate: time.Now(), CompletionDate: nil, Id: fmt.Sprintf("%d", habitId)}
				myhabit.AddNewHabit()
				break
			} else {
				fmt.Println("Please enter valid habit name and description")
			}
		}

	case 2:
		jsonData := ListHabit(habitId)
		fmt.Printf("Habit Name: %s\nDescription: %s\nCreated Date: %s\nCompletion Date: %s\n", jsonData["name"], jsonData["description"], jsonData["createdDate"], jsonData["completionDate"])

	case 3:
		fmt.Println("What do you wanna update? \n1. Name\n2. Description\n3. Completion Date\n4. All")
		scanner.Scan()
		updateInput, _ := strconv.Atoi(scanner.Text())
		switch updateInput {
		case 1:
			var habitname string
			fmt.Print("Enter the habit name: ")
			scanner.Scan()
			habitname = scanner.Text()
			UpdateHabitName(habitId, habitname)
		case 2:
			var habitDescription string
			fmt.Print("Enter the habit description: ")
			scanner.Scan()
			habitDescription = scanner.Text()
			fmt.Println("Updated Description: ", habitDescription)
			UpdateHabitDescription(habitId, habitDescription)
		case 3:
			var completionDateString string
			fmt.Print("Enter the completion date: ")
			scanner.Scan()
			completionDateString = scanner.Text()
			completionDate, err := time.Parse(time.RFC3339, completionDateString) // Assuming RFC3339 format
			if err != nil {
				fmt.Println("Error parsing date:", err)
				return
			}
			UpdateHabitCompletionDate(habitId, completionDate)
		case 4:
			var habitName, habitDescription string
			for {
				fmt.Print("Enter the habit name (or 'q' to quit): ")
				scanner.Scan()
				habitName = scanner.Text()
				if habitName == "q" {
					break
				}
				fmt.Print("\nEnter the habit description: ")
				scanner.Scan()
				habitDescription = scanner.Text()
				fmt.Printf("Habit: %s\nDescription: %s\n\n", habitName, habitDescription)
				if habitName != "" && habitDescription != "" {
					myhabit := habits{Name: habitName, Description: habitDescription, CreatedDate: time.Now(), CompletionDate: nil, Id: fmt.Sprintf("%d", habitId)}
					myhabit.UpdateHabit()
					break
				} else {
					fmt.Println("Please enter valid habit name and description")
				}
			}

		default:
			fmt.Println("Invalid Input")
		}
	}
}
