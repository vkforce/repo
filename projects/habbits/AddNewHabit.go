package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (habit *habits) AddNewHabit() {
	filename := habit.Id + ".json"
	data := map[string]interface{}{
		"name":           habit.Name,
		"description":    habit.Description,
		"createdDate":    habit.CreatedDate,
		"completionDate": habit.CompletionDate,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	// Create the file if it doesn't exist
	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		var file, err = os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Added a New Habit Successfully.")

}
