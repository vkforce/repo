package main

import "time"

type habits struct {
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	CreatedDate    time.Time       `json:"createdDate"`
	CompletionDate map[string]bool `json:"completionDate"`
}
