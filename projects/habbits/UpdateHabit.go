package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func (habit *habits) UpdateHabit() {
	err := os.Remove(habit.Id + ".json")
	if err != nil {
		fmt.Println("Error deleting the file")
	}
	jsonData, err := json.Marshal(habit)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(habit.Id+".json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error updating the habit!!")
	}
	fmt.Println("Habit Updated!")
}

func UpdateHabitName(habitId int, habitName string) {
	filename := fmt.Sprint(habitId) + ".json"
	JsonData := ListHabit(habitId)
	JsonData["name"] = habitName
	jsonData, err := json.Marshal(JsonData)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update the habit Name!")
}

func UpdateHabitDescription(habitId int, habitDescription string) {
	filename := fmt.Sprint(habitId) + ".json"
	JsonData := ListHabit(habitId)
	JsonData["description"] = habitDescription
	jsonData, err := json.Marshal(JsonData)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Update the habit Description!")
}

func UpdateHabitCompletionDate(habitId int, completionDate time.Time) {
	filename := fmt.Sprint(habitId) + ".json"
	JsonData := ListHabit(habitId)
	if !completionDate.IsZero() {
		JsonData["completionDate"] = completionDate
	} else {
		JsonData["completionDate"] = time.Now()
	}

	jsonData, err := json.Marshal(JsonData)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	createdDate, err := time.Parse(time.RFC3339, JsonData["createdDate"].(string))
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	difference := completionDate.Sub(createdDate)
	days := int64(difference.Hours() / 24)
	fmt.Println("Update the habit Completion Date!")
	fmt.Println("Habit completed in", days, "days!")
}
