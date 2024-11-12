package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ListHabit(fileid int) map[string]interface{} {
	filename := fmt.Sprint(fileid) + ".json"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("No Habit Found!")
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file")
	}
	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error unmarshalling json")
	}
	return jsonData
}
